package model

// formatTableにInsertする
func CreateMatchingFormat(matchingFormat MatchingFormat) (MatchingFormat, error) {
	err := db.Create(&matchingFormat).Error
	if err != nil {
		return MatchingFormat{}, err
	}
	return matchingFormat, nil
}
