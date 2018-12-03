package view

import (
	"fmt"
	"wyatt/api/constant"
	"wyatt/api/model"
)

type Story struct {
	StoryId    string `json:"storyId"`    //故事流水号id
	Title      string `json:"title"`      //标题
	Desc       string `json:"desc"`       //简介
	Classify   string `json:"classify"`   //分类
	CoverImg   string `json:"coverImg"`   //本节故事的封面图
	Author     string `json:"author"`     //作者账号
	AvatarUrl  string `json:"avatarUrl"`  //作者头像
	MajorId    string `json:"majorId"`    //主角的流水号id
	ViewedNum  int64  `json:"viewedNum"`  //浏览量
	ZanNum     int64  `json:"zanNum"`     //点赞量
	CommentNum int64  `json:"commentNum"` //评论数量
}
type Series struct {
	SeriesId   string `json:"seriesId"`   //系列流水号id
	Title      string `json:"title"`      //标题
	Desc       string `json:"desc"`       //简介
	Classify   string `json:"classify"`   //分类
	CoverImg   string `json:"coverImg"`   //本节故事的封面图
	Author     string `json:"author"`     //作者账号
	AvatarUrl  string `json:"avatarUrl"`  //作者头像
	ViewedNum  int64  `json:"viewedNum"`  //浏览量
	ZanNum     int64  `json:"zanNum"`     //点赞量
	CommentNum int64  `json:"commentNum"` //评论数量
}

func (s *Story) List(msList []model.Story, u model.User, major model.StoryRole) []Story {
	var list = make([]Story, 0, len(msList))
	for _, v := range msList {
		list = append(list, Story{
			StoryId:   fmt.Sprintf("%s.%d.%d", constant.S, v.ID, v.FlowId),
			Title:     v.Title,
			Desc:      v.Desc,
			Classify:  v.Classify,
			Author:    u.Account,
			AvatarUrl: u.AvatarUrl,
			MajorId:   fmt.Sprintf("%s.%d.%d", constant.R, major.ID, major.FlowId),
			CoverImg:  v.CoverImg,
			ViewedNum: v.ViewedNum,
			ZanNum:    v.ZanNum,
		})
	}
	return list
}

func (s *Series) List(msList []model.Series, u model.User) []Series {
	var list = make([]Series, 0, len(msList))
	for _, v := range msList {
		list = append(list, Series{
			SeriesId:   fmt.Sprintf("%s.%d.%d", constant.SR, v.ID, v.FlowId),
			Title:      v.Title,
			Desc:       v.Desc,
			Classify:   v.Classify,
			CoverImg:   v.CoverImg,
			Author:     u.Account,
			AvatarUrl:  u.AvatarUrl,
			ViewedNum:  v.ViewedNum,
			ZanNum:     v.ZanNum,
			CommentNum: v.CommentNum,
		})
	}
	return list
}

//故事信息
func (s *Story) Info(ms model.Story, mu model.User, major model.StoryRole) {
	s.StoryId = fmt.Sprintf("%s.%d.%d", constant.S, ms.ID, ms.FlowId)
	s.Title = ms.Title
	s.Desc = ms.Desc
	s.Classify = ms.Classify
	s.CoverImg = ms.CoverImg
	s.Author = mu.Account
	s.AvatarUrl = mu.AvatarUrl
	s.MajorId = fmt.Sprintf("%s.%d.%d", constant.R, major.ID, major.FlowId)
	s.ViewedNum = ms.ViewedNum
	s.ZanNum = ms.ZanNum
	s.CommentNum = ms.CommentNum
}
