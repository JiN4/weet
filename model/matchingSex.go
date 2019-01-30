package model

// TableにInsertする
func CreateMatchingSex(matchingSex MatchingSex) (MatchingSex, error) {
	err := db.Create(&matchingSex).Error
	if err != nil {
		return MatchingSex{}, err
	}
	return matchingSex, nil
}

func GetMatchingSexByUserID(userId uint) (uint, error) {
	matchingSex := MatchingSex{}
	err := db.Select("sex_id").Where("user_id = ?", userId).First(&matchingSex).Error

	sexId := matchingSex.SexID

	return sexId, err
}
