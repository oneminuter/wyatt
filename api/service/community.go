package service

import (
	"errors"
	"wyatt/api/model"
	"wyatt/util"
)

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

//保存logo， 返回logo 存储的 path
func (Community) SaveLogo(imgBase64 string) (string, error) {
	filename, err := util.SaveFile(imgBase64)
	if err != nil {
		util.LoggerError(err)
		return "", errors.New("Save file error")
	}
	return util.FilePath + filename, nil
}

/*
判断是不是管理员或者是不是创建者
	cId: 社区号 - 10位的时间错
	userId: 用户id
*/

func (Community) IsManager(cId int64, userId int64) bool {
	var (
		mc  model.Community
		mcm model.CommunityManager
	)
	err := mc.QueryOne("*", "c_id = ?", cId)
	if err != nil {
		util.LoggerError(err)
		return false
	}

	if mc.CreatorId == userId && 0 != userId {
		return true
	}

	managers, err := mcm.QueryList("*", "community_id = ?", mc.ID)
	if err != nil {
		util.LoggerError(err)
		return false
	}

	var isTrue = false
	for _, v := range managers {
		if v.UserId == userId {
			isTrue = true
			break
		}
	}
	return isTrue
}

/*
判断是否是创建者 - 删除社区，管理管理员
*/
func (Community) IsAdmin(cId int64, userId int64) bool {
	var mc model.Community
	count := mc.QueryCount("c_id = ? AND creator_id = ?", cId, userId)
	return count > 0
}
