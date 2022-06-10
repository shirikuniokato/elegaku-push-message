package main

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

type Rank struct {
	Rank   int    `json:"rank"`
	GirlId string `json:"girl_id"`
}

// ランキング更新
func main() {
	webPage := ("https://www.elegaku.com/rank/")
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

	// ランキングを取得（１位～１０位）
	rank := []Rank{}
	rank = append(rank, one_three(doc)...)
	rank = append(rank, four_five(doc)...)
	rank = append(rank, six_ten(doc)...)

	// TODO■DynamoDBへの登録・更新処理が必要
}

// １～３位の情報を取得
func one_three(doc *goquery.Document) []Rank {
	results := []Rank{}

	doc.Find("#one_three").Find("#rank_com").Each(func(i int, sGirl *goquery.Selection) {
		// GirlId・順位の取得
		g, _ := sGirl.Find("div.g_image > a").Attr("href")
		r, _ := sGirl.Find("span").Attr("class")

		// 初期化・セット
		rank := Rank{}
		rank.GirlId = regexp.MustCompile("[^0-9]").ReplaceAllString(g, "")
		rank.Rank, _ = strconv.Atoi(regexp.MustCompile("[^0-9]").ReplaceAllString(r, ""))
		results = append(results, rank)
	})

	return results
}

// ４～５位の情報を取得
func four_five(doc *goquery.Document) []Rank {
	results := []Rank{}

	doc.Find("#four_five").Find("#rank_com").Each(func(i int, sGirl *goquery.Selection) {
		// GirlId・順位の取得
		g, _ := sGirl.Find("div.g_image > a").Attr("href")
		r, _ := sGirl.Find("span").Attr("class")

		// 初期化・セット
		rank := Rank{}
		rank.GirlId = regexp.MustCompile("[^0-9]").ReplaceAllString(g, "")
		rank.Rank, _ = strconv.Atoi(regexp.MustCompile("[^0-9]").ReplaceAllString(r, ""))
		results = append(results, rank)
	})

	return results
}

// ６～１０位の情報を取得
func six_ten(doc *goquery.Document) []Rank {
	results := []Rank{}

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
		rank := Rank{}
		rank.GirlId = regexp.MustCompile("[^0-9]").ReplaceAllString(g, "")
		rank.Rank, _ = strconv.Atoi(r3)
		results = append(results, rank)
	})

	return results
}
