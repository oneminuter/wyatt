package model

import (
	"time"
	"wyatt/api/constant"
	"wyatt/db"
	"wyatt/util"
)

//故事的每条内容
type StoryContent struct {
	TableModel

	StoryId string  `json:"storyId"` //属于哪个故事的id
	Type    string  `json:"type"`    //1 角色对白，2 旁白
	RoleId  int64   `json:"roleId"`  //角色id, 如果是旁白，该字段为空
	Context string  `json:"context"` //内容
	Order   float64 `json:"order"`   //权重
}

func (m *StoryContent) BeforeCreate() (err error) {
	m.FlowId = time.Now().Unix()
	return
}

func (m *StoryContent) Add() error {
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

func (m *StoryContent) Delete(where interface{}, args ...interface{}) error {
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

func (m *StoryContent) Update(update, where interface{}, args ...interface{}) error {
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

//查询列表，按 order 排序
func (m *StoryContent) QueryList(field string, page, limit int, where interface{}, args ...interface{}) ([]StoryContent, error) {
	if 0 > page {
		page = 0
	}
	if 0 > limit || constant.MAX_QUERY_COUNT < limit {
		limit = constant.MAX_QUERY_COUNT
	}

	mdb := db.GetMysqlDB()
	var list = make([]StoryContent, 0)
	err := mdb.Model(m).Select(field).Where(where, args...).Offset(page * limit).Limit(limit).Order("order asc").Find(&list).Error
	if err != nil {
		util.LoggerError(err)
		return make([]StoryContent, 0), err
	}
	return list, nil
}

func (m *StoryContent) QueryOne(field string, where interface{}, args ...interface{}) error {
	mdb := db.GetMysqlDB()
	err := mdb.Model(m).Select(field).Where(where, args...).Last(m).Error
	if err != nil {
		util.LoggerError(err)
		return err
	}

	return nil
}

func (m *StoryContent) QueryCount(where interface{}, args ...interface{}) (int, error) {
	mdb := db.GetMysqlDB()
	var count int
	err := mdb.Model(m).Where(where, args...).Count(&count).Error
	if err != nil {
		util.LoggerError(err)
		return 0, err
	}
	return count, nil
}

func (m *StoryContent) QueryGrounp(field string, group string, where interface{}, args ...interface{}) ([]StoryContent, error) {
	mdb := db.GetMysqlDB()
	var list = make([]StoryContent, 0)
	err := mdb.Model(m).Select(field).Where(where, args...).Group(group).Find(&list).Error
	if err != nil {
		util.LoggerError(err)
		return make([]StoryContent, 0), err
	}
	return list, nil
}
