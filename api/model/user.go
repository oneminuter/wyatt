package model

import (
	"wyatt/db"
	"wyatt/util"
)

//用户
type User struct {
	TableModel

	Account        string `json:"account" gorm:"unique"` //账号, 临时账号为 10 位随机字符串
	Password       string `json:"password"`              //密码
	RandomStr      string `json:"-"`                     //密码加密随机字符串
	Phone          string `json:"phone"`                 //手机号
	UUID           string `json:"uuid"`                  //用户标识, MD5随机字符串
	NickName       string `json:"nickName"`              //昵称
	Sex            int    `json:"sex"`                   //性别 0 未知, 1 男, 2 女
	Name           string `json:"name"`                  //姓名
	Email          string `json:"email"`                 //邮箱
	AvatarUrl      string `json:"avatarUrl"`             //头像
	Country        string `json:"country"`               //国家
	Province       string `json:"province"`              //省份
	City           string `json:"city"`                  //城市
	RegisterIp     string `json:"registerIp"`            //注册ip
	Status         int64  `json:"status" gorm:"size:4"`  //用户状态: -1 封禁用户, 0 临时用户
	IsSetedAccount bool   `json:"-"`                     //是否设置过account, 每人只能设置一次
}

func (u *User) Add() error {
	mdb := db.GetMysqlDB()
	tx := mdb.Begin()
	err := tx.Model(u).Create(u).Error
	if err != nil {
		tx.Rollback()
		util.LoggerError(err)
	}
	tx.Commit()

	return err
}

func (u *User) Update(update, where interface{}, args ...interface{}) error {
	mdb := db.GetMysqlDB()
	tx := mdb.Begin()
	err := tx.Model(u).Where(where, args...).Updates(update).Error
	if err != nil {
		tx.Rollback()
		util.LoggerError(err)
	}
	tx.Commit()

	return err
}

func (u *User) QueryList(field string, where interface{}, args ...interface{}) (list []User, err error) {
	mdb := db.GetMysqlDB()
	err = mdb.Model(u).Select(field).Where(where, args...).Find(&list).Error
	if err != nil {
		util.LoggerError(err)
		return make([]User, 0), err
	}

	return list, nil
}

func (u *User) QueryOne(field string, where interface{}, args ...interface{}) error {
	mdb := db.GetMysqlDB()
	err := mdb.Model(u).Select(field).Where(where, args).Last(u).Error
	if err != nil {
		util.LoggerError(err)
	}

	return err
}

func (u *User) QueryCount(where interface{}, args ...interface{}) (int, error) {
	mdb := db.GetMysqlDB()
	var count int
	err := mdb.Model(u).Where(where, args...).Count(&count).Error
	if err != nil {
		util.LoggerError(err)
		return 0, err
	}
	return count, nil
}
