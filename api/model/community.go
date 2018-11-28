package model

import (
	"time"
	"wyatt/api/constant"
	"wyatt/db"
	"wyatt/util"
)

//社区
type Community struct {
	TableModel

	Logo      string `json:"logo"`                        //社区logo
	Name      string `json:"name" gorm:"unique, size:30"` //社区名
	Desc      string `json:"desc"`                        //社区描述
	CreatorId int64  `json:"creatorId" sql:"index"`       //创建者id
	Status    int    `json:"status" gorm:"size:4"`        //社区状态: -1 封禁下架, 0 申请中, 1 正常, 2 解散删除
}

func (m *Community) BeforeCreate() (err error) {
	m.FlowId = time.Now().Unix()
	return
}

func (m *Community) Add() error {
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

func (m *Community) Delete(where interface{}, args ...interface{}) error {
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

func (m *Community) Update(update, where interface{}, args ...interface{}) error {
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

func (m *Community) QueryList(field string, page, limit int, where interface{}, args ...interface{}) ([]Community, error) {
	if 0 > page {
		page = 0
	}
	if 0 > limit || constant.MAX_QUERY_COUNT < limit {
		limit = constant.MAX_QUERY_COUNT
	}

	mdb := db.GetMysqlDB()
	var list = make([]Community, 0)
	err := mdb.Model(m).Select(field).Where(where, args...).Offset(page * limit).Limit(limit).Find(&list).Error
	if err != nil {
		util.LoggerError(err)
		return make([]Community, 0), err
	}
	return list, nil
}

func (m *Community) QueryOne(field string, where interface{}, args ...interface{}) error {
	mdb := db.GetMysqlDB()
	err := mdb.Model(m).Select(field).Where(where, args...).Last(m).Error
	if err != nil {
		util.LoggerError(err)
		return err
	}

	return nil
}

func (m *Community) QueryCount(where interface{}, args ...interface{}) (int, error) {
	mdb := db.GetMysqlDB()
	var count int
	err := mdb.Model(m).Where(where, args...).Count(&count).Error
	if err != nil {
		util.LoggerError(err)
		return 0, err
	}
	return count, nil
}

func (m *Community) QueryGrounp(field string, group string, where interface{}, args ...interface{}) ([]Community, error) {
	mdb := db.GetMysqlDB()
	var list = make([]Community, 0)
	err := mdb.Model(m).Select(field).Where(where, args...).Group(group).Find(&list).Error
	if err != nil {
		util.LoggerError(err)
		return make([]Community, 0), err
	}
	return list, nil
}
