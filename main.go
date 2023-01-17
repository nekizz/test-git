package main

import (
	"fmt"
	"test/connection"
	"test/model"
)

//user := &model.User{
//Model: gorm.Model{},
//Name:  "Minh",
//Age:   12,
//}
//query := connection.Conn.Model(&model.User{}).Create(user)
//if err := query.Error; nil != err {
//fmt.Println(err)
//}

//hotel := &model.Hotel{
//Model: gorm.Model{},
//Name:  "Molina's hotel",
//}
//query := connection.Conn.Model(&model.Hotel{}).Create(hotel)
//if err := query.Error; nil != err {
//fmt.Println(err)
//}

func main() {
	var users model.User
	err := connection.Conn.Model(&model.User{}).Where("id = ?", 1).Association("hotel").Find(&users)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(users)
}

func CreatOne(e *model.Hotel) (*model.Hotel, error) {
	query := connection.Conn.Model(&model.Hotel{}).Create(e)
	if err := query.Error; nil != err {
		return nil, err
	}

	return e, nil
}

func DeleteOne(id int) error {
	query := connection.Conn.Delete(&model.Hotel{}, "id = ?", id)
	if err := query.Error; nil != err {
		return err
	}

	return nil
}

func FindOne(id int) (*model.Hotel, error) {
	var hotel model.Hotel

	query := connection.Conn.Model(&model.Hotel{}).Where("id = ?", id).Find(&hotel)
	if err := query.Error; nil != err {
		return nil, err
	}

	return &hotel, nil
}
