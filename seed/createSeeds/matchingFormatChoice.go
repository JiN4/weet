package createSeeds

import (
	"fmt"
	"strconv"

	"github.com/weet/model"
)

func CreateSeedMatchingFormatChoices() {

	matchingFormatChoices_infos := []map[string]string{
		map[string]string{
			"UserID":    "2",
			"Love":      "1",
			"Marriage":  "1",
			"Roommate ": "0",
		},
		map[string]string{
			"UserID":    "3",
			"Love":      "1",
			"Marriage":  "1",
			"Roommate ": "1",
		},
		map[string]string{
			"UserID":    "4",
			"Love":      "1",
			"Marriage":  "1",
			"Roommate ": "1",
		},
		map[string]string{
			"UserID":    "5",
			"Love":      "1",
			"Marriage":  "1",
			"Roommate ": "1",
		},
		map[string]string{
			"UserID":    "7",
			"Love":      "1",
			"Marriage":  "1",
			"Roommate ": "1",
		},
		map[string]string{
			"UserID":    "8",
			"Love":      "1",
			"Marriage":  "1",
			"Roommate ": "1",
		},
		map[string]string{
			"UserID":    "9",
			"Love":      "0",
			"Marriage":  "0",
			"Roommate ": "0",
		},
	}

	for _, info := range matchingFormatChoices_infos {
		userID, _ := strconv.Atoi(info["UserID"])
		love, _ := strconv.Atoi(info["Love"])
		marriage, _ := strconv.Atoi(info["Marriage"])
		roommate, _ := strconv.Atoi(info["Roommate"])
		createMatchingFormatChoice(model.MatchingFormatChoice{
			UserID:   uint(userID),
			Love:     uint(love),
			Marriage: uint(marriage),
			Roommate: uint(roommate),
		})
	}
}

func createMatchingFormatChoice(matchingFormatChoice model.MatchingFormatChoice) {
	matchingFormatChoice, err := model.CreateMatchingFormatChoice(matchingFormatChoice)
	if err != nil {
		panic(err)
	}
	fmt.Printf("matchingFormatChoice created\n")
}
