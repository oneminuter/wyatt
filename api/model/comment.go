package model

import (
	"time"
	"wyatt/api/constant"
	"wyatt/db"
	"wyatt/util"
)

//评论
type Comment struct {
	TableModel
	CreatorId    int64  `json:"creatorId"`                         //发送者用户id
	Content      string `json:"content" gorm:"type:varchar(5000)"` //评论内容
	SourceFlowId string `json:"sourceFlowId"`                      //被评论对象的完整流水号
	ReplyCId     string `json:"replyCid"`                          //被回复评论的完整流水号
}

func (bc *Comment) BeforeCreate() (err error) {
	bc.FlowId = time.Now().Unix()
	return
}

func (c *Comment) Add() error {
	mdb := db.GetMysqlDB()
	tx := mdb.Begin()
	defer tx.Commit()
	err := tx.Create(c).Error
	if err != nil {
		util.LoggerError(err)
		tx.Rollback()
		return err
	}
	return nil
}
func (c *Comment) Delete(where interface{}, args ...interface{}) error {
	mdb := db.GetMysqlDB()
	tx := mdb.Begin()
	defer tx.Commit()
	err := tx.Model(c).Where(where, args...).Delete(c).Error
	if err != nil {
		util.LoggerError(err)
		tx.Rollback()
		return err
	}
	return nil
}

func (c *Comment) QueryList(field string, page int, limit int, where interface{}, args ...interface{}) ([]Comment, error) {
	if 0 > page {
		page = 0
	}
	if 0 > limit || constant.MAX_QUERY_COUNT < limit {
		limit = constant.MAX_QUERY_COUNT
	}

	mdb := db.GetMysqlDB()
	var list = make([]Comment, 0)
	err := mdb.Model(c).Select(field).Where(where, args...).Offset(page * limit).Limit(limit).Find(&list).Error
	if err != nil {
		util.LoggerError(err)
		return make([]Comment, 0), err
	}
	return list, nil
}

func (c *Comment) QueryOne(field string, where interface{}, args ...interface{}) error {
	mdb := db.GetMysqlDB()
	err := mdb.Model(c).Select(field).Where(where, args...).Last(c).Error
	if err != nil {
		util.LoggerError(err)
		return err
	}
	return nil
}
