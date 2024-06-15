package model

import (
	"gorm.io/gorm"
)

type UserTag struct {
	gorm.Model
	UserId    string
	TagName   string
	IsStarred bool
}

func CreateUserTag(userId string, tagName string, isStarred bool) error {
	userTag := UserTag{UserId: userId, TagName: tagName, IsStarred: isStarred}
	result := db.Create(&userTag)
	return result.Error
}

func GetUserTagsByUserId(userId string) ([]UserTag, error) {
	var userTag []UserTag
	result := db.Find(&userTag, "userId = ?", userId)
	return userTag, result.Error
}
