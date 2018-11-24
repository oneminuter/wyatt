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
	ViewedNum  int64  `json:"viewedNum"`  //浏览量
	ZanNum     int64  `json:"zanNum"`     //点赞量
	CommentNum int64  `json:"commentNum"` //评论数量
}

func (s *Story) List(msList []model.Story, u model.User) []Story {
	var list = make([]Story, 0, len(msList))
	for _, v := range msList {
		list = append(list, Story{
			StoryId:   fmt.Sprintf("%s.%d.%d", constant.S, v.ID, v.FlowId),
			Title:     v.Title,
			Desc:      v.Desc,
			Classify:  v.Classify,
			CoverImg:  v.CoverImg,
			ViewedNum: v.ViewedNum,
			ZanNum:    v.ZanNum,
		})
	}
	return list
}
