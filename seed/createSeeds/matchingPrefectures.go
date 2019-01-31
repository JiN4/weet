package createSeeds

import (
	"fmt"
	"strconv"

	"github.com/weet/model"
)

func CreateSeedMatchingPrefectures() {

	matchingPrefectures_infos := []map[string]string{
		map[string]string{
			"UserID":        "1",
			"PrefecturesID": "26,27,28",
		},
		map[string]string{
			"UserID":        "2",
			"PrefecturesID": "26,27,28",
		},
		map[string]string{
			"UserID":        "3",
			"PrefecturesID": "26,27,28",
		},
		map[string]string{
			"UserID":        "4",
			"PrefecturesID": "26,27,28",
		},
		map[string]string{
			"UserID":        "5",
			"PrefecturesID": "26,27,28",
		},
		map[string]string{
			"UserID":        "6",
			"PrefecturesID": "26,27,28",
		},
		map[string]string{
			"UserID":        "7",
			"PrefecturesID": "26,27,28",
		},
		map[string]string{
			"UserID":        "8",
			"PrefecturesID": "26,27,28",
		},
		map[string]string{
			"UserID":        "9",
			"PrefecturesID": "26,27,28",
		},
	}

	for _, info := range matchingPrefectures_infos {
		userID, _ := strconv.Atoi(info["UserID"])
		createMatchingPrefectures(model.MatchingPrefectures{
			UserID:        uint(userID),
			PrefecturesID: info["PrefecturesID"],
		})
	}
}

func createMatchingPrefectures(matchingPrefectures model.MatchingPrefectures) {
	matchingPrefectures, err := model.CreateMatchingPrefectures(matchingPrefectures)
	if err != nil {
		panic(err)
	}
	fmt.Printf("matchingPrefectures created\n")
}
