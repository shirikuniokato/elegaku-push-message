package elegaku

// プッシュ通知用のFlexMessage
// １番目：出勤日
// ２番目：女の子の画像URL
// ３番目：女の子の名前・年齢
// ４番目：GirlId
// ５番目：threeSize
// ６番目：catchCopy
// ７番目：GirlId
const PUSH_MESSAGE_FOMAT = `{
	"type": "bubble",
	"hero": {
		"type": "box",
		"layout": "vertical",
		"contents": [{
				"type": "box",
				"layout": "vertical",
				"contents": [{
						"type": "text",
						"text": "出勤情報確定！",
						"size": "lg",
						"weight": "bold",
						"align": "center",
						"gravity": "center",
						"color": "#ffffff"
					},
					{
						"type": "text",
						"text": "%s",
						"color": "#ffffff",
						"weight": "bold"
					}
				],
				"backgroundColor": "#84dcfd",
				"justifyContent": "center",
				"alignItems": "center",
				"margin": "none",
				"paddingAll": "md"
			},
			{
				"type": "image",
				"url": "%s",
				"aspectMode": "cover",
				"aspectRatio": "185:247",
				"size": "full"
			}
		]
	},
	"body": {
		"type": "box",
		"layout": "vertical",
		"contents": [{
				"type": "text",
				"text": "%s",
				"size": "xl",
				"color": "#3db3e9",
				"action": {
					"type": "uri",
					"label": "action",
					"uri": "https://www.elegaku.com/profile/top/castCode/%s/"
				}
			},
			{
				"type": "text",
				"text": "%s"
			},
			{
				"type": "text",
				"text": "%s"
			}
		],
		"alignItems": "center"
	},
	"footer": {
		"type": "box",
		"layout": "horizontal",
		"spacing": "sm",
		"contents": [{
				"type": "button",
				"style": "primary",
				"height": "sm",
				"action": {
					"type": "uri",
					"label": "ウェブサイト",
					"uri": "https://www.elegaku.com/profile/top/castCode/%s/"
				}
			},
			{
				"type": "button",
				"action": {
					"type": "uri",
					"uri": "tel:0442465322",
					"label": "TEL"
				},
				"style": "primary",
				"height": "sm",
				"color": "#f3892b"
			}
		],
		"flex": 0
	}
}`
