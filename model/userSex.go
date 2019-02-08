package model

// sampleTableにInsertする
func CreateUserSex(userSex UserSex) (UserSex, error) {
	err := db.Create(&userSex).Error
	if err != nil {
		return UserSex{}, err
	}
	return userSex, nil
}
