package main

import (
	_ "github.com/jinzhu/gorm"
	"github.com/weet/model"
)

func main() {

	db := model.GetDBConn()

	// migrate

	db.DropTableIfExists(&model.Sample{})
	db.DropTableIfExists(&model.UserBasics{})
	db.DropTableIfExists(&model.UserQuestionAndAnswer{})
	db.DropTableIfExists(&model.Question{})
	db.DropTableIfExists(&model.Answer{})
	db.DropTableIfExists(&model.MatchingFormat{})
	db.DropTableIfExists(&model.UserIdealQuestionAndAnswer{})
	db.DropTableIfExists(&model.UserSex{})
	db.DropTableIfExists(&model.Prefectures{})
	db.DropTableIfExists(&model.MatchingPrefectures{})
	db.DropTableIfExists(&model.MatchingSex{})
	db.DropTableIfExists(&model.MatchingAge{})
	db.DropTableIfExists(&model.MatchingFormatChoice{})
	db.DropTableIfExists(&model.FavoUser{})
	db.DropTableIfExists(&model.MutualFavoUser{})
	db.DropTableIfExists(&model.MessageUser{})

	db.CreateTable(&model.Sample{})
	db.CreateTable(&model.UserBasics{})
	db.CreateTable(&model.UserQuestionAndAnswer{})
	db.CreateTable(&model.Question{})
	db.CreateTable(&model.Answer{})
	db.CreateTable(&model.MatchingFormat{})
	db.CreateTable(&model.UserIdealQuestionAndAnswer{})
	db.CreateTable(&model.UserSex{})
	db.CreateTable(&model.Prefectures{})
	db.CreateTable(&model.MatchingPrefectures{})
	db.CreateTable(&model.MatchingSex{})
	db.CreateTable(&model.MatchingAge{})
	db.CreateTable(&model.MatchingFormatChoice{})
	db.CreateTable(&model.FavoUser{})
	db.CreateTable(&model.MutualFavoUser{})
	db.CreateTable(&model.MessageUser{})
}
