package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name string
	Id   string
}

func CreateUser(name string, id string) error {
	user := User{Name: name, Id: id}
	result := db.Create(&user)
	return result.Error
}

func GetUser(id string) (User, error) {
	var user User
	result := db.First(&user, "id = ?", id)
	return user, result.Error
}
