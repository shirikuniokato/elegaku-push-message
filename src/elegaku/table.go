package elegaku

// table_name
const (
	TBLNM_GIRLS    string = "girls"
	TBLNM_NEW_FACE string = "new_face"
	TBLNM_RANK     string = "rank"
	TBLNM_USERS    string = "users"
	// scheculeのテーブル名はyyyyMMddとなるためここでは定義しない
)

// girlsテーブル
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

// girlsカラム
const (
	G_GIRL_ID          string = "girl_id"
	G_NAME             string = "name"
	G_AGE              string = "age"
	G_THREE_SIZE       string = "three_size"
	G_CATCH_COPY       string = "catch_copy"
	G_IMAGE            string = "image"
	G_CREATE_DATE_TIME string = "create_datetime"
	G_UPDATE_DATE_TIME string = "update_datetime"
)

// new_faceテーブル
type NewFace struct {
	GirlId         string `dynamo:"girl_id,hash"`
	CreateDatetime string `dynamo:"create_datetime"`
	UpdateDatetime string `dynamo:"update_datetime"`
}

// new_faceカラム
const (
	N_GIRL_ID          string = "girl_id"
	N_CREATE_DATE_TIME string = "create_datetime"
	N_UPDATE_DATE_TIME string = "update_datetime"
)

// rankテーブル
type Rank struct {
	Rank           int    `dynamo:"rank,hash"`
	GirlId         string `dynamo:"girl_id"`
	CreateDatetime string `dynamo:"create_datetime"`
	UpdateDatetime string `dynamo:"update_datetime"`
}

// rankカラム
const (
	R_RANK             string = "rank"
	R_GIRL_ID          string = "girl_id"
	R_CREATE_DATE_TIME string = "create_datetime"
	R_UPDATE_DATE_TIME string = "update_datetime"
)

// scheduleテーブル
type Schedule struct {
	GirlId         string `dynamo:"girl_id,hash"`
	Time           string `dynamo:"time"`
	NoticeFlg      int    `dynamo:"notice_flg"`
	CreateDatetime string `dynamo:"create_datetime"`
	UpdateDatetime string `dynamo:"update_datetime"`
}

// scheduleカラム
const (
	S_GIRL_ID          string = "girl_id"
	S_TIME             string = "time"
	S_NOTICE_FLG       string = "notice_flg"
	S_CREATE_DATE_TIME string = "create_datetime"
	S_UPDATE_DATE_TIME string = "update_datetime"
)

// usersテーブル
type User struct {
	UserId          string   `dynamo:"user_id,hash"`
	UserName        string   `dynamo:"user_name"`
	FavoriteGirlIds []string `dynamo:"favorite_girl_ids"`
}

// usersカラム
const (
	U_USER_ID           string = "user_id"
	U_USER_NAME         string = "user_name"
	U_FAVORITE_GIRL_IDS string = "favorite_girl_ids"
)
