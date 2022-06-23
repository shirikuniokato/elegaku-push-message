package elegaku

type Girl struct {
	GirlId         string `dynamo:"girl_id,hash"`
	Name           string `dynamo:"name"`
	Age            int    `dynamo:"age"`
	ThreeSize      string `dynamo:"three_size"`
	CatchCopy      string `dynamo:"catch_copy"`
	Image          string `dynamo:"image"`
	CreateDatetime string `dynamo:"create_datetime"`
	UpdateDatetime string `dynamo:"update_datetime"`
}

type NewFace struct {
	GirlId         string `dynamo:"girl_id,hash"`
	CreateDatetime string `dynamo:"create_datetime"`
	UpdateDatetime string `dynamo:"update_datetime"`
}

type Rank struct {
	Rank           int    `dynamo:"rank,hash"`
	GirlId         string `dynamo:"girl_id"`
	CreateDatetime string `dynamo:"create_datetime"`
	UpdateDatetime string `dynamo:"update_datetime"`
}

type Schedule struct {
	GirlId         string `dynamo:"girl_id,hash"`
	Time           string `dynamo:"time"`
	NoticeFlg      int    `dynamo:"notice_flg"`
	CreateDatetime string `dynamo:"create_datetime"`
	UpdateDatetime string `dynamo:"update_datetime"`
}
