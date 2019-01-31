package model

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

// sampleTableにInsertする
func CreateMatchingFormatChoice(matchingFormatChoice MatchingFormatChoice) (MatchingFormatChoice, error) {
	err := db.Create(&matchingFormatChoice).Error
	if err != nil {
		return MatchingFormatChoice{}, err
	}
	return matchingFormatChoice, nil
}

func PutMatchingFormatChoices(c *gin.Context, userId uint) error {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))

	var matchingFormatChoice MatchingFormatChoice
	err = json.Unmarshal(body, &matchingFormatChoice)

	err = db.Raw("UPDATE matching_format_choices SET love = ?, marriage = ?, roommate = ? WHERE user_id = ?", matchingFormatChoice.Love, matchingFormatChoice.Marriage, matchingFormatChoice.Roommate, userId).Scan(&matchingFormatChoice).Error

	return err
}
