package model

import (
	"time"
	"wyatt/api/constant"
	"wyatt/db"
	"wyatt/util"

	"github.com/jinzhu/gorm"
)

//用户积分等级规则
type GrowthlLevel struct {
	TableModel
	Level int `json:"level"` //等级
	Start int `json:"start"` //每个等级的开始值
}

//等级初始化数据
var growthLevelInitDataList = []*GrowthlLevel{
	{TableModel: TableModel{ID: 1}, Level: 0, Start: 0},
	{TableModel: TableModel{ID: 2}, Level: 1, Start: 100},
	{TableModel: TableModel{ID: 3}, Level: 2, Start: 300},
	{TableModel: TableModel{ID: 4}, Level: 3, Start: 600},
	{TableModel: TableModel{ID: 5}, Level: 4, Start: 1000},
	{TableModel: TableModel{ID: 6}, Level: 5, Start: 1500},
	{TableModel: TableModel{ID: 7}, Level: 6, Start: 2100},
	{TableModel: TableModel{ID: 8}, Level: 7, Start: 2800},
	{TableModel: TableModel{ID: 9}, Level: 8, Start: 3600},
	{TableModel: TableModel{ID: 10}, Level: 9, Start: 4500},
	{TableModel: TableModel{ID: 11}, Level: 10, Start: 5500},
	{TableModel: TableModel{ID: 12}, Level: 11, Start: 6600},
	{TableModel: TableModel{ID: 13}, Level: 12, Start: 7800},
	{TableModel: TableModel{ID: 14}, Level: 13, Start: 9100},
	{TableModel: TableModel{ID: 15}, Level: 14, Start: 10500},
	{TableModel: TableModel{ID: 16}, Level: 15, Start: 12000},
	{TableModel: TableModel{ID: 17}, Level: 16, Start: 13600},
	{TableModel: TableModel{ID: 18}, Level: 17, Start: 15300},
	{TableModel: TableModel{ID: 19}, Level: 18, Start: 17100},
	{TableModel: TableModel{ID: 20}, Level: 19, Start: 19000},
	{TableModel: TableModel{ID: 21}, Level: 20, Start: 21000},
	{TableModel: TableModel{ID: 22}, Level: 21, Start: 23100},
	{TableModel: TableModel{ID: 23}, Level: 22, Start: 25300},
	{TableModel: TableModel{ID: 24}, Level: 23, Start: 27600},
	{TableModel: TableModel{ID: 25}, Level: 24, Start: 30000},
	{TableModel: TableModel{ID: 26}, Level: 25, Start: 32500},
	{TableModel: TableModel{ID: 27}, Level: 26, Start: 35100},
	{TableModel: TableModel{ID: 28}, Level: 27, Start: 37800},
	{TableModel: TableModel{ID: 29}, Level: 28, Start: 40600},
	{TableModel: TableModel{ID: 30}, Level: 29, Start: 43500},
	{TableModel: TableModel{ID: 31}, Level: 30, Start: 46500},
}

//初始化等级数据
func (m *GrowthlLevel) initGrowthLevel(mdb *gorm.DB) {
	for _, v := range growthLevelInitDataList {
		err := mdb.Where("id = ?", v.ID).FirstOrCreate(v).Error
		if err != nil {
			util.LoggerError(err)
		}
	}
}

func (m *GrowthlLevel) BeforeCreate() (err error) {
	m.FlowId = time.Now().Unix()
	return
}

func (m *GrowthlLevel) Add() error {
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

func (m *GrowthlLevel) Delete(where interface{}, args ...interface{}) error {
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

func (m *GrowthlLevel) Update(update, where interface{}, args ...interface{}) error {
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

func (m *GrowthlLevel) QueryList(field string, page, limit int, where interface{}, args ...interface{}) ([]GrowthlLevel, error) {
	if 0 > page {
		page = 0
	}
	if 0 > limit || constant.MAX_QUERY_COUNT < limit {
		limit = constant.MAX_QUERY_COUNT
	}

	mdb := db.GetMysqlDB()
	var list = make([]GrowthlLevel, 0)
	err := mdb.Model(m).Select(field).Where(where, args...).Offset(page * limit).Limit(limit).Find(&list).Error
	if err != nil {
		util.LoggerError(err)
		return make([]GrowthlLevel, 0), err
	}
	return list, nil
}

func (m *GrowthlLevel) QueryOne(field string, where interface{}, args ...interface{}) error {
	mdb := db.GetMysqlDB()
	err := mdb.Model(m).Select(field).Where(where, args...).Last(m).Error
	if err != nil {
		util.LoggerError(err)
		return err
	}

	return nil
}

func (m *GrowthlLevel) QueryCount(where interface{}, args ...interface{}) (int, error) {
	mdb := db.GetMysqlDB()
	var count int
	err := mdb.Model(m).Where(where, args...).Count(&count).Error
	if err != nil {
		util.LoggerError(err)
		return 0, err
	}
	return count, nil
}

func (m *GrowthlLevel) QueryGrounp(field string, group string, where interface{}, args ...interface{}) ([]GrowthlLevel, error) {
	mdb := db.GetMysqlDB()
	var list = make([]GrowthlLevel, 0)
	err := mdb.Model(m).Select(field).Where(where, args...).Group(group).Find(&list).Error
	if err != nil {
		util.LoggerError(err)
		return make([]GrowthlLevel, 0), err
	}
	return list, nil
}
