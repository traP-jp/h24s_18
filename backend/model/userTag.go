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
	result := db.Find(&userTag, "user_id = ?", userId) // db の user_id が userId のものを取得
	return userTag, result.Error
}

// userTag につけられている isStarred を変える
func UpdateUserTags(userId string, tagName string, isStarred bool) error {
	var userTag []UserTag
	result := db.Model(&userTag).Where("user_id = ?", userId).Where("tag_name = ?", tagName).Update("is_starred", isStarred)
	return result.Error
}

// UserTagを削除
func DeleteUserTag(userId string, tagName string) error {
	var userTag []UserTag
	result := db.Delete(&userTag, "user_id = ? AND tag_name = ?", userId, tagName)
	return result.Error
}
