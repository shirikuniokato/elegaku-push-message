package elegaku

import (
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

// DynamoDBに接続
func ConnectDB() *dynamo.DB {
	// クライアントの設定
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("ap-northeast-1"),
		Endpoint:    aws.String("dynamodb.ap-northeast-1.amazonaws.com"),
		Credentials: credentials.NewStaticCredentials(os.Getenv("ACCESS_KEY"), os.Getenv("SECRET_ACCESS_KEY"), "dummy"),
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
