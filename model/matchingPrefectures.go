package model

import (
	"strconv"
	"strings"
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
