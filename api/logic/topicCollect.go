package logic

import (
	"log"
	"wyatt/api/constant"
	"wyatt/api/model"
	"wyatt/api/service"
	"wyatt/api/view"
	"wyatt/util"
)

type TopicCollect struct {
	TId   string `json:"tId" form:"tId" binding:"required"` //话题id, 10位数字-为创建时的时间戳
	Page  int    `json:"page" form:"page"`                  //页码，默认从0开始
	Limit int    `json:"limit" form:"limit"`                //查询条数, 最大查询 constant.MAX_QUERY_COUNT
}

//收藏话题
func (tc *TopicCollect) Add(userId int64) interface{} {
	//查询用户
	var mu model.User
	err := mu.QueryOne("*", "id = ?", userId)
	if err != nil {
		util.LoggerError(err)
		return view.CheckMysqlErr(err)
	}

	_, TableID, _, err := util.SplitFlowNumber(tc.TId)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.IncorrectFlowNumber)
	}

	//查询话题信息
	var mt model.Topic
	err = mt.QueryOne("*", "id = ?", TableID)
	if err != nil {
		util.LoggerError(err)
		return view.CheckMysqlErr(err)
	}

	//添加收藏信息
	var mtc = model.TopicCollect{
		TopicId: mt.ID,
		UserId:  userId,
	}
	err = mtc.Add()
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.AddErr)
	}

	return view.SetErr(constant.Success)
}

//取消收藏话题
func (tc *TopicCollect) Cancel(userId int64) interface{} {
	//查询用户
	var mu model.User
	err := mu.QueryOne("*", "id = ?", userId)
	if err != nil {
		util.LoggerError(err)
		return view.CheckMysqlErr(err)
	}

	_, TableID, _, err := util.SplitFlowNumber(tc.TId)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.IncorrectFlowNumber)
	}
	//查询话题信息
	var mt model.Topic
	err = mt.QueryOne("*", "id = ?", TableID)
	if err != nil {
		util.LoggerError(err)
		return view.CheckMysqlErr(err)
	}

	//取消收藏
	var mtc model.TopicCollect
	err = mtc.Delete("topic_id = ? AND user_id = ?", mt.ID, userId)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.DeleteErr)
	}
	return view.SetErr(constant.Success)
}

//收藏列表
func (tc *TopicCollect) List(userId int64) interface{} {
	//判断参数是否合法
	if 0 > tc.Page || 0 > tc.Limit {
		return view.SetErr(constant.QueryPageOrLimit)
	}
	if constant.MAX_QUERY_COUNT < tc.Limit {
		tc.Limit = constant.MAX_QUERY_COUNT
	}

	//查询用户
	var mu model.User
	err := mu.QueryOne("*", "id = ?", userId)
	if err != nil {
		util.LoggerError(err)
		return view.CheckMysqlErr(err)
	}

	//查询收藏的话题
	var mtc model.TopicCollect
	collects, err := mtc.QueryList("*", "user_id = ?", userId)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.QueryDBErr)
	}

	//提取话题id数组
	var stc service.TopicCollect
	topicIdArr := stc.GetTopicIdArr(collects)
	if 0 == len(topicIdArr) {
		util.Logger("话题id为空")
		return view.SetErr(constant.NoCollect)
	}
	log.Println(topicIdArr)
	//查询话题信息
	var mt model.Topic
	topics, err := mt.QueryList("*", tc.Page, tc.Limit, "id IN (?)", topicIdArr)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.QueryDBErr)
	}
	log.Println(topics)
	//提取社区id数组
	var st service.Topic
	commIdArr := st.GetCommunityIdList(topics)
	if 0 == len(commIdArr) {
		util.Logger("社区id为空")
		return view.SetErr(constant.NoCollect)
	}

	//查询社区信息
	var mc model.Community
	communities, err := mc.QueryList("*", 0, tc.Limit, "id IN (?)", commIdArr)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.QueryDBErr)
	}

	//社区主键id:社区 map
	var sc service.Community
	communityMap := sc.GetCommunityMap(communities)

	//提取话题创建者信息
	creatorIdArr := st.GetCreatorIdList(topics)

	//查询创建者信息
	users, err := mu.QueryList("*", "id IN (?)", creatorIdArr)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.QueryDBErr)
	}

	//用户信息map对应，用户id:用户信息
	var su service.User
	uMap := su.TransformToMap(users)

	//处理返回
	var vt view.Topic
	list := vt.HandlerRespCollectList(topics, communityMap, uMap)

	return view.SetRespData(list)
}
