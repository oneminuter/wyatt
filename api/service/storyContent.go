package service

import (
	"wyatt/api/model"
)

type StoryContent struct{}

//故事内容详情字段
type Content struct {
	RoleId  string `json:"roleId" form:"roleId"`   //角色流水号id, 旁白为空
	Type    string `json:"type" form:"type"`       //类型 1 角色对白，2 旁白
	Context string `json:"context" form:"context"` //内容
}

//从故事详情中提取角色id数组
func (*StoryContent) GetRoleIdList(cList []model.StoryContent) []int64 {
	var list = make([]int64, 0)
	for _, v := range cList {
		list = append(list, v.RoleId)
	}
	return list
}
