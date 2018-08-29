package model

import (
	"time"
	"wyatt/db"

	_ "github.com/jinzhu/gorm"
)

type TableModel struct {
	ID        int64 `json:"-" gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func init() {
	mdb := db.GetMysqlDB()
	mdb.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&Comment{}, &Community{}, &JoinedCommunity{}, &Message{}, &Topic{}, &Zan{}, &User{})
}
