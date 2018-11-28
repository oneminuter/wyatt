package model

import (
	"time"
	"wyatt/api/constant"
	"wyatt/db"
	"wyatt/util"
)

//话题
type Topic struct {
	TableModel

	Title       string `json:"title" gorm:"size:30"`    //标题
	Desc        string `json:"desc" gorm:"type:text"`   //简介，详情，或者内容
	CommunityId int64  `json:"communityId" sql:"index"` //所属社区id
	CreatorId   int64  `json:"creatorId" sql:"index"`   //创建者id
	ViewedNum   int64  `json:"viewedNum"`               //浏览量
	ZanNum      int64  `json:"zanNum"`                  //点赞量
	CommentNum  int64  `json:"commentNum"`              //评论数量
	Count       int    `json:"-" gorm:"-"`
}

func (m *Topic) BeforeCreate() (err error) {
	m.FlowId = time.Now().Unix()
	return
}

func (m *Topic) Add() error {
	mdb := db.GetMysqlDB()
	tx := mdb.Begin()
	defer tx.Commit()
	err := tx.Create(m).Error
	if err != nil {
		util.LoggerError(err)
		tx.Rollback()
		return err
	}
	return nil
}

func (m *Topic) Delete(where interface{}, args ...interface{}) error {
	mdb := db.GetMysqlDB()
	tx := mdb.Begin()
	defer tx.Commit()
	err := tx.Model(m).Where(where, args...).Delete(m).Error
	if err != nil {
		util.LoggerError(err)
		tx.Rollback()
		return err
	}
	return nil
}

func (m *Topic) Update(update, where interface{}, args ...interface{}) error {
	mdb := db.GetMysqlDB()
	tx := mdb.Begin()
	defer tx.Commit()
	err := tx.Model(m).Where(where, args...).Updates(update).Error
	if err != nil {
		util.LoggerError(err)
		tx.Rollback()
		return err
	}
	return nil
}

func (m *Topic) QueryList(field string, page, limit int, where interface{}, args ...interface{}) ([]Topic, error) {
	if 0 > page {
		page = 0
	}
	if 0 > limit || constant.MAX_QUERY_COUNT < limit {
		limit = constant.MAX_QUERY_COUNT
	}

	mdb := db.GetMysqlDB()
	var list = make([]Topic, 0)
	err := mdb.Model(m).Select(field).Where(where, args...).Offset(page * limit).Limit(limit).Find(&list).Error
	if err != nil {
		util.LoggerError(err)
		return make([]Topic, 0), err
	}
	return list, nil
}

func (m *Topic) QueryOne(field string, where interface{}, args ...interface{}) error {
	mdb := db.GetMysqlDB()
	err := mdb.Model(m).Select(field).Where(where, args...).Last(m).Error
	if err != nil {
		util.LoggerError(err)
		return err
	}

	return nil
}

func (m *Topic) QueryCount(where interface{}, args ...interface{}) (int, error) {
	mdb := db.GetMysqlDB()
	var count int
	err := mdb.Model(m).Where(where, args...).Count(&count).Error
	if err != nil {
		util.LoggerError(err)
		return 0, err
	}
	return count, nil
}

func (m *Topic) QueryGrounp(field string, group string, where interface{}, args ...interface{}) ([]Topic, error) {
	mdb := db.GetMysqlDB()
	var list = make([]Topic, 0)
	err := mdb.Model(m).Select(field).Where(where, args...).Group(group).Find(&list).Error
	if err != nil {
		util.LoggerError(err)
		return make([]Topic, 0), err
	}
	return list, nil
}
