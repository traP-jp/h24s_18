package model

import (
	"fmt"
	"github.com/traP-jp/h24s_18/gemini"
	"sort"
)

type User struct {
	model
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
	Score float64
	Tags  []TagScore
}

type TagScore struct {
	Name      string
	IsStarred bool
	Score     float64
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

	var userTagMap = make(map[string][]UserTag)

	var userScoreMap = make(map[string]float64)

	for _, userTag := range userTags {
		if _, ok := userTagMap[userTag.UserId]; !ok {
			userTagMap[userTag.UserId] = make([]UserTag, 0)
			userScoreMap[userTag.UserId] = 0
		}
		userTagMap[userTag.UserId] = append(userTagMap[userTag.UserId], userTag)
		similarity := gemini.CosineSimilarity(tagEmb, tagsMap[userTag.TagName])

		score := 0.0
		switch {
		case similarity > 0:
			score = similarity
		case similarity < 0:
			score = 0
		}
		userScoreMap[userTag.UserId] += score
	}

	res := make([]RecommendedUser, 0)

	for _, user := range users {
		if _, ok := userTagMap[user.Id]; !ok {
			continue
		}

		tagRes := make([]TagScore, 0)
		for _, tag := range userTagMap[user.Id] {
			tagRes = append(tagRes, TagScore{Name: tag.TagName, Score: gemini.Dot(tagEmb, tagsMap[tag.TagName]), IsStarred: tag.IsStarred})
		}

		sort.Slice(tagRes, func(i, j int) bool {
			return tagRes[i].Score > tagRes[j].Score
		})

		res = append(res, RecommendedUser{
			User: user, Score: userScoreMap[user.Id], Tags: tagRes})
	}

	// sort
	sort.Slice(res, func(i, j int) bool {
		return res[i].Score > res[j].Score
	})

	return res, nil
}
