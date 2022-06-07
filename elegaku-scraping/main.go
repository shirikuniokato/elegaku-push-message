package main

import (
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Girl struct {
	GirlId    string `json:"girl_id"`
	Name      string `json:"name"`
	Age       int    `json:"age"`
	ThreeSize string `json:"three_size"`
	CatchCopy string `json:"catch_copy"`
	Image     string `json:"image"`
}

// 在籍情報の追加・更新
func main() {
	webPage := ("https://www.elegaku.com/cast/")
	resp, err := http.Get(webPage)
	if err != nil {
		log.Printf("failed to get html: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Fatalf("failed to fetch data: %d %s", resp.StatusCode, resp.Status)
		return
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Printf("failed to load html: %s", err)
		return
	}

	girls := []Girl{}
	doc.Find("#companion_box").Each(func(i int, sGirl *goquery.Selection) {
		// GirlIdを取得
		girlId, _ := sGirl.Find("div.g_image > a").Attr("href")

		// 名前と年齢を取得
		nameAndAge := strings.TrimSpace(sGirl.Find(".name > a").Text())
		length := len(nameAndAge)

		// 初期化・セット
		girl := Girl{}
		girl.GirlId = regexp.MustCompile("[^0-9]").ReplaceAllString(girlId, "")
		girl.Name = nameAndAge[0 : length-2]
		girl.Age, _ = strconv.Atoi(nameAndAge[length-2 : length])
		girl.ThreeSize = sGirl.Find(".size").Text()
		girl.CatchCopy = sGirl.Find(".catch").Text()
		girl.Image, _ = sGirl.Find("div.g_image > a").Children().Attr("src")

		girls = append(girls, girl)
	})

	// TODO■DynamoDBへの登録・更新処理が必要
}
