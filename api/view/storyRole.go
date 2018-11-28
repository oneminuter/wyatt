package view

import (
	"fmt"
	"wyatt/api/constant"
	"wyatt/api/model"
)

type StoryRole struct {
	RolerId   string `json:"rolerId"`   //角色流水号
	AvatarUrl string `json:"avatarUrl"` //角色头像
	Nickname  string `json:"nickname"`  //角色昵称
	Sex       int    `json:"sex"`       //角色性别 0 未知， 1 男， 2 女
	Introduce string `json:"introduce"` //角色介绍
}

func (*StoryRole) HandlerRespList(roles []model.StoryRole) []StoryRole {
	var list = make([]StoryRole, 0, len(roles))
	for _, v := range roles {
		list = append(list, StoryRole{
			RolerId:   fmt.Sprintf("%s.%d.%d", constant.R, v.ID, v.FlowId),
			AvatarUrl: v.AvatarUrl,
			Nickname:  v.Nickname,
			Sex:       v.Sex,
			Introduce: v.Introduce,
		})
	}
	return list
}

func (sr *StoryRole) HandlerRespInfo(r model.StoryRole) {
	sr.RolerId = fmt.Sprintf("%s.%d.%d", constant.R, r.ID, r.FlowId)
	sr.AvatarUrl = r.AvatarUrl
	sr.Nickname = r.Nickname
	sr.Sex = r.Sex
	sr.Introduce = r.Introduce
}
