package model

import (
	"time"
	"wyatt/api/constant"
	"wyatt/db"
	"wyatt/util"
)

//社区人员管理
type CommunityManager struct {
	TableModel

	CommunityId int64 `json:"communityId" sql:"index"` // community 的主键
	UserId      int64 `json:"userId"`                  //用户id
	Role        int   `json:"role" gorm:"size:4"`      // -1 封禁， 0 普通成员，1 管理员  其他为游客，封禁状态不能参与该社区的话题和发言
}

func (m *CommunityManager) BeforeCreate() (err error) {
	m.FlowId = time.Now().Unix()
	return
}

func (m *CommunityManager) Add() error {
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

func (m *CommunityManager) Delete(where interface{}, args ...interface{}) error {
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

func (m *CommunityManager) Update(update, where interface{}, args ...interface{}) error {
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

func (m *CommunityManager) QueryList(field string, page, limit int, where interface{}, args ...interface{}) ([]CommunityManager, error) {
	if 0 > page {
		page = 0
	}
	if 0 > limit || constant.MAX_QUERY_COUNT < limit {
		limit = constant.MAX_QUERY_COUNT
	}

	mdb := db.GetMysqlDB()
	var list = make([]CommunityManager, 0)
	err := mdb.Model(m).Select(field).Where(where, args...).Offset(page * limit).Limit(limit).Find(&list).Error
	if err != nil {
		util.LoggerError(err)
		return make([]CommunityManager, 0), err
	}
	return list, nil
}

func (m *CommunityManager) QueryOne(field string, where interface{}, args ...interface{}) error {
	mdb := db.GetMysqlDB()
	err := mdb.Model(m).Select(field).Where(where, args...).Last(m).Error
	if err != nil {
		util.LoggerError(err)
		return err
	}

	return nil
}

func (m *CommunityManager) QueryCount(where interface{}, args ...interface{}) (int, error) {
	mdb := db.GetMysqlDB()
	var count int
	err := mdb.Model(m).Where(where, args...).Count(&count).Error
	if err != nil {
		util.LoggerError(err)
		return 0, err
	}
	return count, nil
}

func (m *CommunityManager) QueryGrounp(field string, group string, where interface{}, args ...interface{}) ([]CommunityManager, error) {
	mdb := db.GetMysqlDB()
	var list = make([]CommunityManager, 0)
	err := mdb.Model(m).Select(field).Where(where, args...).Group(group).Find(&list).Error
	if err != nil {
		util.LoggerError(err)
		return make([]CommunityManager, 0), err
	}
	return list, nil
}
