package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Schedule struct {
	GirlId    string `json:"girl_id"`
	Time      string `json:"time"`
	NoticeFlg int    `json:"notice_flg"`
}

func main() {
	// １週間分の加算値
	week := []int{0, 1, 2, 3, 4, 5, 6}

	t := time.Now()
	for _, v := range week {
		updateSchedule(t.AddDate(0, 0, v).Format("y/2006/MM/01/dd/02"))
	}
}

// ランキング更新
func updateSchedule(date string) {
	webPage := ("https://www.elegaku.com/cast/schedule/" + date)
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

	schedules := []Schedule{}

	doc.Find("#companion_box").Each(func(i int, sGirl *goquery.Selection) {
		s := Schedule{}

		// GirlIdの取得
		g, _ := sGirl.Find("div.g_image > a").Attr("href")

		// 初期化・セット
		s.GirlId = regexp.MustCompile("[^0-9]").ReplaceAllString(g, "")
		s.Time = strings.TrimSpace(sGirl.Find(".time").Text())
		// TODO■既に登録済みの場合はフラグを引き継ぐ必要がある
		s.NoticeFlg = 0

		schedules = append(schedules, s)
	})

	fmt.Println(schedules)
	// TODO■DynamoDBへの登録・更新処理が必要
	// TODO■テーブル名は日付（yyyyMMdd）とする
}
