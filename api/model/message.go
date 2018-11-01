package model

import "time"

//消息中心
type Message struct {
	TableModel

	UserId  int64  `json:"userId"`  //用户id
	MsgType string `json:"msgType"` //消息类型， 系统消息，自定义消息，可以根据该字段来判断是不是定向消息
	Content string `json:"content"` //消息内容
}

func (bc *Message) BeforeCreate() (err error) {
	bc.FlowId = time.Now().Unix()
	return
}
