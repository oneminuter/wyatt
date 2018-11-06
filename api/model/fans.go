package model

import (
	"time"
	"wyatt/api/constant"
	"wyatt/db"
	"wyatt/util"
)

//关注粉丝
type Fans struct {
	TableModel
	UserId     int64 `json:"userId"`     //用户id
	FansUserId int64 `json:"fansUserId"` //订阅者，粉丝的用户id
}

func (bc *Fans) BeforeCreate() (err error) {
	bc.FlowId = time.Now().Unix()
	return
}

func (f *Fans) Add() error {
	mdb := db.GetMysqlDB()
	tx := mdb.Begin()
	defer tx.Commit()
	err := mdb.Create(f).Error
	if err != nil {
		util.LoggerError(err)
		tx.Rollback()
		return err
	}
	return nil
}

func (f *Fans) Delete(where interface{}, args ...interface{}) error {
	mdb := db.GetMysqlDB()
	tx := mdb.Begin()
	defer tx.Commit()
	err := tx.Model(f).Where(where, args...).Delete(f).Error
	if err != nil {
		util.LoggerError(err)
		tx.Rollback()
		return err
	}
	return nil
}

func (f *Fans) QueryList(field string, page int, limit int, where interface{}, args ...interface{}) ([]Fans, error) {
	if 0 > page {
		page = 0
	}
	if 0 > limit || constant.MAX_QUERY_COUNT < limit {
		limit = constant.MAX_QUERY_COUNT
	}

	mdb := db.GetMysqlDB()
	var list = make([]Fans, 0)
	err := mdb.Model(f).Select(field).Where(where, args...).Offset(page * limit).Limit(limit).Find(&list).Error
	if err != nil {
		util.LoggerError(err)
		return make([]Fans, 0), err
	}
	return list, err
}
