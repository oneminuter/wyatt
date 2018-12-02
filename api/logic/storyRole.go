package logic

import (
	"log"
	"strings"
	"wyatt/api/constant"
	"wyatt/api/model"
	"wyatt/api/service"
	"wyatt/api/view"
	"wyatt/util"
)

type StoryRoleAdd struct {
	StoryId   string `json:"storyId" form:"storyId"  binding:"required"`  //故事流水号
	Nickname  string `json:"nickname" form:"nickname" binding:"required"` //角色昵称
	AvatarUrl string `json:"avatarUrl" form:"avatarUrl"`                  //头像
	Sex       int    `json:"sex" form:"sex"`                              //角色性别
	Introduce string `json:"introduce" form:"introduce"`                  //角色介绍
}
type StoryRoleModify struct {
	RolerId   string `json:"rolerId" form:"rolerId"  binding:"required"`  //角色流水号
	Nickname  string `json:"nickname" form:"nickname" binding:"required"` //角色昵称
	AvatarUrl string `json:"avatarUrl" form:"avatarUrl"`                  //头像
	Sex       int    `json:"sex" form:"sex"`                              //角色性别
	Introduce string `json:"introduce" form:"introduce"`                  //角色介绍
}
type StoryRole struct {
	RolerId string `json:"rolerId" form:"rolerId"  binding:"required"` //角色流水号
}
type StoryRoleList struct {
	StoryId string `json:"storyId" form:"storyId"  binding:"required"` //故事流水号
	Page    int    `json:"page" form:"page"`                           //页码，从0开始，默认为0
	Limit   int    `json:"limit" form:"limit"`                         //查询条数, 最大查询20条
}

//添加角色
func (sr *StoryRoleAdd) Add(userId int64) interface{} {
	log.Println(sr.StoryId)
	_, TableID, _, err := util.SplitFlowNumber(sr.StoryId)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.IncorrectFlowNumber)
	}

	msr := model.StoryRole{
		AvatarUrl: sr.AvatarUrl,
		Nickname:  sr.Nickname,
		Sex:       0,
		Introduce: sr.Introduce,
		StoryId:   TableID,
		CreatorId: userId,
	}
	err = msr.Add()
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.AddErr)
	}
	return view.SetErr(constant.Success)
}

//修改角色信息
func (sr *StoryRoleModify) Modify(userId int64) interface{} {
	_, TableID, _, err := util.SplitFlowNumber(sr.RolerId)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.IncorrectFlowNumber)
	}
	//判断是否角色的创建者
	var ssr service.StoryRole
	if !ssr.IsCreator(TableID, userId) {
		return view.SetErr(constant.NoAuth)
	}

	//更新
	var (
		msr    model.StoryRole
		modify = make(map[string]interface{})
	)
	//昵称
	if "" != strings.TrimSpace(sr.Nickname) {
		modify["nickname"] = sr.Nickname
	}
	//头像
	if "" != strings.TrimSpace(sr.AvatarUrl) {
		modify["avatar_url"] = sr.AvatarUrl
	}
	//性别
	if -1 < sr.Sex || 3 > sr.Sex {
		modify["sex"] = sr.Sex
	}
	//介绍
	if "" != strings.TrimSpace(sr.Introduce) {
		modify["introduce"] = sr.Introduce
	}

	err = msr.Update(modify, "id = ?", TableID)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.ModifyErr)
	}

	return view.SetErr(constant.Success)
}

//故事的角色列表
func (sr *StoryRole) Delete(userId int64) interface{} {
	_, TableID, _, err := util.SplitFlowNumber(sr.RolerId)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.IncorrectFlowNumber)
	}

	//判断是否角色的创建者
	var ssr service.StoryRole
	if !ssr.IsCreator(TableID, userId) {
		return view.SetErr(constant.NoAuth)
	}

	//删除
	var msr model.StoryRole
	err = msr.Delete("id = ?", TableID)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.DeleteErr)
	}
	return view.SetErr(constant.Success)
}

func (sr *StoryRoleList) List() interface{} {
	_, TableID, _, err := util.SplitFlowNumber(sr.StoryId)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.IncorrectFlowNumber)
	}

	var msr model.StoryRole
	roles, err := msr.QueryList("*", sr.Page, sr.Limit, "story_id = ?", TableID)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.QueryDBErr)
	}

	var vsr view.StoryRole
	retData := vsr.HandlerRespList(roles)

	return view.SetRespData(retData)
}

//角色信息
func (sr *StoryRole) Info() interface{} {
	_, TableID, _, err := util.SplitFlowNumber(sr.RolerId)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.IncorrectFlowNumber)
	}

	var msr model.StoryRole
	err = msr.QueryOne("*", "id = ?", TableID)
	if err != nil {
		util.LoggerError(err)
		return view.CheckMysqlErr(err)
	}

	var vsr view.StoryRole
	vsr.HandlerRespInfo(msr)

	return view.SetRespData(vsr)
}
