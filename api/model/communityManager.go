package model

import (
	"wyatt/db"
	"wyatt/util"
)

//社区人员管理
type CommunityManager struct {
	TableModel

	CommunityId int64 `json:"communityId"` // community 的主键
	UserId      int64 `json:"userId"`      //用户id
	Role        int   `json:"role"`        // -1 封禁， 0 普通成员，1 管理员  其他为游客，封禁状态不能参与该社区的话题和发言
}

func (cm *CommunityManager) Add() error {
	mdb := db.GetMysqlDB()
	tx := mdb.Begin()
	defer tx.Commit()

	err := tx.Create(cm).Error
	if err != nil {
		tx.Rollback()
		util.LoggerError(err)
		return err
	}
	return nil
}
func (cm *CommunityManager) QueryList(field string, where interface{}, args ...interface{}) ([]CommunityManager, error) {
	mdb := db.GetMysqlDB()
	var list []CommunityManager
	err := mdb.Model(cm).Where(where, args...).Find(&list).Error
	if err != nil {
		util.LoggerError(err)
		return make([]CommunityManager, 0), err
	}
	return list, nil
}

func (cm *CommunityManager) Delete(where interface{}, args ...interface{}) error {
	mdb := db.GetMysqlDB()
	tx := mdb.Begin()
	defer tx.Commit()

	err := tx.Model(cm).Where(where, args...).Delete(cm).Error
	if err != nil {
		tx.Rollback()
		util.LoggerError(err)
		return err
	}
	return nil
}
