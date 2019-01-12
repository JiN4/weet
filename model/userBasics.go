package model

// userBasicsTableにInsertする
func CreateUserBasics(userBasics UserBasics) (UserBasics, error) {
	err := db.Create(&userBasics).Error
	if err != nil {
		return UserBasics{}, err
	}
	return userBasics, nil
}
