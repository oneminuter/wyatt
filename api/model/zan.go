package model

import (
	"time"
	"wyatt/db"
	"wyatt/util"
)

//赞记录
type Zan struct {
	TableModel
	UserId       int64  `json:"userId"` //触发者用户id
	OwnerId      int64  `json:"ownerId"`
	SourceFlowId string `json:"sourceFlowId"` //赞来源流水号，完整流水号
}

func (bc *Zan) BeforeCreate() (err error) {
	bc.FlowId = time.Now().Unix()
	return
}

func (z *Zan) Add() error {
	mdb := db.GetMysqlDB()
	tx := mdb.Begin()
	defer tx.Commit()
	err := tx.Create(z).Error
	if err != nil {
		util.LoggerError(err)
		tx.Rollback()
		return err
	}
	return nil
}

func (z *Zan) Delete(where interface{}, args ...interface{}) error {
	mdb := db.GetMysqlDB()
	tx := mdb.Begin()
	defer tx.Commit()
	err := tx.Model(z).Where(where, args...).Delete(z).Error
	if err != nil {
		util.LoggerError(err)
		tx.Rollback()
		return err
	}
	return nil
}

func (z *Zan) QueryOne(field string, where interface{}, args ...interface{}) error {
	mdb := db.GetMysqlDB()
	err := mdb.Model(z).Select(field).Where(where, args...).Last(z).Error
	if err != nil {
		util.LoggerError(err)
		return err
	}
	return nil
}

func (z *Zan) QueryList(field string, where interface{}, args ...interface{}) ([]Zan, error) {
	mdb := db.GetMysqlDB()
	var list = make([]Zan, 0)
	err := mdb.Model(z).Select(field).Where(where, args...).Find(&list).Error
	if err != nil {
		util.LoggerError(err)
		return make([]Zan, 0), err
	}
	return list, nil
}

func (z *Zan) QueryCount(where interface{}, args ...interface{}) (int, error) {
	mdb := db.GetMysqlDB()
	var count int
	err := mdb.Model(z).Where(where, args...).Count(&count).Error
	if err != nil {
		util.LoggerError(err)
		return 0, err
	}
	return count, nil
}
