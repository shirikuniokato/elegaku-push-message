package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"local.packages/src/elegaku"

	"github.com/PuerkitoBio/goquery"
)

// ランキングの更新
func main() {
	// DynamoDBに接続
	db := elegaku.ConnectDB()
	table := db.Table(elegaku.TBLNM_RANK)

	// 最新の在籍情報を取得
	rank, err := getRank()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// 取得した在籍情報を登録する。
	for _, r := range rank {
		table.Delete(elegaku.R_RANK, r.Rank).Run()
		table.Put(r).Run()
	}
}

// 最新のランキング更新
func getRank() ([]elegaku.Rank, error) {
	webPage := ("https://www.elegaku.com/rank/")
	resp, err := http.Get(webPage)
	if err != nil {
		log.Printf("failed to get html: %s", err)
		return nil, errors.New("スクレイピング失敗！")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("failed to get html: %s", err)
		return nil, errors.New("スクレイピング失敗！")
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Printf("failed to get html: %s", err)
		return nil, errors.New("スクレイピング失敗！")
	}

	// ランキングを取得（１位～１０位）
	results := []elegaku.Rank{}
	results = append(results, oneToThree(doc)...)
	results = append(results, fourToFive(doc)...)
	results = append(results, sixToTen(doc)...)

	return results, nil
}

// １～３位の情報を取得
func oneToThree(doc *goquery.Document) []elegaku.Rank {
	results := []elegaku.Rank{}

	doc.Find("#one_three").Find("#rank_com").Each(func(i int, sGirl *goquery.Selection) {
		// GirlId・順位の取得
		g, _ := sGirl.Find("div.g_image > a").Attr("href")
		r, _ := sGirl.Find("span").Attr("class")

		// 初期化・セット
		rank := elegaku.Rank{}
		rank.GirlId = regexp.MustCompile("[^0-9]").ReplaceAllString(g, "")
		rank.Rank, _ = strconv.Atoi(regexp.MustCompile("[^0-9]").ReplaceAllString(r, ""))
		rank.CreateDatetime = elegaku.GetTimestamp()
		rank.UpdateDatetime = elegaku.GetTimestamp()
		results = append(results, rank)
	})

	return results
}

// ４～５位の情報を取得
func fourToFive(doc *goquery.Document) []elegaku.Rank {
	results := []elegaku.Rank{}

	doc.Find("#four_five").Find("#rank_com").Each(func(i int, sGirl *goquery.Selection) {
		// GirlId・順位の取得
		g, _ := sGirl.Find("div.g_image > a").Attr("href")
		r, _ := sGirl.Find("span").Attr("class")

		// 初期化・セット
		rank := elegaku.Rank{}
		rank.GirlId = regexp.MustCompile("[^0-9]").ReplaceAllString(g, "")
		rank.Rank, _ = strconv.Atoi(regexp.MustCompile("[^0-9]").ReplaceAllString(r, ""))
		rank.CreateDatetime = elegaku.GetTimestamp()
		rank.UpdateDatetime = elegaku.GetTimestamp()
		results = append(results, rank)
	})

	return results
}

// ６～１０位の情報を取得
func sixToTen(doc *goquery.Document) []elegaku.Rank {
	results := []elegaku.Rank{}

	doc.Find("#castBox").Find("#companion_box").Each(func(i int, sGirl *goquery.Selection) {
		// GirlIdの取得
		g, _ := sGirl.Find("div.g_image > a").Attr("href")

		// 順位の取得
		// ６～１０位だけclass名がすべて06になっていたので画像パスから特定
		// 右の画像のパスから特定する//cdn1.fu-kakumei.com/69/pc_bak/images/rank/no10.png
		r1, _ := sGirl.Find("span.ranking06 > img").Attr("src")
		r2 := regexp.MustCompile(".*no").ReplaceAllString(r1, "")
		r3 := regexp.MustCompile("\\.png").ReplaceAllString(r2, "")

		// 初期化・セット
		rank := elegaku.Rank{}
		rank.GirlId = regexp.MustCompile("[^0-9]").ReplaceAllString(g, "")
		rank.Rank, _ = strconv.Atoi(r3)
		rank.CreateDatetime = elegaku.GetTimestamp()
		rank.UpdateDatetime = elegaku.GetTimestamp()
		results = append(results, rank)
	})

	return results
}
