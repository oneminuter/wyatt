package model

import (
	"wyatt/db"
	"wyatt/util"
)

//社区
type Community struct {
	TableModel

	CId       int64  `json:"cId"`       //社区号，创建时的时间戳(s)
	Logo      string `json:"logo"`      //社区logo
	Name      string `json:"name"`      //社区名
	Desc      string `json:"desc"`      //社区描述
	CreatorId int64  `json:"creatorId"` //创建者id
	Status    int    `json:"status"`    //社区状态: -1 封禁下架, 0 申请中, 1 正常, 2 解散删除
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

func (c *Community) QueryList(field string, where interface{}, args ...interface{}) ([]Community, error) {
	mdb := db.GetMysqlDB()
	var list []Community
	err := mdb.Model(c).Select(field).Where(where, args...).Find(&list).Error
	if err != nil {
		util.LoggerError(err)
		return make([]Community, 0), err
	}
	return list, nil
}
