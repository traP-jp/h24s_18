package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name          string
	Id            string
	Bio           string
	TwitterId     string
	HomeChannelId string
}

func CreateUserIfNotExist(name string, id string, bio string, twitterId string, homeChannelId string) error {
	user := User{
		Name:          name,
		Id:            id,
		Bio:           bio,
		TwitterId:     twitterId,
		HomeChannelId: homeChannelId,
	}

	if err := db.First(&User{}, "id = ?", id).Error; err == nil {
		return nil
	}

	result := db.Create(&user)
	return result.Error
}

func GetUser(id string) (User, error) {
	var user User
	result := db.First(&user, "id = ?", id)
	return user, result.Error
}
