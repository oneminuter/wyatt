package model

import (
	"wyatt/api/constant"
	"wyatt/db"
	"wyatt/util"
)

//话题
type Topic struct {
	TableModel

	TId         int64  `json:"tId"`                   //话题id，创建的时间戳
	Title       string `json:"title"`                 //标题
	Desc        string `json:"desc" gorm:"type:text"` //简介，详情，或者内容
	CommunityId int64  `json:"communityId"`           //所属社区id
	CreatorId   int64  `json:"creator"`               //创建者id
	ViewedNum   int64  `json:"viewedNum"`             //浏览量
	ZanNum      int64  `json:"zanNum"`                //点赞量
	CommentNum  int64  `json:"commentNum"`            //评论数量
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

func (t *Topic) Add() error {
	mdb := db.GetMysqlDB()
	tx := mdb.Begin()
	defer tx.Commit()
	err := tx.Create(t).Error
	if err != nil {
		tx.Rollback()
		util.LoggerError(err)
		return err
	}
	return nil
}

func (t *Topic) QueryOne(field string, where interface{}, args ...interface{}) error {
	mdb := db.GetMysqlDB()

	err := mdb.Model(t).Select(field).Where(where, args...).Last(t).Error
	if err != nil {
		util.LoggerError(err)
		return err
	}
	return nil
}

func (t *Topic) QueryList(field string, page int, limit int, where interface{}, args ...interface{}) ([]Topic, error) {
	if 0 > page {
		page = 0
	}
	if 0 > limit || constant.MAX_QUERY_COUNT < limit {
		limit = constant.MAX_QUERY_COUNT
	}

	mdb := db.GetMysqlDB()
	var list []Topic
	err := mdb.Model(t).Select(field).Where(where, args...).Offset(page * limit).Limit(limit).Find(&list).Error
	if err != nil {
		util.LoggerError(err)
		return make([]Topic, 0), err
	}
	return list, nil
}

func (t *Topic) Delete(where interface{}, args ...interface{}) error {
	mdb := db.GetMysqlDB()
	tx := mdb.Begin()
	defer tx.Commit()
	err := tx.Where(where, args...).Delete(t).Error
	if err != nil {
		tx.Rollback()
		util.LoggerError(err)
		return err
	}
	return nil
}

func (t *Topic) Update(update interface{}, where interface{}, args ...interface{}) error {
	mdb := db.GetMysqlDB()
	tx := mdb.Begin()
	defer tx.Commit()
	err := tx.Model(t).Where(where, args...).Updates(update).Error
	if err != nil {
		tx.Rollback()
		util.LoggerError(err)
		return err
	}
	return nil
}
