package service

import (
	"wyatt/api/model"
	"wyatt/util"
)

type StoryRole struct{}

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
