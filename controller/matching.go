package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/weet/model"
	"github.com/weet/service"
)

func GetMatchingUser(c *gin.Context) {
	matchingUser := service.MatchingUser{}
	userSpecial := service.UserSpecial{}
	userSpecials := service.UserSpecials{}
	userIdealSpecial := service.UserIdealSpecial{}
	userIdealSpecials := service.UserIdealSpecials{}

	var userId uint
	var matchingFormatId uint
	matchingFormatIds := []uint{1}
	var matchingPrefectures []uint
	var matchingSex uint
	var matchingFirstAge uint
	var matchingLastAge uint
	var candidateMatchingUserId uint
	var matchingQuestions []uint
	var matchingIdealQuestions []uint

	var err error

	userId, err = GetUint(c, "user_id")
	matchingFormatId, err = GetUint(c, "matching_format_id")

	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	matchingPrefectures, err = model.GetMatchingPrefecturesByUserID(userId)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	matchingSex, err = model.GetMatchingSexByUserID(userId)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	matchingFirstAge, matchingLastAge, err = model.GetMatchingAgeByUserID(userId)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	candidateMatchingUserId, err = model.GetCandidateMatchingUserId(userId, matchingPrefectures, matchingSex, matchingFirstAge, matchingLastAge, matchingFormatId)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	userSpecial.MatchingFormatName = "友達"
	userIdealSpecial.MatchingFormatName = "友達"
	userSpecial.UserQuestionsAndAnswers, err = model.GetUserQuestionAndAnswerByUserIDAndFormatID(candidateMatchingUserId, uint(1))
	userSpecials = append(userSpecials, userSpecial)
	userIdealSpecial.UserIdealQuestionsAndAnswers, err = model.GetUserIdealQuestionAndAnswerByUserIDAndFormatID(candidateMatchingUserId, uint(1))
	userIdealSpecials = append(userIdealSpecials, userIdealSpecial)

	matchingUser.UserSpecials = userSpecials
	matchingUser.UserIdealSpecials = userIdealSpecials

	if matchingFormatId == 2 {
		matchingFormatIds = append(matchingFormatIds, 2)

		userSpecial.MatchingFormatName = "恋愛"
		userIdealSpecial.MatchingFormatName = "恋愛"

		userSpecial.UserQuestionsAndAnswers, err = model.GetUserQuestionAndAnswerByUserIDAndFormatID(candidateMatchingUserId, uint(2))
		userSpecials = append(userSpecials, userSpecial)

		userIdealSpecial.UserIdealQuestionsAndAnswers, err = model.GetUserIdealQuestionAndAnswerByUserIDAndFormatID(candidateMatchingUserId, uint(2))
		userIdealSpecials = append(userIdealSpecials, userIdealSpecial)

		matchingUser.UserSpecials = userSpecials
		matchingUser.UserIdealSpecials = userIdealSpecials
	} else if matchingFormatId == 3 {
		matchingFormatIds = append(matchingFormatIds, 3)

		userSpecial.MatchingFormatName = "婚活"
		userIdealSpecial.MatchingFormatName = "婚活"

		userSpecial.UserQuestionsAndAnswers, err = model.GetUserQuestionAndAnswerByUserIDAndFormatID(candidateMatchingUserId, uint(3))
		userSpecials = append(userSpecials, userSpecial)

		userIdealSpecial.UserIdealQuestionsAndAnswers, err = model.GetUserIdealQuestionAndAnswerByUserIDAndFormatID(candidateMatchingUserId, uint(3))
		userIdealSpecials = append(userIdealSpecials, userIdealSpecial)

		matchingUser.UserSpecials = userSpecials
		matchingUser.UserIdealSpecials = userIdealSpecials
	} else if matchingFormatId == 4 {
		matchingFormatIds = append(matchingFormatIds, 4)

		userSpecial.MatchingFormatName = "ルームメイト"
		userIdealSpecial.MatchingFormatName = "ルームメイト"

		userSpecial.UserQuestionsAndAnswers, err = model.GetUserQuestionAndAnswerByUserIDAndFormatID(candidateMatchingUserId, uint(4))
		userSpecials = append(userSpecials, userSpecial)

		userIdealSpecial.UserIdealQuestionsAndAnswers, err = model.GetUserIdealQuestionAndAnswerByUserIDAndFormatID(candidateMatchingUserId, uint(4))
		userIdealSpecials = append(userIdealSpecials, userIdealSpecial)

		matchingUser.UserSpecials = userSpecials
		matchingUser.UserIdealSpecials = userIdealSpecials
	}

	matchingQuestions, matchingIdealQuestions, err = model.ComparisonUserSpecial(userId, candidateMatchingUserId, matchingFormatIds)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	matchingUser.UserBasics, err = model.GetUserBasicsById(candidateMatchingUserId)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	matchingUser.MatchingQuestions = matchingQuestions

	matchingUser.MatchingIdealQuestions = matchingIdealQuestions

	c.JSON(http.StatusOK, matchingUser)
}

func PostUserSexes(c *gin.Context){
	userId, err := GetUint(c, "user_id")
	err = model.PostUserSexes(c, userId)
	if err != nil {
		log.Println(err)
	}
}

func PostMatchingFormatChoices(c *gin.Context){
	userId, err := GetUint(c, "user_id")
	err = model.PostMatchingFormatChoices(c, userId)
	if err != nil {
		log.Println(err)
	}
}

func PostMatcingAges(c *gin.Context){
	userId, err := GetUint(c, "user_id")
	err = model.PostMatcingAges(c, userId)
	if err != nil {
		log.Println(err)
	}
}