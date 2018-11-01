package model

import "time"

//关注粉丝
type Fans struct {
	TableModel
	UserId     int64 `json:"userId"`
	FansUserId int64 `json:"fansUserId"`
}

func (bc *Fans) BeforeCreate() (err error) {
	bc.FlowId = time.Now().Unix()
	return
}
