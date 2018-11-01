package model

import "time"

//内容举报
type TipOff struct {
	TableModel
	UserId       int64  `json:"userId"`       //用户id
	SourceFlowId string `json:"sourceFlowId"` //举报内容完整流水号
	Status       int    `json:"status"`       //处理状态
	Remark       string `json:"remark"`       //处理备注
}

func (bc *TipOff) BeforeCreate() (err error) {
	bc.FlowId = time.Now().Unix()
	return
}
