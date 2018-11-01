package model

import "time"

//赞记录
type Zan struct {
	TableModel
	UserId       int64  `json:"userId"`       //触发者用户id
	SourceFlowId string `json:"sourceFlowId"` //赞来源流水号，完整流水号
}

func (bc *Zan) BeforeCreate() (err error) {
	bc.FlowId = time.Now().Unix()
	return
}
