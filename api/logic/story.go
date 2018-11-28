package logic

import (
	"strings"
	"wyatt/api/constant"
	"wyatt/api/model"
	"wyatt/api/service"
	"wyatt/api/view"
	"wyatt/util"
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
	content
}
type content struct {
	RoleId  string `json:"roleId" form:"roleId"` //角色流水号id, 旁白为空
	Type    int    `json:"type" form:"type"`     //类型 1 角色对白，2 旁白
	Context string `json:"context"`              //内容
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

	//返回
	var vs view.Story
	retData := vs.List(stories, mu)
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
	return nil
}
