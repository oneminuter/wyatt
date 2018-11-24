package model

import (
	"time"
	"wyatt/api/constant"
	"wyatt/db"
	"wyatt/util"
)

//系列故事集合，整合系列和故事的关系，系列：故事（1：n）
type SeriseStoryGather struct {
	TableModel
	SeriesId int64 `json:"seriesId"` //系列id
	StoryId  int64 `json:"storyId"`  //故事id
}

func (m *SeriseStoryGather) BeforeCreate() (err error) {
	m.FlowId = time.Now().Unix()
	return
}

func (m *SeriseStoryGather) Add() error {
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

func (m *SeriseStoryGather) Delete(where interface{}, args ...interface{}) error {
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

func (m *SeriseStoryGather) Update(update, where interface{}, args ...interface{}) error {
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

func (m *SeriseStoryGather) QueryList(field string, page, limit int, where interface{}, args ...interface{}) ([]SeriseStoryGather, error) {
	if 0 > page {
		page = 0
	}
	if 0 > limit || constant.MAX_QUERY_COUNT < limit {
		limit = constant.MAX_QUERY_COUNT
	}

	mdb := db.GetMysqlDB()
	var list = make([]SeriseStoryGather, 0)
	err := mdb.Model(m).Select(field).Where(where, args...).Offset(page * limit).Limit(limit).Find(&list).Error
	if err != nil {
		util.LoggerError(err)
		return make([]SeriseStoryGather, 0), err
	}
	return list, nil
}

func (m *SeriseStoryGather) QueryOne(field string, where interface{}, args ...interface{}) error {
	mdb := db.GetMysqlDB()
	err := mdb.Model(m).Select(field).Where(where, args...).Last(m).Error
	if err != nil {
		util.LoggerError(err)
		return err
	}

	return nil
}

func (m *SeriseStoryGather) QueryCount(where interface{}, args ...interface{}) (int, error) {
	mdb := db.GetMysqlDB()
	var count int
	err := mdb.Model(m).Where(where, args...).Count(&count).Error
	if err != nil {
		util.LoggerError(err)
		return 0, err
	}
	return count, nil
}

func (m *SeriseStoryGather) QueryGrounp(field string, group string, where interface{}, args ...interface{}) ([]SeriseStoryGather, error) {
	mdb := db.GetMysqlDB()
	var list = make([]SeriseStoryGather, 0)
	err := mdb.Model(m).Select(field).Where(where, args...).Group(group).Find(&list).Error
	if err != nil {
		util.LoggerError(err)
		return make([]SeriseStoryGather, 0), err
	}
	return list, nil
}
