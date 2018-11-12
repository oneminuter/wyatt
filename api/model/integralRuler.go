package model

import (
	"time"
	"wyatt/api/constant"
	"wyatt/db"
	"wyatt/util"
)

//积分获取和消费规则，定义每个操作对应获取的积分或者消耗的积分
type IntegralRuler struct {
	TableModel

	Operate   string  `json:"operate" sql:"index" gorm:"size:50"` //操作字符串，1 发表文章，2 参与话题并发表评论...
	Integral  int     `json:"integral"`                           //本次操作对应积分，小于0为消耗的积分，大于0为获得的积分
	SpeedRate float64 `json:"speedRate" gorm:"default:1"`         //获得成长值的倍率，一般情况下，获得1积分同时可获得1成长值，特殊活动可得到不同倍率的成长值
}

func (m *IntegralRuler) BeforeCreate() (err error) {
	m.FlowId = time.Now().Unix()
	return
}

func (m *IntegralRuler) Add() error {
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

func (m *IntegralRuler) Delete(where interface{}, args ...interface{}) error {
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

func (m *IntegralRuler) Update(update, where interface{}, args ...interface{}) error {
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

func (m *IntegralRuler) QueryList(field string, page, limit int, where interface{}, args ...interface{}) ([]IntegralRuler, error) {
	if 0 > page {
		page = 0
	}
	if 0 > limit || constant.MAX_QUERY_COUNT < limit {
		limit = constant.MAX_QUERY_COUNT
	}

	mdb := db.GetMysqlDB()
	var list = make([]IntegralRuler, 0)
	err := mdb.Model(m).Select(field).Where(where, args...).Offset(page * limit).Limit(limit).Find(&list).Error
	if err != nil {
		util.LoggerError(err)
		return make([]IntegralRuler, 0), err
	}
	return list, nil
}

func (m *IntegralRuler) QueryOne(field string, where interface{}, args ...interface{}) error {
	mdb := db.GetMysqlDB()
	err := mdb.Model(m).Select(field).Where(where, args...).Last(m).Error
	if err != nil {
		util.LoggerError(err)
		return err
	}

	return nil
}

func (m *IntegralRuler) QueryCount(where interface{}, args ...interface{}) (int, error) {
	mdb := db.GetMysqlDB()
	var count int
	err := mdb.Model(m).Where(where, args...).Count(&count).Error
	if err != nil {
		util.LoggerError(err)
		return 0, err
	}
	return count, nil
}

func (m *IntegralRuler) QueryGrounp(field string, group string, where interface{}, args ...interface{}) ([]IntegralRuler, error) {
	mdb := db.GetMysqlDB()
	var list = make([]IntegralRuler, 0)
	err := mdb.Model(m).Select(field).Where(where, args...).Group(group).Find(&list).Error
	if err != nil {
		util.LoggerError(err)
		return make([]IntegralRuler, 0), err
	}
	return list, nil
}
