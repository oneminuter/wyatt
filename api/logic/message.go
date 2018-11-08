package logic

import (
	"wyatt/api/constant"
	"wyatt/api/model"
	"wyatt/api/view"
	"wyatt/util"
)

type Message struct {
	MId string `json:"mId" form:"mId" binding:"required"` //消息完整流水号
}

type MessageList struct {
	Page  int `json:"page" form:"page"`
	Limit int `json:"limit" form:"limit"`
}

func (ml *MessageList) List(userId int64) interface{} {
	//查询消息列表
	var mm model.Message
	messages, err := mm.QueryList("*", ml.Page, ml.Limit, "user_id = ?", userId)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.QueryDBEmptyErr)
	}

	var vm view.Message
	respData := vm.HandleRespList(messages)

	return view.SetRespData(respData)
}

func (m *Message) Detail(userId int64) interface{} {
	_, TableID, timestamp, err := util.SplitFlowNumber(m.MId)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.IncorrectFlowNumber)
	}

	//查询消息
	var mm model.Message
	err = mm.QueryOne("*", "id = ? AND flow_id = ?", TableID, timestamp)
	if err != nil {
		util.LoggerError(err)
		return view.CheckMysqlErr(err)
	}

	//判断是够有权限
	if userId != mm.UserId {
		return view.SetErr(constant.NoAuth)
	}

	var vm view.Message
	vm.HandleRespDetail(mm)
	return view.SetRespData(vm)
}

func (m *Message) Delete(userId int64) interface{} {
	_, TableID, timestamp, err := util.SplitFlowNumber(m.MId)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.IncorrectFlowNumber)
	}

	//查询消息
	var mm model.Message
	err = mm.QueryOne("*", "id = ? AND flow_id = ?", TableID, timestamp)
	if err != nil {
		util.LoggerError(err)
		return view.CheckMysqlErr(err)
	}

	//判断是够有权限
	if userId != mm.UserId {
		return view.SetErr(constant.NoAuth)
	}

	err = mm.Delete("id = ?", mm.ID)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.DeleteErr)
	}
	return view.SetErr(constant.Success)
}
