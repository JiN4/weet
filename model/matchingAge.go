package model

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
)

// TableにInsertする
func CreateMatchingAge(matchingAge MatchingAge) (MatchingAge, error) {
	err := db.Create(&matchingAge).Error
	if err != nil {
		return MatchingAge{}, err
	}
	return matchingAge, nil
}

func GetMatchingAgeByUserID(userId uint) (uint, uint, error) {
	matchingAge := MatchingAge{}

	err := db.Select("first_age, last_age").Where("user_id = ?", userId).First(&matchingAge).Error

	firstAge := matchingAge.FirstAge
	lastAge := matchingAge.LastAge

	return firstAge, lastAge, err
}

func PostMatcingAges(c *gin.Context, userId uint) error {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))

	var matchingAge MatchingAge
	err = json.Unmarshal(body, &matchingAge)

	if 20 <= matchingAge.FirstAge && matchingAge.FirstAge <= matchingAge.LastAge {
		err = db.Model(&matchingAge).Where("user_id = ?", userId).Updates(MatchingAge{
			FirstAge: matchingAge.FirstAge,
			LastAge:  matchingAge.LastAge,
		}).Error
	}

	return err
}