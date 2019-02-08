package model

import "github.com/weet/service"

// userQuestionAndAnswerTableにInsertする
func CreateUserQuestionAndAnswer(userQuestionAndAnswer UserQuestionAndAnswer) (UserQuestionAndAnswer, error) {
	err := db.Create(&userQuestionAndAnswer).Error
	if err != nil {
		return UserQuestionAndAnswer{}, err
	}
	return userQuestionAndAnswer, nil
}

//user_idとFormat_Idで条件を絞り、質問と答えが１セットの配列が欲しい(N+1問題の解決のためにクエリを手打ち)
func GetUserQuestionAndAnswerByUserIDAndFormatID(userId uint, matchingFormatId uint) ([]service.UserQuestionAndAnswer, error) {
	userQuestionAndAnswers := []service.UserQuestionAndAnswer{}

	err := db.Raw("select questions.id AS question_id, questions.name AS question_name , answers.name AS answer_name from user_question_and_answers join user_basics on (user_question_and_answers.user_id = user_basics.id) left join questions on (user_question_and_answers.question_id = questions.id) left join answers on (user_question_and_answers.answer_id = answers.id) where user_question_and_answers.user_id = ? and user_question_and_answers.matching_format_id = ?", userId, matchingFormatId).Scan(&userQuestionAndAnswers).Error
	return userQuestionAndAnswers, err
}
