package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
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

func PutMatchingSexes(c *gin.Context, userId uint) error {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))

	var matchingSex MatchingSex
	err = json.Unmarshal(body, &matchingSex)

	fmt.Println(matchingSex)

	i := strconv.Itoa(int(matchingSex.SexID))

	err = db.Model(&matchingSex).Where("user_id = ?", userId).Update("sex_id", i).Error
	return err
}
