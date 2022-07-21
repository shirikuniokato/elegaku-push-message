package elegaku

import "time"

// 日付フォーマット
const (
	YMD_FMT         = "20060102"
	MD_FMT          = "0102"
	ELEGAKU_YMD_FMT = "y/2006/MM/01/dd/02"
)

// contains(文字列型)
func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// タイムスタンプ取得（JST)
func GetTimeJst() time.Time {
	jst, _ := time.LoadLocation("Asia/Tokyo")
	return time.Now().In(jst)
}

// 曜日取得
func GetYoubi(t time.Time) string {
	YOUBI := []string{"日", "月", "火", "水", "木", "金", "土"}
	return YOUBI[t.Weekday()]
}
