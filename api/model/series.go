package model

import (
	"time"
	"wyatt/api/constant"
	"wyatt/db"
	"wyatt/util"
)

//系列故事，如小说之类的，有很多篇的，连续的故事
type Series struct {
	TableModel

	Title      string `json:"title"`      //标题，在系列里，相当于书名，在故事章节里相当于每节的标题
	Desc       string `json:"desc"`       //简介
	Classify   string `json:"classify"`   //分类，比如情感类，武侠类，转机类
	AuthorId   int64  `json:"author"`     //作者,对应用户的id
	CoverImg   string `json:"coverImg"`   //本节故事的封面图
	ViewedNum  int64  `json:"viewedNum"`  //浏览量, 此字段在故事章节里是每章节的总浏览量，在系列里是该系列所有故事的总浏览量
	ZanNum     int64  `json:"zanNum"`     //点赞量，此字段在故事章节里是每章节的总点赞量，在系列里是该系列所有故事的总点赞量
	CommentNum int64  `json:"commentNum"` //评论数量，此字段在故事章节里是每章节的总评论量，在系列里是该系列所有故事的总评论量
}

func (m *Series) BeforeCreate() (err error) {
	m.FlowId = time.Now().Unix()
	return
}

func (m *Series) Add() error {
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

func (m *Series) Delete(where interface{}, args ...interface{}) error {
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

func (m *Series) Update(update, where interface{}, args ...interface{}) error {
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

func (m *Series) QueryList(field string, page, limit int, where interface{}, args ...interface{}) ([]Series, error) {
	if 0 > page {
		page = 0
	}
	if 0 > limit || constant.MAX_QUERY_COUNT < limit {
		limit = constant.MAX_QUERY_COUNT
	}

	mdb := db.GetMysqlDB()
	var list = make([]Series, 0)
	err := mdb.Model(m).Select(field).Where(where, args...).Offset(page * limit).Limit(limit).Find(&list).Error
	if err != nil {
		util.LoggerError(err)
		return make([]Series, 0), err
	}
	return list, nil
}

func (m *Series) QueryOne(field string, where interface{}, args ...interface{}) error {
	mdb := db.GetMysqlDB()
	err := mdb.Model(m).Select(field).Where(where, args...).Last(m).Error
	if err != nil {
		util.LoggerError(err)
		return err
	}

	return nil
}

func (m *Series) QueryCount(where interface{}, args ...interface{}) (int, error) {
	mdb := db.GetMysqlDB()
	var count int
	err := mdb.Model(m).Where(where, args...).Count(&count).Error
	if err != nil {
		util.LoggerError(err)
		return 0, err
	}
	return count, nil
}

func (m *Series) QueryGrounp(field string, group string, where interface{}, args ...interface{}) ([]Series, error) {
	mdb := db.GetMysqlDB()
	var list = make([]Series, 0)
	err := mdb.Model(m).Select(field).Where(where, args...).Group(group).Find(&list).Error
	if err != nil {
		util.LoggerError(err)
		return make([]Series, 0), err
	}
	return list, nil
}
