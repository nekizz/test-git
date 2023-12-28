package main

import (
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

func main() {
	db, err := New(mysql.Open("test20:secret@tcp(localhost:3306)/test20?charset=utf8mb4&parseTime=true&loc=UTC&autocommit=true"), false)
	if err != nil {
		fmt.Println(err)
	}

	var user User

	err = db.Model(&User{}).Preload("Address").Preload("Address.State").First(&user, 1).Error
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(user.ID, user.Address.ID)
	user.Address.City = "HAHA"
	err = db.Model(&User{}).Session(&gorm.Session{FullSaveAssociations: true}).Where("id = ?", 1).Updates(&user).Error
	if err != nil {
		fmt.Println(err)
	}

	UJson, _ := json.Marshal(user)

	fmt.Println(string(UJson))
}

func New(dialect gorm.Dialector, enableLog bool) (*gorm.DB, error) {
	orm, err := gorm.Open(dialect, &gorm.Config{})
	if nil != err {
		return nil, err
	}

	sqlDB, err := orm.DB()
	if nil != err {
		panic(err)
	}

	// TODO: use config for these values
	sqlDB.SetConnMaxLifetime(300 * time.Minute)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(15)

	return orm, nil
}

type Model struct {
	ID        uint `gorm:"primaryKey;type:int;autoIncrement"`
	CreatedBy uint `gorm:"type:int"`
	UpdatedBy uint `gorm:"type:int"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type User struct {
	Model
	Name          string
	DateOfBirth   *time.Time
	Email         string
	ContactNumber string
	VerifiedAt    *time.Time
	Budget        float64
	Status        int
	AddressID     uint
	Address       *Address
}

type Address struct {
	Model
	Address  string
	City     string
	Postcode string
	StateID  uint
	State    State
}

type State struct {
	Model
	Name string
	Code string
}

// Goodbye function for say goodbye
type Asoda struct {
	Say uint
	Bye uint
}

// Hello a function for say hello
type Hello struct {
	Say string
	Bye string
}

type Hate struct {
	Name string
	Age  uint
}

type Bello struct {
	Love uint
	Hate uint
}
