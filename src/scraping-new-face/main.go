package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"local.packages/src/elegaku"

	"github.com/PuerkitoBio/goquery"
)

// 新入生取得
func main() {
	// クライアントの設定
	db := elegaku.ConnectDB()
	table := db.Table(elegaku.TBLNM_NEW_FACE)
	newFaces, err := getNewFaces()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// 取得した在籍情報を登録する。
	for _, n := range newFaces {
		table.Delete(elegaku.N_GIRL_ID, n.GirlId).Run()
		table.Put(n).Run()
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
