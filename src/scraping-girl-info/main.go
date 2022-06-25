package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"local.packages/src/elegaku"

	"github.com/PuerkitoBio/goquery"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

// 本来はenvから取得した方が良い
const AWS_REGION = "ap-northeast-1"
const DYNAMO_ENDPOINT = "http://localhost:8000"

// 在籍情報の追加・更新
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

	table := db.Table("girls")

	// 最新の在籍情報を取得
	girls, err := getGitls()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// 取得した在籍情報を登録する。
	for _, g := range girls {
		err := table.Put(g).Run()

		if err != nil {
			fmt.Println(err.Error())
			break
		}
	}
}

// 最新の在籍情報を取得（WEBスクレイピング）
func getGitls() ([]elegaku.Girl, error) {
	webPage := ("https://www.elegaku.com/cast/")
	resp, err := http.Get(webPage)
	if err != nil {
		log.Printf("failed to get html: %s", err)
		return nil, errors.New("スクレイピング失敗！")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("failed to fetch data: %d %s", resp.StatusCode, resp.Status)
		return nil, errors.New("スクレイピング失敗！")
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Printf("failed to load html: %s", err)
		return nil, errors.New("スクレイピング失敗！")
	}

	results := []elegaku.Girl{}
	doc.Find("#companion_box").Each(func(i int, sGirl *goquery.Selection) {
		// GirlIdを取得
		girlId, _ := sGirl.Find("div.g_image > a").Attr("href")

		// 名前と年齢を取得
		nameAndAge := strings.TrimSpace(sGirl.Find(".name > a").Text())
		length := len(nameAndAge)

		// 初期化・セット
		girl := elegaku.Girl{}
		girl.GirlId = regexp.MustCompile("[^0-9]").ReplaceAllString(girlId, "")
		girl.Name = nameAndAge[0 : length-2]
		girl.Age, _ = strconv.Atoi(nameAndAge[length-2 : length])
		girl.ThreeSize = sGirl.Find(".size").Text()
		girl.CatchCopy = sGirl.Find(".catch").Text()
		girl.Image, _ = sGirl.Find("div.g_image > a").Children().Attr("src")
		girl.CreateDatetime = elegaku.GetTimestamp()
		girl.UpdateDatetime = elegaku.GetTimestamp()

		results = append(results, girl)
	})

	return results, nil
}
