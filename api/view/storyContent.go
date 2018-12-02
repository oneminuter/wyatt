package view

import (
	"errors"
	"fmt"
	"wyatt/api/constant"
	"wyatt/api/model"
	"wyatt/util"
)

type StoryContent struct {
	SCId      string  `json:"scId"`      //内容流水号
	Type      int     `json:"type"`      //1 角色对白，2 旁白
	RolerId   string  `json:"rolerId"`   //角色流水号
	AvatarUrl string  `json:"avatarUrl"` //角色头像
	Nickname  string  `json:"nickname"`  //角色昵称
	Sex       int     `json:"sex"`       //角色性别 0 未知， 1 男， 2 女
	Context   string  `json:"context"`   //内容
	Order     float64 `json:"order"`     //权重
}

func (sc *StoryContent) HandlerRespList(cList []model.StoryContent, rMap map[int64]model.StoryRole) []StoryContent {
	var (
		list = make([]StoryContent, 0, len(cList))
		role model.StoryRole
		ok   bool
	)
	for _, v := range cList {
		role, ok = rMap[v.RoleId]
		if !ok {
			util.LoggerError(errors.New("roler info not exist"))
			role.AvatarUrl = ""
			role.Nickname = ""
			role.ID = 0
			role.Sex = 0
		}

		list = append(list, StoryContent{
			SCId:      fmt.Sprintf("%s.%d.%d", constant.SC, v.ID, v.FlowId),
			Type:      v.Type,
			RolerId:   fmt.Sprintf("%s.%d.%d", constant.R, role.ID, role.FlowId),
			AvatarUrl: role.AvatarUrl,
			Nickname:  role.Nickname,
			Sex:       role.Sex,
			Context:   v.Context,
			Order:     v.Order,
		})
	}
	return list
}
