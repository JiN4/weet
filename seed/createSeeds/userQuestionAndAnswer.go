package createSeeds

import (
	"fmt"
	"strconv"

	"github.com/weet/model"
)

func CreateSeedUserQuestionAndAnswers() {

	questionAndAnswer_infos := []map[string]string{
		map[string]string{
			"UserID":     "1",
			"QuestionID": "1",
			"AnswerID":   "1",
		},
		map[string]string{
			"UserID":     "1",
			"QuestionID": "2",
			"AnswerID":   "5",
		},
		map[string]string{
			"UserID":     "1",
			"QuestionID": "3",
			"AnswerID":   "3",
		},
		map[string]string{
			"UserID":     "1",
			"QuestionID": "4",
			"AnswerID":   "13",
		},
		map[string]string{
			"UserID":     "1",
			"QuestionID": "5",
			"AnswerID":   "17",
		},
	}

	for _, info := range questionAndAnswer_infos {
		userID, _ := strconv.Atoi(info["UserID"])
		questionID, _ := strconv.Atoi(info["QuestionID"])
		answerID, _ := strconv.Atoi(info["AnswerID"])
		createUserQuestionAndAnswer(model.UserQuestionAndAnswer{
			UserID:     uint(userID),
			QuestionID: uint(questionID),
			AnswerID:   uint(answerID),
		})
	}
}

func createUserQuestionAndAnswer(userQuestionAndAnswer model.UserQuestionAndAnswer) {
	userQuestionAndAnswer, err := model.CreateUserQuestionAndAnswer(userQuestionAndAnswer)
	if err != nil {
		panic(err)
	}
	fmt.Printf("userQuestionAndAnswer created\n")
}
