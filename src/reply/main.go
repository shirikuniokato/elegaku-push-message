package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/line/line-bot-sdk-go/linebot"
	"local.packages/src/elegaku"
)

/* 返信 */
func reply(bot *linebot.Client, webhook elegaku.WebHook) error {
	// リクエストされたイベントの件数分処理する
	for _, event := range webhook.Events {
		_, err := bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("テスト")).Do()
		fmt.Println(err)

		switch event.Type {
		case linebot.EventTypeFollow: // フォローイベント
			fmt.Println("*** event follow")
			Follow()
		case linebot.EventTypeUnfollow: // フォロー解除イベント
			fmt.Println("*** event unfollow")
			UnFollow()
		case linebot.EventTypePostback: // ポストバックイベント
			fmt.Println("*** event postback")
			Postback()
		default:
			fmt.Println("*** event " + event.Type)
			fmt.Println("処理対象外のイベント")
		}
	}
	return nil
}

func handler(request events.APIGatewayProxyRequest) error {
	// リクエスト内容をデコードする
	var webhook elegaku.WebHook
	if err := json.Unmarshal([]byte(request.Body), &webhook); err != nil {
		return nil
	}

	// ボットの定義
	fmt.Println("*** linebot new")
	bot, err := linebot.New(
		os.Getenv("LINE_CHANNEL_SECRET"),
		os.Getenv("LINE_CHANNEL_ACCESS_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	// 返信
	err = reply(bot, webhook)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return nil
}

func main() {
	lambda.Start(handler)
}
