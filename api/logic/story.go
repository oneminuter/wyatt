package logic

import (
	"strings"
	"wyatt/api/constant"
	"wyatt/api/model"
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

type StoryModify struct {
	StoryId  string `json:"storyId" form:"storyId" binding:"required"` //故事流水号id
	Title    string `json:"title" form:"title"`
	Desc     string `json:"desc" form:"desc"`
	CoverImg string `json:"coverImg" form:"coverImg"`
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

	var modify = make(map[string]string)
	//标题
	if "" != strings.TrimSpace(sm.Title) {
		modify["title"] = sm.Title
	}
	//描述
	if "" != strings.TrimSpace(sm.Desc) {
		modify["desc"] = sm.Desc
	}
	//封面图
	if "" != strings.TrimSpace(sm.CoverImg) {
		modify["cover_img"] = sm.CoverImg
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
