package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

// TableにInsertする
func CreateMatchingPrefectures(matchingPrefectures MatchingPrefectures) (MatchingPrefectures, error) {
	err := db.Create(&matchingPrefectures).Error
	if err != nil {
		return MatchingPrefectures{}, err
	}
	return matchingPrefectures, nil
}

func GetMatchingPrefecturesByUserID(userId uint) ([]uint, error) {
	matchingPrefecturesAll := MatchingPrefectures{}
	var stringPrefecturesIds []string
	var prefecturesId int
	var prefecturesIds []uint
	err := db.Select("prefectures_id").Where("user_id = ?", userId).First(&matchingPrefecturesAll).Error

	stringPrefecturesIds = strings.Split(matchingPrefecturesAll.PrefecturesID, ",")

	for _, stringPrefecturesId := range stringPrefecturesIds {
		prefecturesId, _ = strconv.Atoi(stringPrefecturesId)
		prefecturesIds = append(prefecturesIds, uint(prefecturesId))
	}

	return prefecturesIds, err
}

func PutMatchingPrefectures(c *gin.Context, userId uint) error {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))

	var matchingPrefectures MatchingPrefectures
	err = json.Unmarshal(body, &matchingPrefectures)

	fmt.Println(matchingPrefectures)

	err = db.Model(&matchingPrefectures).Where("user_id = ?", userId).Update("prefectures_id", matchingPrefectures.PrefecturesID).Error
	return err
}
