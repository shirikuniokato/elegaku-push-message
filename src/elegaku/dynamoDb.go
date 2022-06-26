package elegaku

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

// TODO■Lambdaに乗っける前に環境変数から取得するように修正する
const AWS_REGION = "ap-northeast-1"
const DYNAMO_ENDPOINT = "http://localhost:8000"

// DynamoDBに接続
func ConnectDB() *dynamo.DB {
	// クライアントの設定
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(AWS_REGION),
		Endpoint:    aws.String(DYNAMO_ENDPOINT),
		Credentials: credentials.NewStaticCredentials("dummy", "dummy", "dummy"),
	})
	if err != nil {
		panic(err)
	}

	return dynamo.New(sess)
}

// タイムスタンプの取得
func GetTimestamp() string {
	// 現在時刻をyyyy-MM-dd hh:mm:ss形式で取得
	return time.Now().Format("2006-01-02 15:04:05")
}
