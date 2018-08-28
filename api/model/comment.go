package model

type Comment struct {
	TableModel
	UserId  int64  `json:"userId"`  //发送者用户id
	Content string `json:"content"` //评论内容
}
