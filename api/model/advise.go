package model

//用户建议
type Advise struct {
	TableModel
	UserId  int64  `json:"userId"`  //用户id
	Phone   string `json:"phone"`   //用户电话
	Email   string `json:"email"`   //邮箱
	Content string `json:"content"` //建议内容
	Status  int    `json:"status"`  //处理状态， 默认 0, 1 已看
	Remark  string `json:"remark"`  //备注
}
