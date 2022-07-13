package elegaku

import "time"

const YMD_FMT = "20060102"

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
