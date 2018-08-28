package model

import (
	"time"
	"wyatt/db"

	_ "github.com/jinzhu/gorm"
)

type TableModel struct {
	ID        int64 `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

type Model interface {
	Add() error
	QueryOne(field string, where interface{}, args ...interface{}) (Model, error)
	QueryList(field string, where interface{}, args ...interface{}) ([]Model, error)
}

func init() {
	mdb := db.GetMysqlDB()
	defer mdb.Close()
	mdb.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Comment{}, &Community{}, &JoinedCommunity{}, &Message{}, &Topic{}, &Zan{}, &User{})
}
