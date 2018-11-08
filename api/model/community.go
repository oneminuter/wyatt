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
	CreatorId int64  `json:"creatorId"`                   //创建者id
	Status    int    `json:"status" gorm:"size:4"`        //社区状态: -1 封禁下架, 0 申请中, 1 正常, 2 解散删除
}

func (bc *Community) BeforeCreate() (err error) {
	bc.FlowId = time.Now().Unix()
	return
}

func (c *Community) QueryOne(field string, where interface{}, args ...interface{}) error {
	mdb := db.GetMysqlDB()
	err := mdb.Model(c).Select(field).Where(where, args...).Last(c).Error
	if err != nil {
		util.LoggerError(err)
		return err
	}
	return nil
}

func (c *Community) QueryList(field string, page int, limit int, where interface{}, args ...interface{}) ([]Community, error) {
	if 0 > page {
		page = 0
	}
	if 0 > limit || constant.MAX_QUERY_COUNT < limit {
		limit = constant.MAX_QUERY_COUNT
	}

	mdb := db.GetMysqlDB()
	var list []Community
	err := mdb.Model(c).Select(field).Where(where, args...).Offset(page * limit).Limit(limit).Find(&list).Error
	if err != nil {
		util.LoggerError(err)
		return make([]Community, 0), err
	}
	return list, nil
}

func (c *Community) Add() error {
	mdb := db.GetMysqlDB()
	tx := mdb.Begin()
	defer tx.Commit()

	err := tx.Create(c).Error
	if err != nil {
		tx.Rollback()
		util.LoggerError(err)
		return err
	}
	return nil
}

/*
更新
update: 可以是结构体 或者是 map
*/
func (c *Community) Update(update interface{}, where interface{}, args ...interface{}) error {
	mdb := db.GetMysqlDB()
	tx := mdb.Begin()
	defer tx.Commit()

	err := tx.Model(c).Where(where, args...).Updates(update).Error
	if err != nil {
		tx.Rollback()
		util.LoggerError(err)
		return err
	}
	return nil
}

func (c *Community) QueryCount(where interface{}, args ...interface{}) int {
	mdb := db.GetMysqlDB()
	var count int
	err := mdb.Model(c).Where(where, args...).Count(&count).Error
	if err != nil {
		util.LoggerError(err)
		return 0
	}
	return count
}
