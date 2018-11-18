package logic

import (
	"strings"
	"wyatt/api/constant"
	"wyatt/api/model"
	"wyatt/api/service"
	"wyatt/api/view"
	"wyatt/util"
)

type Topic struct {
	CId   string `json:"cId" form:"cId"`     //表别名+10位数字的社区id
	Page  int    `json:"page" form:"page"`   //页码，默认从0开始
	Limit int    `json:"limit" form:"limit"` //查询条数, 最大查询 constant.MAX_QUERY_COUNT
}
type TopicAdd struct {
	CId   string `json:"cId" form:"cId" binding:"required"`
	Title string `json:"title" form:"title" binding:"required"` //话题标题
	Desc  string `json:"desc" form:"desc" binding:"required"`   //话题内容
}
type TopicDelete struct {
	Tid string `json:"tid" form:"tId" binding:"required"` //话题id
}
type TopicModify struct {
	Tid   string `json:"tid" form:"tId" binding:"required"` //话题id
	Title string `json:"title" form:"title"`                //话题标题
	Desc  string `json:"desc" form:"desc"`                  //话题内容
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
	_, TableID, timesteamp, err := util.SplitFlowNumber(t.CId)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.IncorrectFlowNumber)
	}
	//查询社区信息
	err = mc.QueryOne("*", "id = ?", TableID)
	if err != nil {
		util.LoggerError(err)
		return view.CheckMysqlErr(err)
	}
	//查询社区下的话题列表
	topics, err := mt.QueryList("*", t.Page, t.Limit, "community_id = ?", mc.ID)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.QueryDBErr)
	}

	//提取创建者id
	creatorIds := st.GetCreatorIdList(topics)

	//获取用户信息
	ulist, err := mu.QueryList("*", 0, t.Limit, "id IN (?)", creatorIds)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.QueryDBErr)
	}

	//将用户信息转换为map
	uMap := su.TransformToMap(ulist)

	//返回
	list := vt.HandlerRespList(topics, TableID, timesteamp, uMap)
	return view.SetRespData(list)
}

//增加话题
func (ta *TopicAdd) Add(creatorId int64) interface{} {
	_, TableID, _, err := util.SplitFlowNumber(ta.CId)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.IncorrectFlowNumber)
	}
	//查询社区
	var mc model.Community
	err = mc.QueryOne("*", "id = ?", TableID)
	if err != nil {
		util.LoggerError(err)
		return view.CheckMysqlErr(err)
	}

	//够造数据
	mt := model.Topic{
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

	//送积分
	var sir service.IntegralRecord
	go sir.AddIntegral(creatorId, model.OPT_ADD_TOPIC)

	return view.SetErr(constant.Success)
}

//删除话题
func (td *TopicDelete) Delete(userId int64) interface{} {
	_, TableID, _, err := util.SplitFlowNumber(td.Tid)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.IncorrectFlowNumber)
	}
	var mt model.Topic
	err = mt.QueryOne("*", "id = ?", TableID)
	if err != nil {
		util.LoggerError(err)
		view.CheckMysqlErr(err)
	}

	//判断是否是话题的创建者
	if userId != mt.CreatorId {
		return view.SetErr(constant.NoAuth)
	}
	err = mt.Delete("id = ?", TableID)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.DeleteErr)
	}
	return view.SetErr(constant.Success)
}

//修改话题
func (tm *TopicModify) Modify(userId int64) interface{} {
	_, TableID, _, err := util.SplitFlowNumber(tm.Tid)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.IncorrectFlowNumber)
	}
	var mt model.Topic
	err = mt.QueryOne("*", "id = ?", TableID)
	if err != nil {
		util.LoggerError(err)
		view.CheckMysqlErr(err)
	}
	//判断是否是话题的创建者
	if userId != mt.CreatorId {
		return view.SetErr(constant.NoAuth)
	}

	var (
		m        = make(map[string]string)
		isModify = false
	)

	//判断title是否为空
	if "" != strings.TrimSpace(tm.Title) {
		m["title"] = tm.Title
		isModify = true
	}
	//判断内容是否为空
	if "" != strings.TrimSpace(tm.Desc) {
		m["desc"] = tm.Desc
		isModify = true
	}

	if !isModify {
		return view.SetErr(constant.NoModify)
	}

	err = mt.Update(m, "id = ?", TableID)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.ModifyErr)
	}

	return view.SetErr(constant.Success)
}

//查看话题详情
func (Topic) Detail(tId string) interface{} {
	_, TableID, _, err := util.SplitFlowNumber(tId)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.IncorrectFlowNumber)
	}

	var mt model.Topic
	err = mt.QueryOne("*", "id = ?", TableID)
	if err != nil {
		util.LoggerError(err)
		return view.CheckMysqlErr(err)
	}

	//查询话题所属社区号
	var mc model.Community
	err = mc.QueryOne("*", "id = ?", mt.CommunityId)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.QueryDBErr)
	}

	//查询创建者信息
	var mu model.User
	err = mu.QueryOne("*", "id = ?", mt.CreatorId)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.QueryDBErr)
	}

	//返回视图
	var vt view.Topic
	vt.HandlerRespDetail(mt, mc.ID, mc.FlowId, mu)

	return view.SetRespData(vt)
}
