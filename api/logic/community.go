package logic

import (
	"wyatt/api/constant"
	"wyatt/api/model"
	"wyatt/api/service"
	"wyatt/api/view"
	"wyatt/util"
)

type Community struct{}

//查询所有状态为 1 的社区
func (c *Community) ListAll() interface{} {
	var m model.Community
	//获取状态=1(正常)的所有社区
	list, err := m.QueryList("*", "status = 1")
	if err != nil {
		util.LoggerError(err)
		return view.CheckMysqlErr(err)
	}

	//社区主键id 数组
	var sc service.Community
	communityIdArr := sc.GetCommunityIdArr(list)

	//获取各社区的加入人数
	var mjc model.JoinedCommunity
	joinList, err := mjc.QueryGrounp("community_id, count(*) count", "community_id", "community_id in (?)", communityIdArr)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.QueryDBErr)
	}

	// 转换为社区id:加入人数map
	joinNumMap := sc.GetCommunityJoinNumMap(joinList)

	//获取社区文章数
	var t model.Topic
	topicList, err := t.QueryGrounp("community_id, count(*) count", "community_id", "community_id in (?)", communityIdArr)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.QueryDBErr)
	}
	//装换为社区id:文章数map
	topicNumMap := sc.GetCommunityTopicNumMap(topicList)

	var vc view.Community
	resp := vc.RenderListAll(list, joinNumMap, topicNumMap)
	return view.SetRespData(resp)
}
