package model

import (
	"wyatt/db"
	"wyatt/util"
)

//话题
type Topic struct {
	TableModel

	Title       string `json:"title"`       //标题
	Desc        string `json:"desc"`        //简介
	CommunityId int64  `json:"communityId"` //所属社区id
	CreatorId   int64  `json:"creator"`     //创建者id
	ViewedNum   int64  `json:"viewedNum"`   //浏览量
	ZanNum      int64  `json:"zanNum"`      //点赞量
	CommentNum  int64  `json:"commentNum"`  //评论数量
	Count       int    `json:"-" gorm:"-"`
}

//按组查询
func (t *Topic) QueryGrounp(field string, group string, where interface{}, args ...interface{}) ([]Topic, error) {
	mdb := db.GetMysqlDB()
	var list []Topic
	err := mdb.Model(t).Select(field).Where(where, args...).Group(group).Find(&list).Error
	if err != nil {
		util.LoggerError(err)
		return make([]Topic, 0), err
	}
	return list, nil
}
