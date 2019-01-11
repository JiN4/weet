package createSeeds

import (
	"fmt"
	"strconv"

	"github.com/weet/model"
)

func CreateSeedQuestion() {

	question_infos := []map[string]string{
		map[string]string{
			"FormatID": "1",
			"Name":     "血液型",
		},
		map[string]string{
			"FormatID": "1",
			"Name":     "インドア派？アウトドア派？",
		},
		map[string]string{
			"FormatID": "2",
			"Name":     "デートの頻度",
		},
		map[string]string{
			"FormatID": "3",
			"Name":     "子供は欲しい？",
		},
		map[string]string{
			"FormatID": "4",
			"Name":     "ルームシェアの目的は？",
		},
	}

	for _, info := range question_infos {
		formatID, _ := strconv.Atoi(info["FormatID"])
		createQuestion(model.Question{
			FormatID: uint(formatID),
			Name:     info["Name"],
		})
	}
}

func createQuestion(question model.Question) {
	question, err := model.CreateQuestion(question)
	if err != nil {
		panic(err)
	}
	fmt.Printf("question created\n")
}
