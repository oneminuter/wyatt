package model

import (
	"time"
	"wyatt/api/constant"
	"wyatt/db"
	"wyatt/util"
)

//用户
type User struct {
	TableModel

	Account        string `json:"account" gorm:"unique" sql:"index"` //账号, 临时账号为 10 位随机字符串
	Password       string `json:"password"`                          //密码
	RandomStr      string `json:"-"`                                 //密码加密随机字符串
	Phone          string `json:"phone" gorm:"size:30" sql:"index"`  //手机号
	UUID           string `json:"uuid" gorm:"size:50"`               //用户标识, MD5随机字符串
	Nickname       string `json:"nickname" gorm:"size:50"`           //昵称
	Sex            int    `json:"sex" gorm:"size:4"`                 //性别 0 未知, 1 男, 2 女
	Name           string `json:"name" gorm:"size:30"`               //姓名
	Email          string `json:"email" gorm:"size:50"`              //邮箱
	AvatarUrl      string `json:"avatarUrl"`                         //头像
	Country        string `json:"country" gorm:"size:50"`            //国家
	Province       string `json:"province" gorm:"size:50"`           //省份
	City           string `json:"city" gorm:"size:50"`               //城市
	RegisterIp     string `json:"registerIp" gorm:"size:30"`         //注册ip
	Status         int64  `json:"status" gorm:"size:4"`              //用户状态: -1 封禁用户, 0 临时用户 1或空位正常用户
	IsSetedAccount bool   `json:"-"`                                 //是否设置过account, 每人只能设置一次
	Signature      string `json:"signature"`                         //个性签名
}

func (m *User) BeforeCreate() (err error) {
	m.FlowId = time.Now().Unix()
	return
}

func (m *User) Add() error {
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

func (m *User) Delete(where interface{}, args ...interface{}) error {
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

func (m *User) Update(update, where interface{}, args ...interface{}) error {
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

func (m *User) QueryList(field string, page, limit int, where interface{}, args ...interface{}) ([]User, error) {
	if 0 > page {
		page = 0
	}
	if 0 > limit || constant.MAX_QUERY_COUNT < limit {
		limit = constant.MAX_QUERY_COUNT
	}

	mdb := db.GetMysqlDB()
	var list = make([]User, 0)
	err := mdb.Model(m).Select(field).Where(where, args...).Offset(page * limit).Limit(limit).Find(&list).Error
	if err != nil {
		util.LoggerError(err)
		return make([]User, 0), err
	}
	return list, nil
}

func (m *User) QueryOne(field string, where interface{}, args ...interface{}) error {
	mdb := db.GetMysqlDB()
	err := mdb.Model(m).Select(field).Where(where, args...).Last(m).Error
	if err != nil {
		util.LoggerError(err)
		return err
	}

	return nil
}

func (m *User) QueryCount(where interface{}, args ...interface{}) (int, error) {
	mdb := db.GetMysqlDB()
	var count int
	err := mdb.Model(m).Where(where, args...).Count(&count).Error
	if err != nil {
		util.LoggerError(err)
		return 0, err
	}
	return count, nil
}

func (m *User) QueryGrounp(field string, group string, where interface{}, args ...interface{}) ([]User, error) {
	mdb := db.GetMysqlDB()
	var list = make([]User, 0)
	err := mdb.Model(m).Select(field).Where(where, args...).Group(group).Find(&list).Error
	if err != nil {
		util.LoggerError(err)
		return make([]User, 0), err
	}
	return list, nil
}
