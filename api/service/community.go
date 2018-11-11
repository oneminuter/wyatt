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

//话题列表转 社区id:话题数 map
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
	cId: 社区号，主键
	userId: 用户id
*/

func (Community) IsManager(cId int64, userId int64) bool {
	var (
		mc  model.Community
		mcm model.CommunityManager
	)
	err := mc.QueryOne("*", "id = ?", cId)
	if err != nil {
		util.LoggerError(err)
		return false
	}

	if mc.CreatorId == userId && 0 != userId {
		return true
	}

	managers, err := mcm.QueryList("*", 0, 999, "community_id = ?", mc.ID)
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
cId: 社区主键号
userId: 用户id
*/
func (Community) IsAdmin(cId int64, userId int64) bool {
	var mc model.Community
	count, err := mc.QueryCount("id = ? AND creator_id = ?", cId, userId)
	if err != nil {
		util.LoggerError(err)
		return false
	}
	return count > 0
}

//提取 社区主键id:社区 map
func (c *Community) GetCommunityMap(list []model.Community) map[int64]model.Community {
	var m = make(map[int64]model.Community)
	for _, v := range list {
		m[v.ID] = v
	}
	return m
}
