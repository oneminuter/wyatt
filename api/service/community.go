package service

import "wyatt/api/model"

type Community struct{}

//提取取所有的社区id
func (Community) GetCommunityIdArr(list []model.Community) []int64 {
	arr := make([]int64, len(list))
	for _, v := range list {
		arr = append(arr, v.ID)
	}
	return arr
}

/*
	从加入设局列表中提取社区id:加入人数
	参数 model.JoinedCommunity 结构体数中，只有id 和 count 字段有值
*/
func (Community) GetCommunityJoinNumMap(list []model.JoinedCommunity) map[int64]int {
	m := make(map[int64]int)
	for _, v := range list {
		m[v.CommunityId] = v.Count
	}
	return m
}

func (Community) GetCommunityTopicNumMap(list []model.Topic) map[int64]int {
	m := make(map[int64]int)
	for _, v := range list {
		m[v.CommunityId] = v.Count
	}
	return m
}
