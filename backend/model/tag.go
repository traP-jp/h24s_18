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
	Tag := Tag{Name: name, Value: value}
	result := db.Create(&Tag)
	return result.Error
}

func GetTag(value string) (Tag, error) {
	var Tag Tag
	result := db.First(&Tag, "value = ?", value)
	return Tag, result.Error
}
