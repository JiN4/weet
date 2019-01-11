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
	db.DropTableIfExists(&model.Format{})
	db.DropTableIfExists(&model.Question2{})

	db.CreateTable(&model.Sample{})
	db.CreateTable(&model.Question{})
	db.CreateTable(&model.Answer{})
	db.CreateTable(&model.Format{})
	db.CreateTable(&model.UserBasics{})
	db.CreateTable(&model.UserQuestionAndAnswer{})
	db.CreateTable(&model.Question2{})
	db.CreateTable(&model.Test1{})
	db.CreateTable(&model.Test2{})

}
