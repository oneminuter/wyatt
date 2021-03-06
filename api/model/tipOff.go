package model

import (
	"time"
	"wyatt/api/constant"
	"wyatt/db"
	"wyatt/util"
)

//内容举报
type TipOff struct {
	TableModel
	UserId       int64  `json:"userId"`                                  //用户id
	SourceFlowId string `json:"sourceFlowId" gorm:"size:30" sql:"index"` //举报内容完整流水号
	Reason       string `json:"reason"`                                  //举报原因
	Status       int    `json:"status" gorm:"size:4"`                    //处理状态, 0 未处理，1 已处理
	Remark       string `json:"remark"`                                  //处理备注
}

func (m *TipOff) BeforeCreate() (err error) {
	m.FlowId = time.Now().Unix()
	return
}

func (m *TipOff) Add() error {
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

func (m *TipOff) Delete(where interface{}, args ...interface{}) error {
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

func (m *TipOff) Update(update, where interface{}, args ...interface{}) error {
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

func (m *TipOff) QueryList(field string, page, limit int, where interface{}, args ...interface{}) ([]TipOff, error) {
	if 0 > page {
		page = 0
	}
	if 0 > limit || constant.MAX_QUERY_COUNT < limit {
		limit = constant.MAX_QUERY_COUNT
	}

	mdb := db.GetMysqlDB()
	var list = make([]TipOff, 0)
	err := mdb.Model(m).Select(field).Where(where, args...).Offset(page * limit).Limit(limit).Find(&list).Error
	if err != nil {
		util.LoggerError(err)
		return make([]TipOff, 0), err
	}
	return list, nil
}

func (m *TipOff) QueryOne(field string, where interface{}, args ...interface{}) error {
	mdb := db.GetMysqlDB()
	err := mdb.Model(m).Select(field).Where(where, args...).Last(m).Error
	if err != nil {
		util.LoggerError(err)
		return err
	}

	return nil
}

func (m *TipOff) QueryCount(where interface{}, args ...interface{}) (int, error) {
	mdb := db.GetMysqlDB()
	var count int
	err := mdb.Model(m).Where(where, args...).Count(&count).Error
	if err != nil {
		util.LoggerError(err)
		return 0, err
	}
	return count, nil
}

func (m *TipOff) QueryGrounp(field string, group string, where interface{}, args ...interface{}) ([]TipOff, error) {
	mdb := db.GetMysqlDB()
	var list = make([]TipOff, 0)
	err := mdb.Model(m).Select(field).Where(where, args...).Group(group).Find(&list).Error
	if err != nil {
		util.LoggerError(err)
		return make([]TipOff, 0), err
	}
	return list, nil
}
