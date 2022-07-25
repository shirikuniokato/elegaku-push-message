package elegaku

import (
	"fmt"
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
)

// ポストバック
func Postback(bot *linebot.Client, event *linebot.Event) {
	// ポストバックデータ
	data := event.Postback.Data

	// ポストバックデータを基に返信内容を判定する
	if strings.Contains(data, PostbackTypeRegister) {
		// 通知登録ボタン押下
		if strings.Contains(data, PostbackTypeAdd) {
			// お気に入り登録
			notificationAdd(bot, event)
		} else if strings.Contains(data, PostbackTypeRemove) {
			// お気に入り解除
			notificationRemove(bot, event)
		} else {
			// お気に入り登録リスト取得
			notificationList(bot, event)
		}
		return
	} else if strings.Contains(data, PostbackTypeSchedule) {
		// 出勤予定表ボタン押下時
		schedule(bot, event)
	} else if strings.Contains(data, PostbackTypeLocation) {
		// アクセスボタン押下時
		location(bot, event)
	} else if strings.Contains(data, PostbackTypeSystem) {
		// 料金表ボタン押下時
		system(bot, event)
	} else if strings.Contains(data, PostbackTypeRank) {
		// ランキングボタン押下時
		rank(bot, event)
	} else if strings.Contains(data, PostbackTypeNewFace) {
		// 新入生紹介ボタン押下時
		newFace(bot, event)
	} else if strings.Contains(data, PostbackTypeVideo) {
		// 動画一覧ボタン押下時
		video(bot, event)
	} else if strings.Contains(data, PostbackTypeMenuSwitch) {
		// メニュー切り替えボタン押下時
		// 何もしない
	} else {
		// それ以外の場合
		// 何もしない
	}
}

// お気に入り登録用リストを送信
func notificationList(bot *linebot.Client, event *linebot.Event) {
	jsonData := []byte(`{
		"type": 'carousel',
		contents: [
			{
				type: 'bubble',
				header: {
					type: 'box',
					layout: 'vertical',
					contents: [
					  {
						type: 'text',
						text: 'お気に入り登録',
						align: 'center',
						size: 'xl',
						weight: 'bold',
						color: '#F0F0F0',
					  },
					],
					backgroundColor: '#84dcfd',
				  },
				body: {
					type: 'box',
					layout: 'vertical',
					contents: [
						{
							type: 'separator',
						},
						{
							type: 'box',
							layout: 'horizontal',
							contents: [
							  {
								type: 'image',
								url: 'https://cdn-fu-kakumei.com/image/c589c70bcc1c290e/0/0/.api',
								align: 'start',
								size: 'xs',
								flex: 1,
							  },
							  {
								type: 'box',
								layout: 'vertical',
								contents: [
								  {
									type: 'text',
									text: 'girlNameAndAge',
									weight: 'bold',
									action: {
									  type: 'uri',
									  label: 'action',
									  uri:
										'https://www.elegaku.com/profile/top/castCode/260317/',
									},
									color: '#3db3e9',
								  },
								  {
									type: 'text',
									text: 'threeSize',
									size: 'xs',
								  },
								  {
									type: 'text',
									text: 'catchCopy',
									size: 'xs',
								  },
								],
								flex: 2,
							  },
							  {
								type: 'button',
								action: {
								  type: 'postback',
								  label: '解除',
								  data: 'register:remove=girlId',
								},
								flex: 1,
								position: 'relative',
								gravity: 'center',
								style: 'primary',
								color: '#F30100',
								adjustMode: 'shrink-to-fit',
							  }
							],
							margin: 'none',
							paddingTop: 'sm',
						  }
	
					]
				}
			}
		]
	  }
	`)

	container, err := linebot.UnmarshalFlexMessageJSON(jsonData)
	if err != nil {
		// 正しくUnmarshalできないinvalidなJSONであればerrが返る
		message := linebot.NewTextMessage(err.Error())
		bot.ReplyMessage(event.ReplyToken, message).Do()
	} else {
		message := linebot.NewFlexMessage("alt text", container)
		bot.ReplyMessage(event.ReplyToken, message).Do()
	}
}

// お気に入り登録し、結果を送信
func notificationAdd(bot *linebot.Client, event *linebot.Event) {
	textMessage := linebot.NewTextMessage("未実装")
	bot.ReplyMessage(event.ReplyToken, textMessage).Do()
}

// お気に入り解除し、結果を送信
func notificationRemove(bot *linebot.Client, event *linebot.Event) {
	textMessage := linebot.NewTextMessage("未実装")
	bot.ReplyMessage(event.ReplyToken, textMessage).Do()
}

// 直近１週間分のボタンを送信（クイックリプライ）
func schedule(bot *linebot.Client, event *linebot.Event) {
	bot.ReplyMessage(event.ReplyToken, createQuickReplyItems()).Do()
}

// 直近１週間分の出勤確認ボタンを作成
func createQuickReplyItems() linebot.SendingMessage {
	items := []linebot.QuickReplyButton{}
	// １週間分の加算値
	w := []int{0, 1, 2, 3, 4, 5, 6}
	for _, v := range w {
		t := GetTimeJst().AddDate(0, 0, v)

		// クイックリプライのボタンに表示する文字列を生成
		var label string
		if v == 0 {
			label = "本日"
		} else if v == 1 {
			label = "明日"
		} else {
			label = fmt.Sprintf("%s(%s)", t.Format(MD_FMT), GetYoubi(t))
		}

		action := NewURIAction(label, fmt.Sprintf("https://www.elegaku.com/cast/schedule/%s", t.Format(ELEGAKU_YMD_FMT)))
		items = append(items, *linebot.NewQuickReplyButton("", action))
	}
	q := linebot.NewQuickReplyItems(&items[0], &items[1], &items[2], &items[3], &items[4], &items[5], &items[6])
	return linebot.NewTextMessage("選択した日付の出勤情報を確認します。").WithQuickReplies(q)
}

// 位置情報を送信
func location(bot *linebot.Client, event *linebot.Event) {
	message := linebot.NewLocationMessage("エレガンス学院", "神奈川県川崎市川崎区堀之内町７−８", 35.533641839733406, 139.70597139350963)
	// 位置情報を送信
	bot.ReplyMessage(event.ReplyToken, message).Do()
}

// 料金表を送信
func system(bot *linebot.Client, event *linebot.Event) {
	message := linebot.NewImageMessage(SystemImageURL, SystemImageURL)
	bot.ReplyMessage(event.ReplyToken, message).Do()
}

// ランキング情報を送信
func rank(bot *linebot.Client, event *linebot.Event) {
	textMessage := linebot.NewTextMessage("未実装")
	bot.ReplyMessage(event.ReplyToken, textMessage).Do()
}

// 新入生一覧を送信
func newFace(bot *linebot.Client, event *linebot.Event) {
	textMessage := linebot.NewTextMessage("未実装")
	bot.ReplyMessage(event.ReplyToken, textMessage).Do()
}

// 動画一覧を送信
func video(bot *linebot.Client, event *linebot.Event) {
	textMessage := linebot.NewTextMessage("未実装")
	bot.ReplyMessage(event.ReplyToken, textMessage).Do()
}
