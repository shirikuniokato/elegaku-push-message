package elegaku

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// SQSに格納する通知情報
type PushInfo struct {
	TargetDate string
	GirlId     string
	Image      string
	UserIds    []string
}

// SQSに接続
func ConnectSQS() *sqs.SQS {
	// クライアントの設定
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1")},
	)
	if err != nil {
		panic(err)
	}

	return sqs.New(sess)
}

// SQSにPushInfoを格納する
func PushSQS(svc *sqs.SQS, info PushInfo) {
	// TODO 実装悩み中
	// 通知情報を格納する
	msg := fmt.Sprintf("%s,%s,%s,%s", info.TargetDate, info.GirlId, info.Image, info.UserIds) // 通知情報を文字列に変換
	svc.SendMessage(&sqs.SendMessageInput{
		QueueUrl:    aws.String("https://sqs.us-west-2.amazonaws.com/123456789012/SQS_QUEUE_NAME"),
		MessageBody: aws.String(msg),
	})
}
