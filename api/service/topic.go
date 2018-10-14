package service

import "wyatt/api/model"

type Topic struct{}

//从话题列表里提取创建者id列表
func (Topic) GetCreatorIdList(tList []model.Topic) []int64 {
	var list = make([]int64, 0, len(tList))
	for _, v := range tList {
		list = append(list, v.CreatorId)
	}
	return list
}
