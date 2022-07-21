package main

import (
	"fmt"

	"local.packages/src/elegaku"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/guregu/dynamo"
)

// 全ユーザー情報
var userList []elegaku.User

// 出勤対象のGirlIdと紐づくUserIDを取得
func getNotificationList() []elegaku.PushInfo {
	// SQSに格納するリスト
	pushInfoList := []elegaku.PushInfo{}
	// DynamoDBに接続
	db := elegaku.ConnectDB()
	// ユーザー情報の取得
	getAllUserList(db)

	// １週間分の加算値
	week := []int{0, 1, 2, 3, 4, 5, 6}
	t := elegaku.GetTimeJst()
	for _, d := range week {
		// テーブル名取得
		targetDate := t.AddDate(0, 0, d)
		tableName := targetDate.Format(elegaku.YMD_NUM_FMT)
		fmt.Println(tableName + " start")

		// テーブル取得
		table := db.Table(tableName)

		// 未通知の出勤情報を取得
		var scheduleList []elegaku.Schedule
		err := table.Scan().Filter("'notice_flg' = ?", 0).All(&scheduleList)

		if err != nil {
			// 通知対象の女の子がいない場合は次の日へ
			continue
		}

		pushInfoList = append(pushInfoList, bindUserAndGirl(db, tableName, scheduleList)...)
	}

	return pushInfoList
}

// ユーザー情報の取得
func getAllUserList(db *dynamo.DB) {
	// テーブル取得
	table := db.Table(elegaku.TBLNM_USERS)
	err := table.Scan().All(&userList)

	// ユーザー情報の取得に失敗した場合は処理を終了する（本来取得できないことはない想定）
	if err != nil {
		panic(err)
	}
}

// ユーザ情報と女の子の情報を紐づける
func bindUserAndGirl(db *dynamo.DB, targetDate string, schedule []elegaku.Schedule) []elegaku.PushInfo {
	results := []elegaku.PushInfo{}

	// イメージを取得するためにGirlTableから取得する
	girlTable := db.Table(elegaku.TBLNM_GIRLS)

	for _, s := range schedule {
		// GirlTableから女の子の情報を取得
		var g = elegaku.Girl{}
		err := girlTable.Get(elegaku.G_GIRL_ID, s.GirlId).One(&g)
		if err != nil {
			// 女の子の情報が取得できなかった場合は次の通知情報へ
			continue
		}

		// SQSに格納する情報を初期化する
		info := elegaku.PushInfo{TargetDate: targetDate, GirlId: s.GirlId, Image: g.Image}

		// 女の子とユーザを紐づける
		for _, u := range userList {
			// 紐づく場合、ユーザーIDを追加する
			if elegaku.Contains(u.FavoriteGirlIds, s.GirlId) {
				info.UserIds = append(info.UserIds, u.UserId)
			}
		}

		// ユーザ情報が紐づかない場合は次の通知情報へ
		if info.UserIds == nil {
			continue
		}

		// 通知情報とユーザが紐づいた場合に返却リストに追加する
		results = append(results, info)
	}

	return results
}

// 出勤情報の追加・更新処理を呼び出す
func HandleLambdaEvent() {
	// 通知情報を取得
	pushInfoList := getNotificationList()
	// SQSに接続
	sqs := elegaku.ConnectSQS()
	// SQSに通知情報を格納する
	elegaku.PushSQS(sqs, pushInfoList)
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
