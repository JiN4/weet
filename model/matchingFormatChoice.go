package model

// sampleTableにInsertする
func CreateMatchingFormatChoice(matchingFormatChoice MatchingFormatChoice) (MatchingFormatChoice, error) {
	err := db.Create(&matchingFormatChoice).Error
	if err != nil {
		return MatchingFormatChoice{}, err
	}
	return matchingFormatChoice, nil
}
