package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/line/line-bot-sdk-go/linebot"
)

type MessageObj struct {
	UserID  string
	Message string
}

/* メッセージ送信 */
func sendMessage(bot *linebot.Client, p *events.SQSMessage) error {
	// 取得したMessageをデコードする。
	message := MessageObj{"", ""}
	if err := json.NewDecoder(strings.NewReader(p.Body)).Decode(&message); err != nil {
		fmt.Println(err)
		return fmt.Errorf("massages decode error.[%s]", p.Body)
	}

	fmt.Println("*** push")
	if _, err := bot.PushMessage(message.UserID, linebot.NewTextMessage(message.Message)).Do(); err != nil {
		log.Fatal(err)
		return fmt.Errorf("massages push error.[%s]", message.UserID)
	}

	return nil
}

func handler(ctx context.Context, sqsEvent events.SQSEvent) error {

	// ボットの定義
	fmt.Println("*** linebot new")
	bot, err := linebot.New(
		os.Getenv("CHANNELSECRET"),
		os.Getenv("ACCESSTOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// メッセージをSQSから取得
	for _, message := range sqsEvent.Records {
		// メッセージ送信
		err := sendMessage(bot, &message)
		if err != nil {
			log.Fatal(err)
		}
	}

	// 終了
	fmt.Println("*** end")
	return nil
}

func main() {
	lambda.Start(handler)
}
