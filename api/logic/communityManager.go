package logic

import (
	"strings"
	"wyatt/api/constant"
	"wyatt/api/model"
	"wyatt/api/service"
	"wyatt/api/view"
	"wyatt/util"
)

type CommunityManager struct {
	CId     int64 `json:"cId" form:"cId" binding:"required"`        //社区号
	Account int64 `json:"userId" form:"account" binding:"required"` //用户ID
}

//添加管理员
func (cm *CommunityManager) Add(creatorId int64) interface{} {
	//判断权限
	var sc service.Community
	if !sc.IsAdmin(cm.CId, creatorId) {
		return view.SetErr(constant.NoAuth)
	}

	//查询对应的用户id
	var mu model.User
	err := mu.QueryOne("*", "account = ?", cm.Account)
	if err != nil && strings.Contains(err.Error(), constant.MysqlNotHaveData) {
		return view.SetErr(constant.UserNotExist)
	}
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.AddErr)
	}

	//判断目标用户状态是否为被封禁
	if -1 == mu.Status {
		return view.SetErr(constant.AccountForbid)
	}

	//添加
	mcm := model.CommunityManager{
		CommunityId: cm.CId,
		UserId:      mu.ID,
		Role:        1,
	}
	err = mcm.Add()
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.AddErr)
	}
	return view.SetErr(constant.Success)
}

//删除管理员
func (cm *CommunityManager) Delete(creatorId int64) interface{} {
	//判断权限
	var sc service.Community
	if !sc.IsAdmin(cm.CId, creatorId) {
		return view.SetErr(constant.NoAuth)
	}

	//查询对应的用户id
	var mu model.User
	err := mu.QueryOne("*", "account = ?", cm.Account)
	if err != nil && strings.Contains(err.Error(), constant.MysqlNotHaveData) {
		return view.SetErr(constant.UserNotExist)
	}
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.DeleteErr)
	}

	var mcm model.CommunityManager
	err = mcm.Delete("community_id = ? AND user_id = ?", cm.CId, mu.ID)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.DeleteErr)
	}
	return view.SetErr(constant.Success)
}
