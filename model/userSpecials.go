package model

import (
	"fmt"

	"github.com/weet/service"
)

func GetUserSpecialsById(userId uint) (service.UserSpecials, error) {
	userSpecial := service.UserSpecial{}
	userSpecials := service.UserSpecials{}
	var err error

	for i := 1; i < 5; i++ {

		//userSpecial.MatchingFormatName = GetMatchingFormatNameByID(uint(i))

		switch i {
		case 1:
			userSpecial.MatchingFormatName = "友達"
		case 2:
			userSpecial.MatchingFormatName = "恋愛"
		case 3:
			userSpecial.MatchingFormatName = "婚活"
		case 4:
			userSpecial.MatchingFormatName = "ルームメイト"
		}
		userSpecial.UserQuestionsAndAnswers, err = GetUserQuestionAndAnswerByUserIDAndFormatID(userId, uint(i))
		userSpecials = append(userSpecials, userSpecial)
	}
	return userSpecials, err
}

func ComparisonUserSpecial(userId uint, candidateMatchingUserId uint, matchingFormatIds []uint) ([]uint, []uint, error) {
	userQuestionAndAnswers := []UserQuestionAndAnswer{}
	userIdealQuestionAndAnswers := []UserIdealQuestionAndAnswer{}
	cmUserQuestionAndAnswers := []UserQuestionAndAnswer{}
	cmUserIdealQuestionAndAnswers := []UserIdealQuestionAndAnswer{}
	userQAndAs := map[uint]uint{}
	userIdealQAndAs := map[uint][]uint{}
	cmUserQAndAs := map[uint]uint{}
	cmUserIdealQAndAs := map[uint][]uint{}
	var firstQuestionId uint
	var idealAnswers []uint
	var matchingQuestions []uint
	var matchingIdealQuestions []uint
	var err error

	//マッチング候補者のマイページQandA取得
	err = db.Where("user_id = ? and matching_format_id in (?) and answer_id < 236", candidateMatchingUserId, matchingFormatIds).Find(&cmUserQuestionAndAnswers).Error

	for _, cmUserQuestionAndAnswer := range cmUserQuestionAndAnswers {
		cmUserQAndAs[cmUserQuestionAndAnswer.QuestionID] = cmUserQuestionAndAnswer.AnswerID
	}
	//fmt.Println(cmUserQAndAs)

	//プレイヤーの理想像QandA取得
	err = db.Where("user_id = ? and matching_format_id in (?) and answer_id < 236", userId, matchingFormatIds).Find(&userIdealQuestionAndAnswers).Error

	for i, userIdealQuestionAndAnswer := range userIdealQuestionAndAnswers {
		if i == 0 {
			firstQuestionId = userIdealQuestionAndAnswer.QuestionID
		} else if userIdealQuestionAndAnswers[i-1].QuestionID != userIdealQuestionAndAnswer.QuestionID {
			idealAnswers = nil
		}
		idealAnswers = append(idealAnswers, userIdealQuestionAndAnswer.AnswerID)
		userIdealQAndAs[userIdealQuestionAndAnswer.QuestionID] = idealAnswers
	}
	//fmt.Println(userIdealQAndAs)

	idealAnswers = nil

	for i := firstQuestionId; i < (firstQuestionId + uint(len(userIdealQAndAs))); i++ {
		idealAnswers = userIdealQAndAs[uint(i)] //マッチング相手の理想像の回答群
		for _, idealAnswer := range idealAnswers {
			if idealAnswer == cmUserQAndAs[uint(i)] { //プレイヤーの答えとマッチング相手の理想像の答えを一つずつ照合
				matchingQuestions = append(matchingQuestions, uint(i))
			}
		}
	}

	fmt.Println(matchingQuestions)
	idealAnswers = nil

	//プレイヤーのマイページQandA取得
	err = db.Where("user_id = ? and matching_format_id in (?) and answer_id < 236", userId, matchingFormatIds).Find(&userQuestionAndAnswers).Error

	for _, userQuestionAndAnswer := range userQuestionAndAnswers {
		userQAndAs[userQuestionAndAnswer.QuestionID] = userQuestionAndAnswer.AnswerID
	}
	//fmt.Println(userQAndAs)

	//マッチング候補者の理想像QandA取得
	err = db.Where("user_id = ? and matching_format_id in (?) and answer_id < 236", candidateMatchingUserId, matchingFormatIds).Find(&cmUserIdealQuestionAndAnswers).Error

	for i, cmUserIdealQuestionAndAnswer := range cmUserIdealQuestionAndAnswers {
		if i == 0 {
			firstQuestionId = cmUserIdealQuestionAndAnswer.QuestionID
		} else if cmUserIdealQuestionAndAnswers[i-1].QuestionID != cmUserIdealQuestionAndAnswer.QuestionID {
			idealAnswers = nil
		}
		idealAnswers = append(idealAnswers, cmUserIdealQuestionAndAnswer.AnswerID)
		cmUserIdealQAndAs[cmUserIdealQuestionAndAnswer.QuestionID] = idealAnswers
	}
	//fmt.Println(cmUserIdealQAndAs)

	idealAnswers = nil

	for i := firstQuestionId; i < (firstQuestionId + uint(len(cmUserIdealQAndAs))); i++ {
		idealAnswers = cmUserIdealQAndAs[uint(i)] //マッチング相手の理想像の回答群
		for _, idealAnswer := range idealAnswers {
			if idealAnswer == userQAndAs[uint(i)] { //プレイヤーの答えとマッチング相手の理想像の答えを一つずつ照合
				matchingIdealQuestions = append(matchingIdealQuestions, uint(i))
			}
		}
	}

	fmt.Println(matchingIdealQuestions)

	return matchingQuestions, matchingIdealQuestions, err
}

/*
//マッチング形式ごとのタイトルと質疑応答
type UserSpecial struct {
	[
		MatchingFormatName
		[
			QuestionID   uint   `json:"question_id"`
			QuestionName string `json:"question_name"`
			AnswerName   string `json:"answer_name"`
		]
	]
}
*/
