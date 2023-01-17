package model

import (
	"gorm.io/gorm"
)

type Hotel struct {
	gorm.Model
	Name string
}

type User struct {
	gorm.Model
	Name  string
	Age   int
	Hotel []*Hotel `gorm:"many2many:whislist;"`
}
