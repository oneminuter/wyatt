package model

import (
	"time"
	"wyatt/api/constant"
	"wyatt/db"
	"wyatt/util"
)

//消息中心
type Message struct {
	TableModel

	UserId   int64  `json:"userId" sql:"index"`     //用户id
	MsgType  string `json:"msgType" gorm:"size:30"` //消息类型， system:系统消息，custom:自定义消息，可以根据该字段来判断是不是定向消息
	Content  string `json:"content"`                //消息内容
	IsViewed int    `json:"isViewed" gorm:"size:4"` //是否查看过
}

func (m *Message) BeforeCreate() (err error) {
	m.FlowId = time.Now().Unix()
	return
}

func (m *Message) Add() error {
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

func (m *Message) Delete(where interface{}, args ...interface{}) error {
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

func (m *Message) Update(update, where interface{}, args ...interface{}) error {
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

func (m *Message) QueryList(field string, page, limit int, where interface{}, args ...interface{}) ([]Message, error) {
	if 0 > page {
		page = 0
	}
	if 0 > limit || constant.MAX_QUERY_COUNT < limit {
		limit = constant.MAX_QUERY_COUNT
	}

	mdb := db.GetMysqlDB()
	var list = make([]Message, 0)
	err := mdb.Model(m).Select(field).Where(where, args...).Offset(page * limit).Limit(limit).Find(&list).Error
	if err != nil {
		util.LoggerError(err)
		return make([]Message, 0), err
	}
	return list, nil
}

func (m *Message) QueryOne(field string, where interface{}, args ...interface{}) error {
	mdb := db.GetMysqlDB()
	err := mdb.Model(m).Select(field).Where(where, args...).Last(m).Error
	if err != nil {
		util.LoggerError(err)
		return err
	}

	return nil
}

func (m *Message) QueryCount(where interface{}, args ...interface{}) (int, error) {
	mdb := db.GetMysqlDB()
	var count int
	err := mdb.Model(m).Where(where, args...).Count(&count).Error
	if err != nil {
		util.LoggerError(err)
		return 0, err
	}
	return count, nil
}

func (m *Message) QueryGrounp(field string, group string, where interface{}, args ...interface{}) ([]Message, error) {
	mdb := db.GetMysqlDB()
	var list = make([]Message, 0)
	err := mdb.Model(m).Select(field).Where(where, args...).Group(group).Find(&list).Error
	if err != nil {
		util.LoggerError(err)
		return make([]Message, 0), err
	}
	return list, nil
}
