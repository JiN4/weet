package model

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
)

// TableにInsertする
func CreateMatchingSex(matchingSex MatchingSex) (MatchingSex, error) {
	err := db.Create(&matchingSex).Error
	if err != nil {
		return MatchingSex{}, err
	}
	return matchingSex, nil
}

func GetMatchingSexByUserID(userId uint) (uint, error) {
	matchingSex := MatchingSex{}
	err := db.Select("sex_id").Where("user_id = ?", userId).First(&matchingSex).Error

	sexId := matchingSex.SexID

	return sexId, err
}

func PostUserSexes(c *gin.Context, userId uint) error {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))

	var userSex UserSex
	err = json.Unmarshal(body, &userSex)

	err = db.Model(&userSex).Where("user_id = ?", userId).Update("SexID", userSex.SexID).Error
	return err
}