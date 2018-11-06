package model

import (
	"time"
	"wyatt/api/constant"
	"wyatt/db"
	"wyatt/util"
)

type TopicCollect struct {
	TableModel

	TopicId int64 `json:"tid"`    //话题id, 主键id
	UserId  int64 `json:"userId"` //用户id
}

func (bc *TopicCollect) BeforeCreate() (err error) {
	bc.FlowId = time.Now().Unix()
	return
}

func (tc *TopicCollect) Add() error {
	mdb := db.GetMysqlDB()
	tx := mdb.Begin()
	defer tx.Commit()
	err := tx.Create(tc).Error
	if err != nil {
		tx.Rollback()
		util.LoggerError(err)
		return err
	}
	return nil
}

func (tc *TopicCollect) Delete(where interface{}, args ...interface{}) error {
	mdb := db.GetMysqlDB()
	tx := mdb.Begin()
	defer tx.Commit()
	err := tx.Model(tc).Where(where, args...).Delete(tc).Error
	if err != nil {
		tx.Rollback()
		util.LoggerError(err)
		return err
	}
	return nil
}
func (tc *TopicCollect) QueryOne(field string, where interface{}, args ...interface{}) error {
	mdb := db.GetMysqlDB()
	err := mdb.Model(tc).Select(field).Where(where, args...).Last(tc).Error
	if err != nil {
		util.LoggerError(err)
		return err
	}
	return nil
}
func (tc *TopicCollect) QueryList(field string, page int, limit int, where interface{}, args ...interface{}) ([]TopicCollect, error) {
	if 0 > page {
		page = 0
	}
	if 0 > limit || constant.MAX_QUERY_COUNT < limit {
		limit = constant.MAX_QUERY_COUNT
	}
	mdb := db.GetMysqlDB()
	var list []TopicCollect
	err := mdb.Model(tc).Select(field).Where(where, args...).Offset(page * limit).Offset(limit).Find(&list).Error
	if err != nil {
		util.LoggerError(err)
		return make([]TopicCollect, 0), err
	}
	return list, nil
}
