package model

import (
	"fmt"
	"github.com/traP-jp/h24s_18/gemini"
	"gorm.io/gorm"
	"sort"
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

func UpdateUserBio(id string, bio string) {
	db.Model(&User{}).Where("Id =?", id).Update("Bio", bio)
}

type RecommendedUser struct {
	User  User
	Score float32
	Tags  []string
}

func RecommendUserByTag(tagName string) ([]RecommendedUser, error) {
	tagsMap := make(map[string]gemini.Embedding)
	{
		tags := make([]Tag, 0)

		result := db.Find(&tags)
		if result.Error != nil {
			return nil, result.Error
		}

		for _, tag := range tags {
			emb, err := gemini.ScanEmb(tag.Value)
			if err != nil {
				return nil, err
			}
			tagsMap[tag.Name] = *emb
		}
	}

	fmt.Printf("tagsMap: %+v\n", tagsMap)

	tagEmb, ok := tagsMap[tagName]
	if !ok {
		return nil, fmt.Errorf("tag not found")
	}

	var users []User
	{
		result := db.Find(&users)
		if result.Error != nil {
			return nil, result.Error
		}
	}

	userTags := make([]UserTag, 0)
	{
		result := db.Find(&userTags)
		if result.Error != nil {
			return nil, result.Error
		}
	}

	var userTagMap = make(map[string][]string)

	var userScoreMap = make(map[string]float32)

	for _, userTag := range userTags {
		if _, ok := userTagMap[userTag.UserId]; !ok {
			userTagMap[userTag.UserId] = make([]string, 0)
			userScoreMap[userTag.UserId] = 0
		}
		userTagMap[userTag.UserId] = append(userTagMap[userTag.UserId], userTag.TagName)
		userScoreMap[userTag.UserId] += gemini.Dot(tagEmb, tagsMap[userTag.TagName])
	}

	res := make([]RecommendedUser, 0)

	for _, user := range users {
		if _, ok := userTagMap[user.Id]; !ok {
			continue
		}
		res = append(res, RecommendedUser{User: user, Score: userScoreMap[user.Id], Tags: userTagMap[user.Id]})
	}

	// sort
	sort.Slice(res, func(i, j int) bool {
		return res[i].Score > res[j].Score
	})

	return res, nil
}
