package elegaku

import (
	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// SQSに格納する通知情報
type PushInfo struct {
	UserIds    []string `json:"user_ids,omitempty"`
	GirlId     string   `json:"girl_id"`
	TargetDate string   `json:"target_date"`
	Image      string   `json:"image"`
	NameAndAge string   `json:"name_and_age"`
	ThreeSize  string   `json:"three_size"`
	CatchCopy  string   `json:"catch_copy"`
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
func PushSQS(svc *sqs.SQS, infoList []PushInfo) {
	// 通知情報をJSONに変換する
	msg, _ := json.Marshal(infoList)

	// 通知情報を格納する
	svc.SendMessage(&sqs.SendMessageInput{
		QueueUrl:    aws.String("https://sqs.ap-northeast-1.amazonaws.com/856051715637/attendance_notification"),
		MessageBody: aws.String(string(msg)),
	})
}
