package versions

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

func Version20221110000009(tx *gorm.DB) error {
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

	return tx.AutoMigrate(&Hotel{}, &User{})
}
