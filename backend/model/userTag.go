package model

type UserTag struct {
	model
	UserId    string `gorm:"primaryKey"`
	TagName   string `gorm:"primaryKey"`
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
