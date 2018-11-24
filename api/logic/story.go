package logic

import (
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

func (s *Story) Add(userId int64) interface{} {
	//查询用户信息，看用户状态
	var mu model.User
	err := mu.QueryOne("*", "id = ?", userId)
	if err != nil {
		util.LoggerError(err)
		return view.SetErr(constant.QueryDBErr)
	}

	if mu.Status == -1 {
		return view.SetErr(constant.AccountForbid)
	}

	//添加
	ms := model.Story{
		Title:    s.Title,
		Desc:     s.Desc,
		Classify: s.Classify,
		AuthorId: userId,
		CoverImg: s.CoverImg,
	}
	err = ms.Add()
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
