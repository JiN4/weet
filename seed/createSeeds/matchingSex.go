package createSeeds

import (
	"fmt"
	"strconv"

	"github.com/weet/model"
)

func CreateSeedMatchingSexes() {

	matchingSexes_infos := []map[string]string{
		map[string]string{
			"UserID": "1",
			"SexID":  "2",
		},
		map[string]string{
			"UserID": "2",
			"SexID":  "2",
		},
		map[string]string{
			"UserID": "3",
			"SexID":  "3",
		},
		map[string]string{
			"UserID": "4",
			"SexID":  "2",
		},
		map[string]string{
			"UserID": "5",
			"SexID":  "1",
		},
		map[string]string{
			"UserID": "7",
			"SexID":  "3",
		},
		map[string]string{
			"UserID": "8",
			"SexID":  "1",
		},
		map[string]string{
			"UserID": "9",
			"SexID":  "3",
		},
	}

	for _, info := range matchingSexes_infos {
		userID, _ := strconv.Atoi(info["UserID"])
		sexID, _ := strconv.Atoi(info["SexID"])
		createMatchingSex(model.MatchingSex{
			UserID: uint(userID),
			SexID:  uint(sexID),
		})
	}
}

func createMatchingSex(matchingSex model.MatchingSex) {
	matchingSex, err := model.CreateMatchingSex(matchingSex)
	if err != nil {
		panic(err)
	}
	fmt.Printf("matchingSex created\n")
}
