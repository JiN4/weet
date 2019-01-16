package model

import (
	"github.com/weet/service"
)

func GetUserIdealSpecialsById(userId uint) service.UserIdealSpecials {
	userIdealSpecial := service.UserIdealSpecial{}
	userIdealSpecials := service.UserIdealSpecials{}

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
		userIdealSpecial.UserIdealQuestionsAndAnswers = GetUserIdealQuestionAndAnswerByUserIDAndFormatID(userId, uint(i))
		userIdealSpecials = append(userIdealSpecials, userIdealSpecial)
	}
	return userIdealSpecials
}
