package createSeeds

import (
	"fmt"
	"strconv"

	"github.com/weet/model"
)

func CreateSeedFavoUsers() {

	favoUsers_infos := []map[string]string{
		map[string]string{
			"MatchingFormatID": "1",
			"PlayerUserID":     "1",
			"FavoUserID":       "9",
		},
		map[string]string{
			"MatchingFormatID": "2",
			"PlayerUserID":     "1",
			"FavoUserID":       "7",
		},
	}

	for _, info := range favoUsers_infos {
		matchingFormatID, _ := strconv.Atoi(info["MatchingFormatID"])
		playerUserID, _ := strconv.Atoi(info["PlayerUserID"])
		favoUserID, _ := strconv.Atoi(info["FavoUserID"])
		createFavoUser(model.FavoUser{
			MatchingFormatID: uint(matchingFormatID),
			PlayerUserID:     uint(playerUserID),
			FavoUserID:       uint(favoUserID),
		})
	}
}

func createFavoUser(favoUser model.FavoUser) {
	favoUser, err := model.CreateFavoUser(favoUser)
	if err != nil {
		panic(err)
	}
	fmt.Printf("favoUser created\n")
}
