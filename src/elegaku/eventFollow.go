package elegaku

import (
	"github.com/line/line-bot-sdk-go/linebot"
)

// 友達追加
func Follow(bot *linebot.Client, event *linebot.Event) {
	// ユーザ名取得
	userName, err1 := getUserName(bot, event.Source.UserID)
	if err1 != nil {
		bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("ユーザ名取得に失敗しました")).Do()
	}

	// ユーザ情報をDynamoDBに登録
	err2 := insertUser(event.Source.UserID, userName)
	if err2 != nil {
		bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("ユーザ登録に失敗しました")).Do()
	}
}

// ユーザー名取得
func getUserName(bot *linebot.Client, userId string) (string, error) {
	res, err := bot.GetProfile(userId).Do()
	if err != nil {
		return "", err
	}
	return res.DisplayName, nil
}

// ユーザー情報をDynamoDBに登録
func insertUser(userId string, userName string) error {
	// ユーザ情報をDynamoDBに登録
	err := ConnectDB().Table(TBLNM_USERS).Put(User{UserId: userId, UserName: userName}).Run()
	if err != nil {
		return err
	}
	return nil
}
