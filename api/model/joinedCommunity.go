package model

import (
	"time"
	"wyatt/db"
	"wyatt/util"
)

//加入的社区
type JoinedCommunity struct {
	TableModel

	UserId      int64 `json:"userId"`
	CommunityId int64 `json:"communityId"` //社区id
	Count       int   `json:"-" gorm:"-"`
}

func (bc *JoinedCommunity) BeforeCreate() (err error) {
	bc.FlowId = time.Now().Unix()
	return
}

//按组查询
func (jc *JoinedCommunity) QueryGrounp(field string, group string, where interface{}, args ...interface{}) ([]JoinedCommunity, error) {
	mdb := db.GetMysqlDB()
	var list []JoinedCommunity
	err := mdb.Model(jc).Select(field).Where(where, args...).Group(group).Find(&list).Error
	if err != nil {
		util.LoggerError(err)
		return make([]JoinedCommunity, 0), err
	}
	return list, nil
}

//查询列表
func (jc *JoinedCommunity) QueryList(field string, where interface{}, args ...interface{}) ([]JoinedCommunity, error) {
	mdb := db.GetMysqlDB()
	var list []JoinedCommunity
	err := mdb.Model(jc).Select(field).Where(where, args...).Find(&list).Error
	if err != nil {
		util.LoggerError(err)
		return make([]JoinedCommunity, 0), err
	}
	return list, nil
}

func (jc *JoinedCommunity) Add() error {
	mdb := db.GetMysqlDB()
	tx := mdb.Begin()
	defer tx.Commit()

	err := tx.Create(jc).Error
	if err != nil {
		util.LoggerError(err)
		tx.Rollback()
		return err
	}
	return nil
}

func (jc *JoinedCommunity) Delete(where interface{}, args ...interface{}) error {
	mdb := db.GetMysqlDB()
	tx := mdb.Begin()
	defer tx.Commit()

	err := tx.Where(where, args...).Delete(jc).Error
	if err != nil {
		tx.Rollback()
		util.LoggerError(err)
		return err
	}
	return nil
}
