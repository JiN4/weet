package model

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/weet/service"
)

func GetUserIdealSpecialsById(userId uint) (service.UserIdealSpecials, error) {
	userIdealSpecial := service.UserIdealSpecial{}
	userIdealSpecials := service.UserIdealSpecials{}
	var err error

	for i := 1; i < 5; i++ {

		//userSpecial.MatchingFormatName = GetMatchingFormatNameByID(uint(i))

		switch i {
		case 1:
			userIdealSpecial.MatchingFormatName = "友達"
		case 2:
			userIdealSpecial.MatchingFormatName = "恋愛"
		case 3:
			userIdealSpecial.MatchingFormatName = "婚活"
		case 4:
			userIdealSpecial.MatchingFormatName = "ルームメイト"
		}
		userIdealSpecial.UserIdealQuestionsAndAnswers, err = GetUserIdealQuestionAndAnswerByUserIDAndFormatID(userId, uint(i))
		userIdealSpecials = append(userIdealSpecials, userIdealSpecial)
	}
	return userIdealSpecials, err
}

func GetUserIdealSpecialsById2(userId uint) (service.UserIdealSpecials2, error) {
	userIdealSpecial := service.UserIdealSpecial2{}
	userIdealSpecials := service.UserIdealSpecials2{}
	var err error

	for i := 1; i < 5; i++ {

		//userSpecial.MatchingFormatName = GetMatchingFormatNameByID(uint(i))

		switch i {
		case 1:
			userIdealSpecial.MatchingFormatName = "友達"
		case 2:
			userIdealSpecial.MatchingFormatName = "恋愛"
		case 3:
			userIdealSpecial.MatchingFormatName = "婚活"
		case 4:
			userIdealSpecial.MatchingFormatName = "ルームメイト"
		}
		userIdealSpecial.UserIdealQuestionsAndAnswers2, err = GetUserIdealQuestionAndAnswerByUserIDAndFormatID2(userId, uint(i))
		userIdealSpecials = append(userIdealSpecials, userIdealSpecial)
	}
	return userIdealSpecials, err
}

func UpdateUserIdealSpecials(c *gin.Context, userId uint) {
	userIdealQuestionAndAnswers := []UserIdealQuestionAndAnswer{}
	matchingFormatId := c.PostForm("matching_format_id")
	questionId := c.PostForm("question_id")
	answerIds := c.PostForm("answer_ids")
	var addAnswerIDs, deleteAnswerIDs []int
	var intAnswerID int
	var answerFlag bool

	fmt.Println("answerIds:", answerIds)

	answerIDs := strings.Split(answerIds, ",")

	intMatchingFormatID, _ := strconv.Atoi(matchingFormatId)
	intQuestionID, _ := strconv.Atoi(questionId)

	//DBに登録されているanswer_idを取得
	err := db.Select("answer_id").Where("user_id = ? and question_id = ?", userId, uint(intQuestionID)).Find(&userIdealQuestionAndAnswers).Error
	if err != nil {
		panic(err)
	}

	//27 28, 1 2
	//削除対象のanswer_idを選定
	for _, userIdealQuestionAndAnswer := range userIdealQuestionAndAnswers {
		for _, answerID := range answerIDs {
			intAnswerID, _ = strconv.Atoi(answerID)
			if userIdealQuestionAndAnswer.AnswerID == uint(intAnswerID) {
				answerFlag = true
				break
			}
		}
		if answerFlag == false {
			deleteAnswerIDs = append(deleteAnswerIDs, int(userIdealQuestionAndAnswer.AnswerID))
		}
		answerFlag = false
	}

	//1 2, 27 28
	//追加対象のanswer_idを選定
	for _, answerID := range answerIDs {
		for _, userIdealQuestionAndAnswer := range userIdealQuestionAndAnswers {
			intAnswerID, _ = strconv.Atoi(answerID)
			if uint(intAnswerID) == userIdealQuestionAndAnswer.AnswerID {
				answerFlag = true
				break
			}
		}
		if answerFlag == false {
			addAnswerIDs = append(addAnswerIDs, intAnswerID)
		}
		answerFlag = false
	}

	fmt.Println("delete:", deleteAnswerIDs)
	fmt.Println("add:", addAnswerIDs)
	//削除対象のanswer_idを削除
	for _, deleteAnswerID := range deleteAnswerIDs {
		err := db.Where("user_id = ? and answer_id = ?", userId, uint(deleteAnswerID)).Delete([]UserIdealQuestionAndAnswer{}).Error
		if err != nil {
			panic(err)
		}
	}

	//追加対象のanswer_idを削除
	for _, addAnswerID := range addAnswerIDs {
		CreateUserIdealQuestionAndAnswer(UserIdealQuestionAndAnswer{
			MatchingFormatID: uint(intMatchingFormatID),
			UserID:           userId,
			QuestionID:       uint(intQuestionID),
			AnswerID:         uint(addAnswerID),
		})
	}

}
