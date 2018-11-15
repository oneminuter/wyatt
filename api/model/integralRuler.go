package model

import (
	"time"
	"wyatt/api/constant"
	"wyatt/db"
	"wyatt/util"

	"github.com/jinzhu/gorm"
)

//积分获取和消费规则，定义每个操作对应获取的积分或者消耗的积分
type IntegralRuler struct {
	TableModel

	Operate   string  `json:"operate" sql:"index" gorm:"size:50"` //操作字符串，1 发表文章，2 参与话题并发表评论...
	Integral  int     `json:"integral"`                           //本次操作对应积分，小于0为消耗的积分，大于0为获得的积分
	SpeedRate float64 `json:"speedRate" gorm:"default:1"`         //获得成长值的倍率，一般情况下，获得1积分同时可获得1成长值，特殊活动可得到不同倍率的成长值
}

//获取和消费积分的一些操作
var (
	OPT_REGISTER         = "register"        //注册
	OPT_LOGIN            = "login"           //登录
	OPT_CREATE_COMMUNITY = "createCommunity" //创建社区
	OPT_ADD_TOPIC        = "addTopic"        //添加话题
	OPT_ADD_COMMENT      = "addComment"      //添加评论
	OPT_GET_ZAN          = "getZan"          //获得赞
)

//初始化数据
var integralRulerInitDataList = []*IntegralRuler{
	{TableModel: TableModel{ID: 1}, Operate: OPT_REGISTER, Integral: 20, SpeedRate: 1},
	{TableModel: TableModel{ID: 2}, Operate: OPT_LOGIN, Integral: 5, SpeedRate: 1},
	{TableModel: TableModel{ID: 3}, Operate: OPT_CREATE_COMMUNITY, Integral: 10, SpeedRate: 1},
	{TableModel: TableModel{ID: 4}, Operate: OPT_ADD_TOPIC, Integral: 10, SpeedRate: 1},
	{TableModel: TableModel{ID: 5}, Operate: OPT_ADD_COMMENT, Integral: 5, SpeedRate: 1},
	{TableModel: TableModel{ID: 6}, Operate: OPT_GET_ZAN, Integral: 1, SpeedRate: 1},
}

//积分规则数据初始化
func (*IntegralRuler) initIntegralRuler(mdb *gorm.DB) {
	for _, v := range integralRulerInitDataList {
		err := mdb.Where("id = ?", v.ID).FirstOrCreate(v).Error
		if err != nil {
			util.LoggerError(err)
		}
	}
}

func (m *IntegralRuler) BeforeCreate() (err error) {
	m.FlowId = time.Now().Unix()
	return
}

func (m *IntegralRuler) Add() error {
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

func (m *IntegralRuler) Delete(where interface{}, args ...interface{}) error {
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

func (m *IntegralRuler) Update(update, where interface{}, args ...interface{}) error {
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

func (m *IntegralRuler) QueryList(field string, page, limit int, where interface{}, args ...interface{}) ([]IntegralRuler, error) {
	if 0 > page {
		page = 0
	}
	if 0 > limit || constant.MAX_QUERY_COUNT < limit {
		limit = constant.MAX_QUERY_COUNT
	}

	mdb := db.GetMysqlDB()
	var list = make([]IntegralRuler, 0)
	err := mdb.Model(m).Select(field).Where(where, args...).Offset(page * limit).Limit(limit).Find(&list).Error
	if err != nil {
		util.LoggerError(err)
		return make([]IntegralRuler, 0), err
	}
	return list, nil
}

func (m *IntegralRuler) QueryOne(field string, where interface{}, args ...interface{}) error {
	mdb := db.GetMysqlDB()
	err := mdb.Model(m).Select(field).Where(where, args...).Last(m).Error
	if err != nil {
		util.LoggerError(err)
		return err
	}

	return nil
}

func (m *IntegralRuler) QueryCount(where interface{}, args ...interface{}) (int, error) {
	mdb := db.GetMysqlDB()
	var count int
	err := mdb.Model(m).Where(where, args...).Count(&count).Error
	if err != nil {
		util.LoggerError(err)
		return 0, err
	}
	return count, nil
}

func (m *IntegralRuler) QueryGrounp(field string, group string, where interface{}, args ...interface{}) ([]IntegralRuler, error) {
	mdb := db.GetMysqlDB()
	var list = make([]IntegralRuler, 0)
	err := mdb.Model(m).Select(field).Where(where, args...).Group(group).Find(&list).Error
	if err != nil {
		util.LoggerError(err)
		return make([]IntegralRuler, 0), err
	}
	return list, nil
}
