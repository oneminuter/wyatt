package logic

import (
	"wyatt/api/constant"
	"wyatt/api/model"
	"wyatt/api/service"
	"wyatt/api/view"
	"wyatt/util"
)

type JoinedCommunity struct {
	CId   string `json:"cId" form:"cId" binding:"required"`
	Page  int    `json:"page" form:"page"`   //页码，默认从0开始
	Limit int    `json:"limit" form:"limit"` //查询条数, 最大查询 constant.MAX_QUERY_COUNT
}

//我加入的社区列表
func (c *JoinedCommunity) MyList(userId int64) interface{} {

	//已加入的社区列表
	var mjc model.JoinedCommunity
	jlist, err := mjc.QueryList("*", "user_id = ?", userId)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.QueryDBErr)
	}

	//获得社区id数组
	var sjc service.JoinedCommunity
	joinedIdArr := sjc.GetJoinedIdArr(jlist)

	//获取社区信息
	var mc model.Community
	communities, err := mc.QueryList("*", c.Page, c.Limit, "status = 1 AND id in (?)", joinedIdArr)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.QueryDBErr)
	}

	//已加入的社区并且status=1的社区id 数组
	var sc service.Community
	communityIdArr := sc.GetCommunityIdArr(communities)

	//获取各社区的加入人数
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
	resp := vc.HandlerRespListAll(communities, joinNumMap, topicNumMap)
	return view.SetRespData(resp)

}

/*加入社区
参数：
	userId: 用户id
	cId: 社区号
*/
func (jc *JoinedCommunity) Join(userId int64) interface{} {
	_, TableID, _, err := util.SplitFlowNumber(jc.CId)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.IncorrectFlowNumber)
	}

	var mc model.Community
	err = mc.QueryOne("*", "id = ?", TableID)
	if err != nil {
		util.LoggerError(err)
		return view.CheckMysqlErr(err)
	}

	// 判断社区状态
	//-1 封禁下架, 0 申请中, 1 正常, 2 解散删除
	switch mc.Status {
	case -1:
		return view.SetErr(constant.CommunityProhibition)
	case 0:
		return view.SetErr(constant.CommunityExamining)
	case 2:
		return view.SetErr(constant.CommunityDissolution)
	default:
	}

	var mjc model.JoinedCommunity
	mjc.CommunityId = mc.ID
	mjc.UserId = userId
	err = mjc.Add()
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.CommunityJoinErr)
	}

	return view.SetErr(constant.Success)
}

//退出社区
func (jc *JoinedCommunity) Exit(userId int64) interface{} {
	var mc model.Community
	err := mc.QueryOne("*", "c_id = ?", jc.CId)
	if err != nil {
		util.LoggerError(err)
		return view.CheckMysqlErr(err)
	}

	var mjc model.JoinedCommunity
	err = mjc.Delete("community_id = ? AND user_id = ?", mc.ID, userId)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.CommunityExitErr)
	}
	return view.SetErr(constant.Success)
}
