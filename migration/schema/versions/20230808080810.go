package versions

import (
	"time"

	"gorm.io/gorm"
)

type Model struct {
	ID        uint `gorm:"primaryKey;type:int;autoIncrement"`
	CreatedBy uint `gorm:"type:int"`
	UpdatedBy uint `gorm:"type:int"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func Version20230808080810(tx *gorm.DB) error {
	type State struct {
		Model
		Name string `gorm:"type:char(32);not null"`
		Code string `gorm:"type:char(3);not null"`
	}

	type Address struct {
		Model
		Address  string `gorm:"type:char(128);not null"`
		City     string `gorm:"type:char(128);not null"`
		Postcode string `gorm:"type:char(32);not null"`
		StateID  uint   `gorm:"type:int;not null"`
		State    State  `gorm:"foreignKey:StateID;references:ID;constraint:OnUpdate:RESTRICT,OnDelete:CASCADE;"`
	}

	type User struct {
		Model
		Name          string `gorm:"type:varchar(255);not null"`
		DateOfBirth   *time.Time
		Email         string `gorm:"unique;type:varchar(255);not null"`
		ContactNumber string `gorm:"type:varchar(16)"`
		VerifiedAt    *time.Time
		Budget        float64
		Status        int     `gorm:"type:tinyint;not null"`
		AddressID     uint    `gorm:"type:int"`
		Address       Address `gorm:"foreignKey:AddressID;references:ID;constraint:OnUpdate:RESTRICT,OnDelete:CASCADE;"`
	}

	return tx.AutoMigrate(&User{}, &Address{}, &State{})
}

func Rollback20230808080810(tx *gorm.DB) error {

	return nil
}
