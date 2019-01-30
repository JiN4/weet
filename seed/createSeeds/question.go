package createSeeds

import (
	"fmt"
	"strconv"

	"github.com/weet/model"
)

func CreateSeedQuestions() {

	question_infos := []map[string]string{
		map[string]string{
			"MatchingFormatID": "1",
			"Name":             "出身地",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"Name":             "血液型",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"Name":             "学歴",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"Name":             "職種",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"Name":             "休日の曜日",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"Name":             "休日の過ごし方",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"Name":             "身長",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"Name":             "体型",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"Name":             "お酒",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"Name":             "煙草",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"Name":             "ギャンブル",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"Name":             "免許",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"Name":             "一ヶ月の娯楽費",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"Name":             "会話",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"Name":             "出会うまでの希望",
		},
		map[string]string{ //16
			"MatchingFormatID": "2",
			"Name":             "交際経験",
		},
		map[string]string{
			"MatchingFormatID": "2",
			"Name":             "メッセージ交換の頻度",
		},
		map[string]string{
			"MatchingFormatID": "2",
			"Name":             "デートの頻度",
		},
		map[string]string{
			"MatchingFormatID": "2",
			"Name":             "初回デートの費用",
		},
		map[string]string{ //20
			"MatchingFormatID": "2",
			"Name":             "同棲の希望",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"Name":             "結婚経験",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"Name":             "家事",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"Name":             "育児",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"Name":             "年収",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"Name":             "結婚式",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"Name":             "子供の有無",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"Name":             "子供はいつ欲しい？",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"Name":             "子供は何人欲しい？",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"Name":             "親との同居",
		},
		map[string]string{ //30
			"MatchingFormatID": "4",
			"Name":             "一番の目的",
		},
		map[string]string{
			"MatchingFormatID": "4",
			"Name":             "期間",
		},
		map[string]string{
			"MatchingFormatID": "4",
			"Name":             "部屋の貸し借り",
		},
		map[string]string{
			"MatchingFormatID": "4",
			"Name":             "ペット",
		},
		map[string]string{
			"MatchingFormatID": "4",
			"Name":             "来客",
		},
	}

	for _, info := range question_infos {
		matchingFormatID, _ := strconv.Atoi(info["MatchingFormatID"])
		createQuestion(model.Question{
			MatchingFormatID: uint(matchingFormatID),
			Name:             info["Name"],
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
