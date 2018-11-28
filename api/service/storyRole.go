package service

import (
	"wyatt/api/model"
	"wyatt/util"
)

type StoryRole struct{}

//判断是否是角色的创建者
func (sr *StoryRole) IsCreator(rolerId, userId int64) bool {
	var msr model.StoryRole
	count, err := msr.QueryCount("rolerId = ? AND creator_id = ?", rolerId, userId)
	if err != nil {
		util.LoggerError(err)
		return false
	}
	if count > 0 {
		return true
	}
	return false
}

//角色列表转map id:role
func (sr *StoryRole) GetRoleMap(rList []model.StoryRole) map[int64]model.StoryRole {
	var m = make(map[int64]model.StoryRole)
	for _, v := range rList {
		m[v.ID] = v
	}
	return m
}
