package logic

import (
	"strings"
	"wyatt/api/constant"
	"wyatt/api/model"
	"wyatt/api/service"
	"wyatt/api/view"
	"wyatt/util"

	"strconv"

	"github.com/json-iterator/go"
)

type Story struct {
	Title    string `json:"title" form:"title" binding:"required"` //故事标题
	Desc     string `json:"desc"`                                  //简介
	Classify string `json:"classify"`                              //分类，比如情感类，武侠类，转机类
	CoverImg string `json:"coverImg"`                              //本节故事的封面图
}

type StoryList struct {
	UserAccount string `json:"userAccount" form:"userAccount" binding:"required"`
	Page        int    `json:"page" form:"page"` //页码，从0开始，默认为0
	Limit       int    `json:"limit" form:"limit"`
}

//修改故事
type StoryModify struct {
	StoryId  string `json:"storyId" form:"storyId" binding:"required"` //故事流水号id
	Title    string `json:"title" form:"title"`
	Desc     string `json:"desc" form:"desc"`
	CoverImg string `json:"coverImg" form:"coverImg"`
}

//添加故事详情
type StoryContentAdd struct {
	StoryId string `json:"storyId" form:"storyId" binding:"required"` //故事流水号id
	Content string `json:"content" form:"content" binding:"required"` //故事详情内容, json数组字符串，json的内部字段为 contentAdd
}

//故事内容详情字段
type contentAdd struct {
	RoleId  string `json:"roleId"`  //角色流水号id, 旁白为空
	Type    string `json:"type"`    //类型 1 角色对白，2 旁白， 这里接收是string类型，因为前端传过来的数据经过 JSON.stringify() 之后，数字会变成字符串
	Context string `json:"context"` //内容
}

//修改故事
type StoryContentModify struct {
	StoryId string `json:"storyId" form:"storyId" binding:"required"` //故事流水号id
	Content string `json:"content" form:"content" binding:"required"` //故事详情内容, json数组字符串，json的内部字段为 contentModify
}
type contentModify struct {
	SCId  string `json:"scId"`  //内容流水号id, 根据该字段的有无判断是新增还是修改
	Order string `json:"order"` //排序
	contentAdd
}

type StoryContentList struct {
	StoryId string `json:"storyId" form:"storyId" binding:"required"` //故事流水号id
	Page    int    `json:"page" form:"page"`                          //页码，从0开始，默认为0
	Limit   int    `json:"limit" form:"limit"`
}

func (s *Story) Add(userId int64) interface{} {
	//添加
	ms := model.Story{
		Title:    s.Title,
		Desc:     s.Desc,
		Classify: s.Classify,
		AuthorId: userId,
		CoverImg: s.CoverImg,
	}
	err := ms.Add()
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.AddErr)
	}
	return view.SetErr(constant.Success)
}

func (sl *StoryList) List() interface{} {
	//查询作者信息
	var mu model.User
	err := mu.QueryOne("*", "account = ?", sl.UserAccount)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.QueryDBErr)
	}

	//查询该作者的故事列表
	var ms model.Story
	stories, err := ms.QueryList("*", sl.Page, sl.Limit, "author_id = ?", mu.ID)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.QueryDBErr)
	}

	//查询主角信息
	var msr model.StoryRole
	if 0 != ms.MajorId {
		err = msr.QueryOne("*", "id = ?", ms.MajorId)
		if err != nil {
			util.LoggerError(err)
			return view.CheckMysqlErr(err)
		}
	}

	//返回
	var vs view.Story
	retData := vs.List(stories, mu, msr)
	return view.SetRespData(retData)
}

func (sm *StoryModify) Modify(userId int64) interface{} {
	_, TableID, _, err := util.SplitFlowNumber(sm.StoryId)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.IncorrectFlowNumber)
	}

	var (
		modify       = make(map[string]string)
		isHaveModify bool
	)
	//标题
	if "" != strings.TrimSpace(sm.Title) {
		modify["title"] = sm.Title
		isHaveModify = true
	}
	//描述
	if "" != strings.TrimSpace(sm.Desc) {
		modify["desc"] = sm.Desc
		isHaveModify = true
	}
	//封面图
	if "" != strings.TrimSpace(sm.CoverImg) {
		modify["cover_img"] = sm.CoverImg
		isHaveModify = true
	}

	if !isHaveModify {
		return view.SetErr(constant.NoModify)
	}

	//更新
	var ms model.Story
	err = ms.Update(modify, "id = ?", TableID)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.ModifyErr)
	}
	return view.SetErr(constant.Success)
}

