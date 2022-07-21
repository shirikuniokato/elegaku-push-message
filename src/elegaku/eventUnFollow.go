package elegaku

import "github.com/line/line-bot-sdk-go/linebot"

// ブロック等
func UnFollow(event *linebot.Event) {
	// ユーザ情報を削除
	ConnectDB().Table(TBLNM_USERS).Delete(event.Source.UserID, User{}).Run()
}
