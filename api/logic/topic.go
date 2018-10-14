package logic

import (
	"time"
	"wyatt/api/constant"
	"wyatt/api/model"
	"wyatt/api/service"
	"wyatt/api/view"
	"wyatt/util"
)

type Topic struct {
	CId int64 `json:"cId"` //10位数字的社区id
}
type TopicAdd struct {
	CId   int64  `json:"cId" form:"cId" binding:"required"`
	Title string `json:"title" form:"title" binding:"required"` //话题标题
	Desc  string `json:"desc" form:"desc" binding:"required"`   //话题内容
}

//获取社区下的话题列表
func (t *Topic) List() interface{} {
	var (
		mt model.Topic
		vt view.Topic
		mc model.Community
		st service.Topic
		mu model.User
		su service.User
	)
	//查询社区信息
	err := mc.QueryOne("*", "c_id = ?", t.CId)
	if err != nil {
		util.LoggerError(err)
		return view.CheckMysqlErr(err)
	}
	//查询社区下的话题列表
	topics, err := mt.QueryList("*", "community_id = ?", mc.ID)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.QueryDBErr)
	}

	//提取创建者id
	creatorIds := st.GetCreatorIdList(topics)

	//获取用户信息
	ulist, err := mu.QueryList("*", "id IN (?)", creatorIds)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.QueryDBErr)
	}

	//将用户信息转换为map
	uMap := su.TransformToMap(ulist)

	//返回
	list := vt.HandlerRespList(topics, t.CId, uMap)
	return view.SetRespData(list)
}

//增加话题
func (ta *TopicAdd) Add(creatorId int64) interface{} {
	//查询社区
	var mc model.Community
	err := mc.QueryOne("*", "c_id = ?", ta.CId)
	if err != nil {
		util.LoggerError(err)
		return view.CheckMysqlErr(err)
	}

	//够造数据
	mt := model.Topic{
		TId:         time.Now().Unix(),
		Title:       ta.Title,
		Desc:        ta.Desc,
		CommunityId: mc.ID,
		CreatorId:   creatorId,
		ViewedNum:   0,
		ZanNum:      0,
		CommentNum:  0,
	}

	//入库
	err = mt.Add()
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.AddErr)
	}
	return view.SetErr(constant.Success)
}
