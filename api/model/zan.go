package model

type Zan struct {
	TableModel
	UserId  int64  `json:"userId"`  //触发者用户id
	Table   string `json:"table"`   //获赞表
	TableId int64  `json:"tableId"` //获赞表id
}
