package model

import (
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
	err := tx.Create(jc).Error
	if err != nil {
		util.LoggerError(err)
		tx.Rollback()
	}
	tx.Commit()
	return err
}