//系列列表
func (sl *StoryList) SeriesList() interface{} {
	//查询用户信息
	var mu model.User
	err := mu.QueryOne("*", "account = ?", sl.UserAccount)
	if err != nil {
		util.LoggerError(err)
		return view.CheckMysqlErr(err)
	}

	//查询系列列表
	var ms model.Series
	series, err := ms.QueryList("*", sl.Page, sl.Limit, "author_id = ?", mu.ID)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.QueryDBErr)
	}

	var vs view.Series
	retData := vs.List(series, mu)
	return view.SetRespData(retData)
}

//故事细节列表
func (scl *StoryContentList) List() interface{} {
	_, TableID, _, err := util.SplitFlowNumber(scl.StoryId)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.IncorrectFlowNumber)
	}

	//查询细节列表
	var msc model.StoryContent
	contents, err := msc.QueryList("*", scl.Page, scl.Limit, "story_id = ?", TableID)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.QueryDBErr)
	}

	//提取角色id
	var ssc service.StoryContent
	roleIdList := ssc.GetRoleIdList(contents)

	//查询角色信息列表
	var msr model.StoryRole
	roles, err := msr.QueryList("*", 0, scl.Limit, "id IN (?)", roleIdList)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.QueryDBErr)
	}
	//角色信息转为map id:role
	var ssr service.StoryRole
	roleMap := ssr.GetRoleMap(roles)

	//返回
	var vsc view.StoryContent
	retData := vsc.HandlerRespList(contents, roleMap)
	return view.SetRespData(retData)
}

//添加故事细节
func (sca *StoryContentAdd) Add(userId int64) interface{} {
	_, storyID, _, err := util.SplitFlowNumber(sca.StoryId)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.IncorrectFlowNumber)
	}

	var (
		clist   []contentAdd
		msc     model.StoryContent
		t       int
		rolerId int64
	)

	//解析内容
	err = jsoniter.Unmarshal([]byte(sca.Content), &clist)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.UnmarshalContentErr)
	}

	for _, v := range clist {
		//故事类型转换数据类型
		t, err = strconv.Atoi(v.Type)
		if err != nil {
			util.LoggerError(err)
			return view.SetErr(constant.UnmarshalContentErr)
		}
		//分割角色id
		_, rolerId, _, err = util.SplitFlowNumber(v.RoleId)
		if err != nil {
			util.LoggerError(err)
			return view.SetErr(constant.IncorrectFlowNumber)
		}

		//构造数据
		msc = model.StoryContent{
			StoryId: storyID,
			Type:    t,
			RoleId:  rolerId,
			Context: v.Context,
		}

		//入库 - 采用一条条插入是因为使用构造sql语句批量插入是，一些自动填充字段为null
		err = msc.Add()
		if err != nil {
			util.LoggerError(err)
			return view.SetErr(constant.AddErr)
		}
	}

	return view.SetErr(constant.Success)
}

//修改故事内容
func (scm *StoryContentModify) Modify(userId int64) interface{} {
	_, storyID, _, err := util.SplitFlowNumber(scm.StoryId)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.IncorrectFlowNumber)
	}

	//判断用户是否是创建者
	var ss service.Story
	if !ss.IsCreator(storyID, userId) {
		return view.SetErr(constant.NoAuth)
	}

	var (
		cList     []contentModify
		msc       model.StoryContent
		t         int
		rolerID   int64
		order     float64
		contentID int64
	)

	//解析内容
	err = jsoniter.Unmarshal([]byte(scm.Content), &cList)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.UnmarshalContentErr)
	}

	for _, v := range cList {
		//故事类型转换数据类型
		t, err = strconv.Atoi(v.Type)
		if err != nil {
			util.LoggerError(err)
			return view.SetErr(constant.UnmarshalContentErr)
		}

		//分割角色id
		_, rolerID, _, err = util.SplitFlowNumber(v.RoleId)
		if err != nil {
			util.LoggerError(err)
			return view.SetErr(constant.IncorrectFlowNumber)
		}

		//排序类型转换
		order, err = strconv.ParseFloat(v.Order, 64)
		if err != nil {
			util.LoggerError(err)
			return view.SetErr(constant.UnmarshalContentErr)
		}

		//构造数据
		msc = model.StoryContent{
			StoryId: storyID,
			Type:    t,
			RoleId:  rolerID,
			Context: v.Context,
			Order:   order,
		}

		if "" != strings.TrimSpace(v.SCId) {
			_, contentID, _, err = util.SplitFlowNumber(v.SCId)
			//修改
			if err != nil {
				util.LoggerError(err)
				return view.SetErr(constant.IncorrectFlowNumber)
			}
			err = msc.Update(msc, "id = ?", contentID)
		} else {
			//添加
			err = msc.Add()
		}

		if err != nil {
			util.LoggerError(err)
			return view.SetErr(constant.ModifyErr)
		}
	}
	return view.SetErr(constant.Success)
}
