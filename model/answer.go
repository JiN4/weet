package model

// answerTableにInsertする
func CreateAnswer(answer Answer) (Answer, error) {
	err := db.Create(&answer).Error
	if err != nil {
		return Answer{}, err
	}
	return answer, nil
}
