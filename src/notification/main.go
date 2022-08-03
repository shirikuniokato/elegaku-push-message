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
	"local.packages/src/elegaku"
)

/* メッセージ送信 */
func sendMessage(bot *linebot.Client, p *events.SQSMessage) error {
	// 取得したMessageをデコードする。
	fmt.Println("*** message decode")
	message := elegaku.PushInfo{}
	if err := json.NewDecoder(strings.NewReader(p.Body)).Decode(&message); err != nil {
		fmt.Println(err)
		return fmt.Errorf("massages decode error.[%s]", p.Body)
	}

	fmt.Println("*** push")
	pushMsg := fmt.Sprintf(
		elegaku.PUSH_MESSAGE_FOMAT,
		message.TargetDate,
		message.Image,
		message.NameAndAge,
		message.GirlId,
		message.ThreeSize,
		message.CatchCopy,
		message.GirlId,
	)

	container, err := linebot.UnmarshalFlexMessageJSON([]byte(pushMsg))
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("flex message decode error.[%s]", message.GirlId)
	}

	if _, err := bot.Multicast(message.UserIds, linebot.NewFlexMessage("", container)).Do(); err != nil {
		log.Fatal(err)
		return fmt.Errorf("massages push error.[%s]", message.GirlId)
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
