package model

// sampleTableにInsertする
func CreateUserSex(userSex UserSex) (UserSex, error) {
	err := db.Create(&userSex).Error
	if err != nil {
		return UserSex{}, err
	}
	return userSex, nil
}

// func GetUserIdBySexId(sexId uint) (uint, error) {
// 	userSexes := []UserSex{}
// 	err := db.Select("user_id").Where("sex_id = ?", sexId).Find(&userSexes).Error

// 	fmt.Println(userSexes)
// 	rand.Seed(time.Now().UnixNano())
// 	userId := userSexes[rand.Intn(len(userSexes))].UserID

// 	return userId, err
// }
