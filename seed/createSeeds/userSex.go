package createSeeds

import (
	"fmt"
	"strconv"

	"github.com/weet/model"
)

func CreateSeedUserSexes() {

	userSex_infos := []map[string]string{
		map[string]string{
			"UserID": "1",
			"SexID":  "1",
		},
		map[string]string{
			"UserID": "2",
			"SexID":  "1",
		},
		map[string]string{
			"UserID": "3",
			"SexID":  "1",
		},
		map[string]string{
			"UserID": "4",
			"SexID":  "1",
		},
		map[string]string{
			"UserID": "5",
			"SexID":  "1",
		},
		// map[string]string{
		// 	"UserID": "6",
		// 	"SexID":  "2",
		// },
		map[string]string{
			"UserID": "7",
			"SexID":  "2",
		},
		map[string]string{
			"UserID": "8",
			"SexID":  "2",
		},
		map[string]string{
			"UserID": "9",
			"SexID":  "2",
		},
		map[string]string{
			"UserID": "10",
			"SexID":  "2",
		},
	}

	for _, info := range userSex_infos {
		userID, _ := strconv.Atoi(info["UserID"])
		sexID, _ := strconv.Atoi(info["SexID"])
		createUserSex(model.UserSex{
			UserID: uint(userID),
			SexID:  uint(sexID),
		})
	}
}

func createUserSex(userSex model.UserSex) {
	userSex, err := model.CreateUserSex(userSex)
	if err != nil {
		panic(err)
	}
	fmt.Printf("userSexcreated\n")
}
