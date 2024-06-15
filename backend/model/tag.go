package model

import (
	"gorm.io/gorm"
)

type Tag struct {
	gorm.Model
	Name  string
	Value string
}

func CreateTag(name string, value string) error {
	tag := Tag{Name: name, Value: value}
	result := db.Create(&tag)
	return result.Error
}

func GetTag(value string) (Tag, error) {
	var tag Tag
	result := db.First(&tag, "value = ?", value)
	return tag, result.Error
}
