package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"local.packages/src/elegaku"

	"github.com/PuerkitoBio/goquery"
	"github.com/aws/aws-lambda-go/lambda"
)

// 出勤情報をDynamoDBに追加・更新
func putSchedule() {
	fmt.Println("scraping start")
	// DynamoDBに接続
	db := elegaku.ConnectDB()
	// １週間分の加算値
	week := []int{0, 1, 2, 3, 4, 5, 6}
	t := time.Now()
	for _, v := range week {
		// テーブル名取得
		targetDate := t.AddDate(0, 0, v)
		tableName := targetDate.Format("20060102")
		fmt.Println(tableName + " start")

		// テーブル作成・取得
		db.CreateTable(tableName, elegaku.Schedule{}).Run()
		table := db.Table(tableName)

		// 出勤情報取得
		schedule, err := getSchedule(targetDate.Format("y/2006/MM/01/dd/02"))
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		// 出勤情報登録
		for _, s := range schedule {
			var base elegaku.Schedule
			table.Get(elegaku.S_GIRL_ID, s.GirlId).One(&base)

			// 出勤情報が未登録の場合はPUT
			if base.GirlId == "" {
				table.Put(s).Run()
				continue
			}

			// 出勤時間が変更していた場合のみ更新する
			// ただし、出勤時間未確定の場合は更新対象外とする
			if s.Time != base.Time && s.Time != "" {
				// 変更している場合
				table.Update(elegaku.S_GIRL_ID, s.GirlId).Set(elegaku.S_NOTICE_FLG, 0).Set(elegaku.S_TIME, s.Time).Set(elegaku.S_UPDATE_DATE_TIME, elegaku.GetTimestamp()).Run()
				continue
			} else {
				// 変更していない場合
				continue
			}
		}
		fmt.Println(tableName + " end")
	}
	fmt.Println("scraping end")
}

// 最新の出勤情報取得
func getSchedule(date string) ([]elegaku.Schedule, error) {
	webPage := ("https://www.elegaku.com/cast/schedule/" + date)
	resp, err := http.Get(webPage)
	if err != nil {
		log.Printf("failed to get html: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("failed to fetch data: %d %s", resp.StatusCode, resp.Status)
		return nil, errors.New("スクレイピング失敗！")
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Printf("failed to load html: %s", err)
		return nil, errors.New("スクレイピング失敗！")
	}

	results := []elegaku.Schedule{}

	doc.Find("#companion_box").Each(func(i int, sGirl *goquery.Selection) {
		s := elegaku.Schedule{}

		// GirlIdの取得
		g, _ := sGirl.Find("div.g_image > a").Attr("href")

		// 初期化・セット
		s.GirlId = regexp.MustCompile("[^0-9]").ReplaceAllString(g, "")
		s.Time = strings.TrimSpace(sGirl.Find(".time").Text())
		s.NoticeFlg = 0
		s.CreateDatetime = elegaku.GetTimestamp()
		s.UpdateDatetime = elegaku.GetTimestamp()

		results = append(results, s)
	})

	return results, nil
}

// 出勤情報の追加・更新処理を呼び出す
func HandleLambdaEvent() {
	putSchedule()
}

func main() {
	lambda.Start(HandleLambdaEvent)
}
