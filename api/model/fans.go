package model

//关注粉丝
type Fans struct {
	TableModel
	UserId int64 `json:"userId"`
	FansUserId int64 `json:"fansUserId"`
}