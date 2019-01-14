package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/weet/model"
)

func GetUserById(c *gin.Context) {
	id, err := GetUint(c, "user_id")
	user, err := model.GetUserById(id)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, user)
}

/*
//マッチング形式ごとのタイトルと質疑応答
type UserSpecial struct {
	MatchingFormatName      string                  `json:"matching_format_name"`
	UserQuestionsAndAnswers UserQuestionsAndAnswers `json:"user_questions_and_answers"`
}

type UserSpecials []UserSpecial

//１セットの質疑応答
//select user_basics.user_name, questions.name, answers.name from user_question_and_answers join user_basics on (user_question_and_answers.user_id = user_basics.id) left join questions on (user_question_and_answers.question_id = questions.id) left join answers on (user_question_and_answers.answer_id = answers.id);
type UserQuestionAndAnswer struct {
	QuestionID   uint   `json:"question_id"`
	QuestionName string `json:"question_name"`
	AnswerName   string `json:"answer_name"`
}
*/
