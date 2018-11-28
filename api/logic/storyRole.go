package logic

import (
	"log"
	"wyatt/api/constant"
	"wyatt/api/model"
	"wyatt/api/view"
	"wyatt/util"
)

type StoryRole struct {
	StoryId   string `json:"storyId" form:"storyId"  binding:"required"`  //故事流水号
	Nickname  string `json:"nickname" form:"nickname" binding:"required"` //角色昵称
	AvatarUrl string `json:"avatarUrl" form:"avatarUrl"`                  //头像
	Sex       int    `json:"sex" form:"sex"`                              //角色性别
	Introduce string `json:"introduce" form:"introduce"`                  //角色介绍
}

func (sr *StoryRole) Add(userId int64) interface{} {
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
