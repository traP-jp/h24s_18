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

func GetTag(name string) (Tag, error) {
	var tag Tag
	result := db.First(&tag, "value = ?", name)
	return tag, result.Error
}

func GetTags() ([]Tag, error) {
	var tags []Tag
	result := db.Find(&tags)
	return tags, result.Error
}

func GetAllTagsName() ([]string, error) {
	tags, err := GetTags()
	if err != nil {
		return nil, err
	}
	names := make([]string, len(tags))
	for i, tag := range tags {
		names[i] = tag.Name
	}
	return names, nil
}


 