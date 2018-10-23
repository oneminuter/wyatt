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
	TId int64 `json:"tId" form:"tId" binding:"required"` //话题id, 10位数字-为创建时的时间戳
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

	//查询话题信息
	var mt model.Topic
	err = mt.QueryOne("*", "t_id = ?", tc.TId)
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

	//查询话题信息
	var mt model.Topic
	err = mt.QueryOne("*", "t_id = ?", tc.TId)
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
	topics, err := mt.QueryList("*", "id IN (?)", topicIdArr)
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
	communities, err := mc.QueryList("*", "id IN (?)", commIdArr)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.QueryDBErr)
	}

	//社区主键id:社区id map
	var sc service.Community
	cidMap := sc.GetCIDMap(communities)

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
	list := vt.HandlerRespCollectList(topics, cidMap, uMap)

	return view.SetRespData(list)
}
