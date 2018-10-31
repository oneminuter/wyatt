package model

import (
	"wyatt/api/constant"
	"wyatt/db"
	"wyatt/util"
)

//评论
type Comment struct {
	TableModel
	CID       int64  `json:"cid"`                               //评论id, 10为数字
	UserId    int64  `json:"userId"`                            //发送者用户id
	Content   string `json:"content" gorm:"type:varchar(5000)"` //评论内容
	Classify  string `json:"classify"`                          //分类：社区话题 还是其他的评论，和表关联
	AritcleId int64  `json:"aritcleId"`                         //文章id, 这条评论属于哪个文章或者哪个话题的id
	ReplyCID  int64  `json:"replyCid"`                          //回复评论id，主键，为空则不是回复
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
	err := mdb.Model(c).Select(field).Where(where, args...).Find(&list).Error
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
