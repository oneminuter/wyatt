package model

import (
	"time"
	"wyatt/db"
	"wyatt/util"
)

//图片 - 图片路劲由 Path + FileName 组成
type Img struct {
	TableModel

	Path     string `json:"path"`     //路径 或者 图片oss域名
	FileName string `json:"fileName"` //图片图片名
}

func (bc *Img) BeforeCreate() (err error) {
	bc.FlowId = time.Now().Unix()
	return
}

func (i *Img) Add() error {
	mdb := db.GetMysqlDB()
	tx := mdb.Begin()
	defer tx.Commit()

	err := tx.Create(i).Error
	if err != nil {
		tx.Rollback()
		util.LoggerError(err)
		return err
	}
	return nil
}
