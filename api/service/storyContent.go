package service

import (
	"wyatt/api/model"
)

type StoryContent struct{}

func (*StoryContent) GetRoleIdList(cList []model.StoryContent) []int64 {
	var list = make([]int64, 0)
	for _, v := range cList {
		list = append(list, v.RoleId)
	}
	return list
}
