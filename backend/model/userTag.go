package model

type UserTag struct {
	model
	UserId    string `gorm:"primaryKey"`
	TagName   string `gorm:"primaryKey"`
	IsStarred bool
}

// UserTagを作成
func CreateUserTag(userId string, tagName string, isStarred bool) error {
	userTag := UserTag{UserId: userId, TagName: tagName, IsStarred: isStarred}
	result := db.Create(&userTag)
	return result.Error
}

// UserIdからそのユーザの持つUserTagを生成
func GetUserTagsByUserId(userId string) ([]UserTag, error) {
	var userTag []UserTag
	result := db.Find(&userTag, "user_id = ?", userId)
	return userTag, result.Error
}

// UserTagを削除
func DeleteUserTag(tagName string) error {
	var userTag []UserTag
	result := db.Delete(&userTag, "tag_name = ?", tagName)
	return result.Error
}
