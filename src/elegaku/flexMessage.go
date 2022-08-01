package elegaku

const PUSH_MESSAGE = `{
	"type": "bubble",
	"hero": {
	  "type": "box",
	  "layout": "vertical",
	  "contents": [
		{
		  "type": "box",
		  "layout": "vertical",
		  "contents": [
			{
			  "type": "text",
			  "text": "出勤情報確定！",
			  "size": "lg",
			  "weight": "bold",
			  "align": "center",
			  "gravity": "center",
			  "color": "#ffffff",
			},
			{
			  "type": "text",
			  "text": systemDateTime,
			  "color": "#ffffff",
			  "weight": "bold",
			},
		  ],
		  "backgroundColor": "#84dcfd",
		  "justifyContent": "center",
		  "alignItems": "center",
		  "margin": "none",
		  "paddingAll": "md",
		},
		{
		  "type": "image",
		  "url": image,
		  "aspectMode": "cover",
		  "aspectRatio": "185:247",
		  "size": "full",
		},
	  ],
	},
	"body": {
	  "type": "box",
	  "layout": "vertical",
	  "contents": [
		{
		  "type": "text",
		  "text": girlNameAndAge,
		  "size": "xl",
		  "color": "#3db3e9",
		  "action": {
			"type": "uri",
			"label": "action",
			"uri": "https://www.elegaku.com/profile/top/castCode/" + girlId + "/",
		  },
		},
		{
		  "type": "text",
		  "text": threeSize,
		},
		{
		  "type": "text",
		  "text": catchCopy,
		},
	  ],
	  "alignItems": "center",
	},
	"footer": {
	  "type": "box",
	  "layout": "horizontal",
	  "spacing": "sm",
	  "contents": [
		{
		  "type": "button",
		  "style": "primary",
		  "height": "sm",
		  "action": {
			"type": "uri",
			"label": "ウェブサイト",
			"uri": "https://www.elegaku.com/profile/top/castCode/" + girlId + "/",
		  },
		},
		{
		  "type": "button",
		  "action": {
			"type": "uri",
			"uri": "tel:0442465322",
			"label": "TEL",
		  },
		  "style": "primary",
		  "height": "sm",
		  "color": "#f3892b",
		},
	  ],
	  "flex": 0,
	}
  }`
