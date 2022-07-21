package elegaku

import (
	"encoding/json"

	"github.com/line/line-bot-sdk-go/linebot"
)

// ↓ ドキュメント
// https://developers.line.biz/ja/docs/messaging-api/

// MessaginAPIのリクエスト用のURL
const (
	LINE_URL_REPLY     = "https://api.line.me/v2/bot/message/reply"     // リプライ用のURL
	LINE_URL_PROFILE   = "https://api.line.me/v2/bot/profile"           // ユーザ情報取得用のURL
	LINE_URL_PUSH      = "https://api.line.me/v2/bot/message/push"      // プッシュ通知送信用のURL
	LINE_URL_MULTICAST = "https://api.line.me/v2/bot/message/multicast" // プッシュ通知送信（一括送信）用のURL
)

// Webhook
type WebHook struct {
	Destination string          `json:"destination"`
	Events      []linebot.Event `json:"events"`
}

// // ポストバックイベントタイプ
const (
	PostbackTypeRegister     = "register"    // お気に入り登録リスト取得
	PostbackTypeAdd          = "add"         // お気に入り登録
	PostbackTypeRemove       = "remove"      // お気に入り解除
	PostbackTypeSchedule     = "schedule"    // 出勤情報一覧取得
	PostbackTypeScheduleDate = "date"        // 出勤情報取得（不要かも）
	PostbackTypeLocation     = "location"    // 位置情報取得
	PostbackTypeSystem       = "system"      // 料金表取得
	PostbackTypeRank         = "ranking"     // ランキング取得
	PostbackTypeNewFace      = "newface"     // 新入生取得
	PostbackTypeVideo        = "video"       // ビデオ取得
	PostbackTypeMenuSwitch   = "menu_switch" // メニュー切り替え
)

const SystemImageURL = "https://cdn1.fu-kakumei.com/69/pc_bak/images/system/system1.jpg" // 料金表の画像URL

// クイックリプライでURIアクションを設定する場合に、line-bot-sdk-go/linebot/action.go#URIActionではインターフェースの型違いで使えないため
// クイックリプライ用に合わせたインターフェールの型で自前で定義する
type URIAction struct {
	Label  string
	URI    string
	AltURI *URIActionAltURI
}
type URIActionAltURI struct {
	Desktop string `json:"desktop"`
}

func (a *URIAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type   linebot.ActionType `json:"type"`
		Label  string             `json:"label,omitempty"`
		URI    string             `json:"uri"`
		AltURI *URIActionAltURI   `json:"altUri,omitempty"`
	}{
		Type:   linebot.ActionTypeURI,
		Label:  a.Label,
		URI:    a.URI,
		AltURI: a.AltURI,
	})
}
func (*URIAction) QuickReplyAction() {}
func NewURIAction(label, uri string) *URIAction {
	return &URIAction{
		Label: label,
		URI:   uri,
	}
}
