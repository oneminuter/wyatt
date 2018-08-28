package model

import (
	"time"

	_ "github.com/jinzhu/gorm"
)

type TableModel struct {
	ID        int64 `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

