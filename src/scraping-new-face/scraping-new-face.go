package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

type NewFace struct {
	GirlId string `json:"girl_id"`
}

type Rank struct {
	Rank   int    `json:"rank"`
	GirlId string `json:"girl_id"`
}

// 新入生取得
func main() {
	webPage := ("https://www.elegaku.com/newface/")
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

	// 新入生を取得を取得
	newFaces := []NewFace{}
	doc.Find("#companion_box").Each(func(i int, sGirl *goquery.Selection) {
		// GirlIdの取得
		g, _ := sGirl.Find("div.g_image > a").Attr("href")

		// 初期化・セット・追加
		newFaces = append(newFaces, NewFace{regexp.MustCompile("[^0-9]").ReplaceAllString(g, "")})
	})

	fmt.Println(newFaces)

	// TODO■DynamoDBへの登録・更新処理が必要
}
