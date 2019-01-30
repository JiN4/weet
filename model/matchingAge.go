package model

// TableにInsertする
func CreateMatchingAge(matchingAge MatchingAge) (MatchingAge, error) {
	err := db.Create(&matchingAge).Error
	if err != nil {
		return MatchingAge{}, err
	}
	return matchingAge, nil
}

func GetMatchingAgeByUserID(userId uint) (uint, uint, error) {
	matchingAge := MatchingAge{}

	err := db.Select("first_age, last_age").Where("user_id = ?", userId).First(&matchingAge).Error

	firstAge := matchingAge.FirstAge
	lastAge := matchingAge.LastAge

	return firstAge, lastAge, err
}
