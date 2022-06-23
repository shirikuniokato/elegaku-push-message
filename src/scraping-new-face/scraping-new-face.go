package main

import (
	"elegaku"
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/PuerkitoBio/goquery"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

// 本来はenvから取得した方が良い
const AWS_REGION = "ap-northeast-1"
const DYNAMO_ENDPOINT = "http://localhost:8000"

// 新入生取得
func main() {
	// クライアントの設定
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(AWS_REGION),
		Endpoint:    aws.String(DYNAMO_ENDPOINT),
		Credentials: credentials.NewStaticCredentials("dummy", "dummy", "dummy"),
	})
	if err != nil {
		panic(err)
	}

	db := dynamo.New(sess)
	table := db.Table("new_face")
	newFaces, err := getNewFaces()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// 取得した在籍情報を登録する。
	for _, n := range newFaces {
		table.Delete("girl_id", n.GirlId).Run()
		err := table.Put(n).Run()

		if err != nil {
			fmt.Println(err.Error())
			break
		}
	}
}

// 最新の新入生情報を取得
func getNewFaces() ([]elegaku.NewFace, error) {
	webPage := ("https://www.elegaku.com/newface/")
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
		return nil, errors.New("スクレイピング失敗！")
	}

	// 新入生を取得を取得
	results := []elegaku.NewFace{}
	doc.Find("#companion_box").Each(func(i int, sGirl *goquery.Selection) {
		// GirlIdの取得
		g, _ := sGirl.Find("div.g_image > a").Attr("href")

		// 初期化・セット・追加
		results = append(results, elegaku.NewFace{GirlId: regexp.MustCompile("[^0-9]").ReplaceAllString(g, ""), CreateDatetime: elegaku.GetTimestamp(), UpdateDatetime: elegaku.GetTimestamp()})
	})
	return results, nil
}
