package model

//社区
type Community struct {
	TableModel

	Name      string `json:"name"`      //社区名
	Desc      string `json:"desc"`      //社区描述
	CreatorId int64  `json:"creatorId"` //创建者id
}
