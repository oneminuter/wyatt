package service

import "wyatt/api/model"

type Comment struct{}

//提取用户id数组
func (c *Comment) GetUserIDArr(mcList []model.Comment) []int64 {
	var list = make([]int64, 0)
	for _, v := range mcList {
		list = append(list, v.UserId)
	}
	return list
}
