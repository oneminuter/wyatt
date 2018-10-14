package model

//评论
type Comment struct {
	TableModel
	UserId  int64  `json:"userId"`                            //发送者用户id
	Content string `json:"content" gorm:"type:varchar(5000)"` //评论内容
}
