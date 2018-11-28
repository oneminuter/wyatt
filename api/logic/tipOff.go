package logic

import (
	"wyatt/api/constant"
	"wyatt/api/model"
	"wyatt/api/view"
	"wyatt/util"
)

type TipOff struct {
	SourceFlowId string `json:"sourceFlowId" form:"sourceFlowId" binding:"required"`
	Reason       string `json:"reason" form:"reason"`
}

//添加举报
func (t *TipOff) Add(userId int64) interface{} {
	//分割所举报内容流水号
	tableName, TableID, timestamp, err := util.SplitFlowNumber(t.SourceFlowId)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.IncorrectFlowNumber)
	}

	//验证流水号是否合法
	if !model.ValidateFlowId(tableName, TableID, timestamp) {
		return view.SetErr(constant.IncorrectFlowNumber)
	}

	//添加到数据库
	mto := model.TipOff{
		UserId:       userId,
		SourceFlowId: t.SourceFlowId,
		Reason:       t.Reason,
		Status:       0,
	}

	err = mto.Add()
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.AddErr)
	}
	return view.SetErr(constant.Success)
}
