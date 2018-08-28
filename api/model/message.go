package model

type Message struct {
	TableModel
	UserId  int64  `json:"userId"`  //用户id
	MsgType string `json:"msgType"` //消息类型， 系统消息，自定义消息
}
