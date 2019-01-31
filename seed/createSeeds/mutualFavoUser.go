package createSeeds

import (
	"fmt"
	"strconv"

	"github.com/weet/model"
)

func CreateSeedMutualFavoUsers() {

	mutualFavoUsers_infos := []map[string]string{
		map[string]string{
			"MatchingFormatID": "1",
			"UserID1":          "1",
			"UserID2":          "5",
		},
		map[string]string{
			"MatchingFormatID": "2",
			"UserID1":          "6",
			"UserID2":          "1",
		},
	}

	for _, info := range mutualFavoUsers_infos {
		matchingFormatID, _ := strconv.Atoi(info["MatchingFormatID"])
		userID1, _ := strconv.Atoi(info["UserID1"])
		userID2, _ := strconv.Atoi(info["UserID2"])
		createMutualFavoUser(model.MutualFavoUser{
			MatchingFormatID: uint(matchingFormatID),
			UserID1:          uint(userID1),
			UserID2:          uint(userID2),
		})
	}
}

func createMutualFavoUser(mutualFavoUser model.MutualFavoUser) {
	mutualFavoUser, err := model.CreateMutualFavoUser(mutualFavoUser)
	if err != nil {
		panic(err)
	}
	fmt.Printf("mutualFavoUser created\n")
}
