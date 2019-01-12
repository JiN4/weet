package model

// userQuestionAndAnswerTableにInsertする
func CreateUserQuestionAndAnswer(userQuestionAndAnswer UserQuestionAndAnswer) (UserQuestionAndAnswer, error) {
	err := db.Create(&userQuestionAndAnswer).Error
	if err != nil {
		return UserQuestionAndAnswer{}, err
	}
	return userQuestionAndAnswer, nil
}
