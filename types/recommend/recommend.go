package recommendType

import "time"

type Record struct {
	Id          int       `xorm:"id"`
	UserId1     int       `xorm:"user_id_1"`
	UserId2     int       `xorm:"user_id_2"`
	PostId      int64     `xorm:"post_id"`
	CreatedTime time.Time `xorm:"created_time"`
	Status      int       `xorm:"status"`
}

type PostWeiboReq struct {
	Req []Req `json:"req"`
}

type Req struct {
	UserId1 int `json:"user_id_1"`
	UserId2 int `json:"user_id_2"`
}

type SuggestResp struct {
	UserId int `json:"user_id"`
}
