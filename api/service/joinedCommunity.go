package service

import "wyatt/api/model"

type JoinedCommunity struct{}

func (JoinedCommunity) GetJoinedIdArr(list []model.JoinedCommunity) []int64 {
	var arr = make([]int64, 0, len(list))
	for _, v := range list {
		arr = append(arr, v.CommunityId)
	}
	return arr
}
