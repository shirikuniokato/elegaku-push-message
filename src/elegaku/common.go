package elegaku

import "time"

func GetTimestamp() time {
	// 現在時刻をyyyy-MM-dd hh:mm:ss形式で取得
	return time.Now().Format("2006-01-02 15:04:05")
}
