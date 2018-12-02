package service

import (
	"wyatt/api/model"
	"wyatt/util"
)

type Story struct{}

func (s *Story) IsCreator(storyId, userId int64) bool {
	var ms model.Story
	count, err := ms.QueryCount("id = ? AND author_id = ?", storyId, userId)
	if err != nil {
		util.LoggerError(err)
		return false
	}
	if 1 > count {
		return false
	}
	return true
}
