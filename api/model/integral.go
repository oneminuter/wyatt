package model

import (
	"time"
	"wyatt/api/constant"
	"wyatt/db"
	"wyatt/util"
)

//积分
type Integral struct {
	TableModel

	Avaliable int   `json:"avaliable"`          //可用积分，可消耗，1积分等于1分钱
	Growth    int   `json:"growth"`             //成长值，判断用户等级
	UserId    int64 `json:"userId" sql:"index"` //用户id
	Level     int   `json:"level"`              //用户等级
}

func (m *Integral) BeforeCreate() (err error) {
	m.FlowId = time.Now().Unix()
	return
}

func (m *Integral) Add() error {
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

func (m *Integral) Delete(where interface{}, args ...interface{}) error {
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

func (m *Integral) Update(update, where interface{}, args ...interface{}) error {
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

func (m *Integral) QueryList(field string, page, limit int, where interface{}, args ...interface{}) ([]Integral, error) {
	if 0 > page {
		page = 0
	}
	if 0 > limit || constant.MAX_QUERY_COUNT < limit {
		limit = constant.MAX_QUERY_COUNT
	}

	mdb := db.GetMysqlDB()
	var list = make([]Integral, 0)
	err := mdb.Model(m).Select(field).Where(where, args...).Offset(page * limit).Limit(limit).Find(&list).Error
	if err != nil {
		util.LoggerError(err)
		return make([]Integral, 0), err
	}
	return list, nil
}

func (m *Integral) QueryOne(field string, where interface{}, args ...interface{}) error {
	mdb := db.GetMysqlDB()
	notFound := mdb.Model(m).Select(field).Where(where, args...).Last(m).RecordNotFound()
	//不存在，则初始化一条记录
	if notFound {
		err := mdb.FirstOrCreate(m).Error
		if err != nil {
			util.LoggerError(err)
			return err
		}
	}
	return nil
}

func (m *Integral) QueryCount(where interface{}, args ...interface{}) (int, error) {
	mdb := db.GetMysqlDB()
	var count int
	err := mdb.Model(m).Where(where, args...).Count(&count).Error
	if err != nil {
		util.LoggerError(err)
		return 0, err
	}
	return count, nil
}

func (m *Integral) QueryGrounp(field string, group string, where interface{}, args ...interface{}) ([]Integral, error) {
	mdb := db.GetMysqlDB()
	var list = make([]Integral, 0)
	err := mdb.Model(m).Select(field).Where(where, args...).Group(group).Find(&list).Error
	if err != nil {
		util.LoggerError(err)
		return make([]Integral, 0), err
	}
	return list, nil
}
