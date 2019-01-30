package createSeeds

import (
	"fmt"
	"strconv"

	"github.com/weet/model"
)

func CreateSeedMatchingAges() {

	matchingAges_infos := []map[string]string{
		map[string]string{
			"UserID":   "2",
			"FirstAge": "23",
			"LastAge":  "30",
		},
		map[string]string{
			"UserID":   "3",
			"FirstAge": "20",
			"LastAge":  "30",
		},
		map[string]string{
			"UserID":   "4",
			"FirstAge": "20",
			"LastAge":  "35",
		},
		map[string]string{
			"UserID":   "5",
			"FirstAge": "20",
			"LastAge":  "35",
		},
		map[string]string{
			"UserID":   "7",
			"FirstAge": "21",
			"LastAge":  "18",
		},
		map[string]string{
			"UserID":   "8",
			"FirstAge": "22",
			"LastAge":  "26",
		},
		map[string]string{
			"UserID":   "9",
			"FirstAge": "21",
			"LastAge":  "30",
		},
	}

	for _, info := range matchingAges_infos {
		userID, _ := strconv.Atoi(info["UserID"])
		firstAge, _ := strconv.Atoi(info["FirstAge"])
		lastAge, _ := strconv.Atoi(info["LastAge"])
		createMatchingAgo(model.MatchingAge{
			UserID:   uint(userID),
			FirstAge: uint(firstAge),
			LastAge:  uint(lastAge),
		})
	}
}

func createMatchingAgo(matchingAge model.MatchingAge) {
	matchingAge, err := model.CreateMatchingAge(matchingAge)
	if err != nil {
		panic(err)
	}
	fmt.Printf("matchingAgo created\n")
}
