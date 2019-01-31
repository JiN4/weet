package createSeeds

import (
	"fmt"
	"strconv"

	"github.com/weet/model"
)

func CreateSeedUserQuestionAndAnswers() {
	questionAndAnswer_infos := []map[string]string{
		map[string]string{
			"MatchingFormatID": "1",  //　質問ドキュメント参照
			"UserID":           "1",  //　ユーザ情報ドキュメント参照
			"QuestionID":       "1",  //　質問ドキュメント参照
			"AnswerID":         "28", //　回答ドキュメント参照
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "1",
			"QuestionID":       "2",
			"AnswerID":         "48",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "1",
			"QuestionID":       "3",
			"AnswerID":         "53",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "1",
			"QuestionID":       "4",
			"AnswerID":         "61",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "1",
			"QuestionID":       "5",
			"AnswerID":         "68",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "1",
			"QuestionID":       "6",
			"AnswerID":         "71",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "1",
			"QuestionID":       "7",
			"AnswerID":         "115",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "1",
			"QuestionID":       "8",
			"AnswerID":         "124",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "1",
			"QuestionID":       "9",
			"AnswerID":         "131",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "1",
			"QuestionID":       "10",
			"AnswerID":         "136",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "1",
			"QuestionID":       "11",
			"AnswerID":         "139",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "1",
			"QuestionID":       "12",
			"AnswerID":         "143",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "1",
			"QuestionID":       "13",
			"AnswerID":         "146",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "1",
			"QuestionID":       "14",
			"AnswerID":         "152",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "1",
			"QuestionID":       "15",
			"AnswerID":         "155",
		},
		map[string]string{
			"MatchingFormatID": "2",
			"UserID":           "1",
			"QuestionID":       "16",
			"AnswerID":         "157",
		},
		map[string]string{
			"MatchingFormatID": "2",
			"UserID":           "1",
			"QuestionID":       "17",
			"AnswerID":         "160",
		},
		map[string]string{
			"MatchingFormatID": "2",
			"UserID":           "1",
			"QuestionID":       "18",
			"AnswerID":         "165",
		},
		map[string]string{
			"MatchingFormatID": "2",
			"UserID":           "1",
			"QuestionID":       "19",
			"AnswerID":         "169",
		},
		map[string]string{
			"MatchingFormatID": "2",
			"UserID":           "1",
			"QuestionID":       "20",
			"AnswerID":         "175",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "1",
			"QuestionID":       "21",
			"AnswerID":         "177",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "1",
			"QuestionID":       "22",
			"AnswerID":         "179",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "1",
			"QuestionID":       "23",
			"AnswerID":         "182",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "1",
			"QuestionID":       "24",
			"AnswerID":         "187",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "1",
			"QuestionID":       "25",
			"AnswerID":         "197",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "1",
			"QuestionID":       "26",
			"AnswerID":         "200",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "1",
			"QuestionID":       "27",
			"AnswerID":         "203",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "1",
			"QuestionID":       "28",
			"AnswerID":         "208",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "1",
			"QuestionID":       "29",
			"AnswerID":         "212",
		},
		map[string]string{
			"MatchingFormatID": "4",
			"UserID":           "1",
			"QuestionID":       "30",
			"AnswerID":         "214",
		},
		map[string]string{
			"MatchingFormatID": "4",
			"UserID":           "1",
			"QuestionID":       "31",
			"AnswerID":         "221",
		},
		map[string]string{
			"MatchingFormatID": "4",
			"UserID":           "1",
			"QuestionID":       "32",
			"AnswerID":         "226",
		},
		map[string]string{
			"MatchingFormatID": "4",
			"UserID":           "1",
			"QuestionID":       "33",
			"AnswerID":         "231",
		},
		map[string]string{
			"MatchingFormatID": "4",
			"UserID":           "1",
			"QuestionID":       "34",
			"AnswerID":         "233",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "2",
			"QuestionID":       "1",
			"AnswerID":         "26",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "2",
			"QuestionID":       "2",
			"AnswerID":         "51",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "2",
			"QuestionID":       "3",
			"AnswerID":         "54",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "2",
			"QuestionID":       "4",
			"AnswerID":         "58",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "2",
			"QuestionID":       "5",
			"AnswerID":         "70",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "2",
			"QuestionID":       "6",
			"AnswerID":         "71",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "2",
			"QuestionID":       "7",
			"AnswerID":         "105",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "2",
			"QuestionID":       "8",
			"AnswerID":         "124",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "2",
			"QuestionID":       "9",
			"AnswerID":         "130",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "2",
			"QuestionID":       "10",
			"AnswerID":         "134",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "2",
			"QuestionID":       "11",
			"AnswerID":         "137",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "2",
			"QuestionID":       "12",
			"AnswerID":         "142",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "2",
			"QuestionID":       "13",
			"AnswerID":         "146",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "2",
			"QuestionID":       "14",
			"AnswerID":         "150",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "2",
			"QuestionID":       "15",
			"AnswerID":         "155",
		},
		map[string]string{
			"MatchingFormatID": "2",
			"UserID":           "2",
			"QuestionID":       "16",
			"AnswerID":         "158",
		},
		map[string]string{
			"MatchingFormatID": "2",
			"UserID":           "2",
			"QuestionID":       "17",
			"AnswerID":         "160",
		},
		map[string]string{
			"MatchingFormatID": "2",
			"UserID":           "2",
			"QuestionID":       "18",
			"AnswerID":         "165",
		},
		map[string]string{
			"MatchingFormatID": "2",
			"UserID":           "2",
			"QuestionID":       "19",
			"AnswerID":         "170",
		},
		map[string]string{
			"MatchingFormatID": "2",
			"UserID":           "2",
			"QuestionID":       "20",
			"AnswerID":         "174",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "2",
			"QuestionID":       "21",
			"AnswerID":         "177",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "2",
			"QuestionID":       "22",
			"AnswerID":         "180",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "2",
			"QuestionID":       "23",
			"AnswerID":         "184",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "2",
			"QuestionID":       "24",
			"AnswerID":         "188",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "2",
			"QuestionID":       "25",
			"AnswerID":         "197",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "2",
			"QuestionID":       "26",
			"AnswerID":         "200",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "2",
			"QuestionID":       "27",
			"AnswerID":         "202",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "2",
			"QuestionID":       "28",
			"AnswerID":         "208",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "2",
			"QuestionID":       "29",
			"AnswerID":         "213",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "2",
			"QuestionID":       "29",
			"AnswerID":         "213",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "2",
			"QuestionID":       "29",
			"AnswerID":         "213",
		},
		map[string]string{
			"MatchingFormatID": "4",
			"UserID":           "2",
			"QuestionID":       "30",
			"AnswerID":         "265",
		},
		map[string]string{
			"MatchingFormatID": "4",
			"UserID":           "2",
			"QuestionID":       "31",
			"AnswerID":         "266",
		},
		map[string]string{
			"MatchingFormatID": "4",
			"UserID":           "2",
			"QuestionID":       "32",
			"AnswerID":         "267",
		},
		map[string]string{
			"MatchingFormatID": "4",
			"UserID":           "2",
			"QuestionID":       "33",
			"AnswerID":         "268",
		},
		map[string]string{
			"MatchingFormatID": "4",
			"UserID":           "2",
			"QuestionID":       "34",
			"AnswerID":         "269",
		},

		//UserId 3
		map[string]string{
			"MatchingFormatID": "1",  //　質問ドキュメント参照
			"UserID":           "3",  //　ユーザ情報ドキュメント参照
			"QuestionID":       "1",  //　質問ドキュメント参照
			"AnswerID":         "31", //　回答ドキュメント参照
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "3",
			"QuestionID":       "2",
			"AnswerID":         "48",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "3",
			"QuestionID":       "3",
			"AnswerID":         "54",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "3",
			"QuestionID":       "4",
			"AnswerID":         "61",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "3",
			"QuestionID":       "5",
			"AnswerID":         "68",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "3",
			"QuestionID":       "6",
			"AnswerID":         "71",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "3",
			"QuestionID":       "7",
			"AnswerID":         "108",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "3",
			"QuestionID":       "8",
			"AnswerID":         "126",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "3",
			"QuestionID":       "9",
			"AnswerID":         "131",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "3",
			"QuestionID":       "10",
			"AnswerID":         "136",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "3",
			"QuestionID":       "11",
			"AnswerID":         "139",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "3",
			"QuestionID":       "12",
			"AnswerID":         "140",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "3",
			"QuestionID":       "13",
			"AnswerID":         "145",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "3",
			"QuestionID":       "14",
			"AnswerID":         "152",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "3",
			"QuestionID":       "15",
			"AnswerID":         "155",
		},
		map[string]string{
			"MatchingFormatID": "2",
			"UserID":           "3",
			"QuestionID":       "16",
			"AnswerID":         "158",
		},
		map[string]string{
			"MatchingFormatID": "2",
			"UserID":           "3",
			"QuestionID":       "17",
			"AnswerID":         "160",
		},
		map[string]string{
			"MatchingFormatID": "2",
			"UserID":           "3",
			"QuestionID":       "18",
			"AnswerID":         "165",
		},
		map[string]string{
			"MatchingFormatID": "2",
			"UserID":           "3",
			"QuestionID":       "19",
			"AnswerID":         "171",
		},
		map[string]string{
			"MatchingFormatID": "2",
			"UserID":           "3",
			"QuestionID":       "20",
			"AnswerID":         "174",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "3",
			"QuestionID":       "21",
			"AnswerID":         "177",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "3",
			"QuestionID":       "22",
			"AnswerID":         "178",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "3",
			"QuestionID":       "23",
			"AnswerID":         "183",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "3",
			"QuestionID":       "24",
			"AnswerID":         "187",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "3",
			"QuestionID":       "25",
			"AnswerID":         "197",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "3",
			"QuestionID":       "26",
			"AnswerID":         "200",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "3",
			"QuestionID":       "27",
			"AnswerID":         "203",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "3",
			"QuestionID":       "28",
			"AnswerID":         "207",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "3",
			"QuestionID":       "29",
			"AnswerID":         "212",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "3",
			"QuestionID":       "30",
			"AnswerID":         "214",
		},
		map[string]string{
			"MatchingFormatID": "4",
			"UserID":           "3",
			"QuestionID":       "31",
			"AnswerID":         "218",
		},
		map[string]string{
			"MatchingFormatID": "4",
			"UserID":           "3",
			"QuestionID":       "32",
			"AnswerID":         "224",
		},
		map[string]string{
			"MatchingFormatID": "4",
			"UserID":           "3",
			"QuestionID":       "33",
			"AnswerID":         "229",
		},
		map[string]string{
			"MatchingFormatID": "4",
			"UserID":           "3",
			"QuestionID":       "34",
			"AnswerID":         "232",
		},

		//userId 4
		map[string]string{
			"MatchingFormatID": "1",  //　質問ドキュメント参照
			"UserID":           "4",  //　ユーザ情報ドキュメント参照
			"QuestionID":       "1",  //　質問ドキュメント参照
			"AnswerID":         "44", //　回答ドキュメント参照
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "4",
			"QuestionID":       "2",
			"AnswerID":         "49",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "4",
			"QuestionID":       "3",
			"AnswerID":         "52",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "4",
			"QuestionID":       "4",
			"AnswerID":         "67",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "4",
			"QuestionID":       "5",
			"AnswerID":         "70",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "4",
			"QuestionID":       "6",
			"AnswerID":         "72",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "4",
			"QuestionID":       "7",
			"AnswerID":         "111",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "4",
			"QuestionID":       "8",
			"AnswerID":         "129",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "4",
			"QuestionID":       "9",
			"AnswerID":         "130",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "4",
			"QuestionID":       "10",
			"AnswerID":         "136",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "4",
			"QuestionID":       "11",
			"AnswerID":         "139",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "4",
			"QuestionID":       "12",
			"AnswerID":         "141",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "4",
			"QuestionID":       "13",
			"AnswerID":         "146",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "4",
			"QuestionID":       "14",
			"AnswerID":         "152",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "4",
			"QuestionID":       "15",
			"AnswerID":         "155",
		},
		map[string]string{
			"MatchingFormatID": "2",
			"UserID":           "4",
			"QuestionID":       "16",
			"AnswerID":         "157",
		},
		map[string]string{
			"MatchingFormatID": "2",
			"UserID":           "4",
			"QuestionID":       "17",
			"AnswerID":         "159",
		},
		map[string]string{
			"MatchingFormatID": "2",
			"UserID":           "4",
			"QuestionID":       "18",
			"AnswerID":         "166",
		},
		map[string]string{
			"MatchingFormatID": "2",
			"UserID":           "4",
			"QuestionID":       "19",
			"AnswerID":         "168",
		},
		map[string]string{
			"MatchingFormatID": "2",
			"UserID":           "4",
			"QuestionID":       "20",
			"AnswerID":         "175",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "4",
			"QuestionID":       "21",
			"AnswerID":         "177",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "4",
			"QuestionID":       "22",
			"AnswerID":         "178",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "4",
			"QuestionID":       "23",
			"AnswerID":         "182",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "4",
			"QuestionID":       "24",
			"AnswerID":         "190",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "4",
			"QuestionID":       "25",
			"AnswerID":         "197",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "4",
			"QuestionID":       "26",
			"AnswerID":         "200",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "4",
			"QuestionID":       "27",
			"AnswerID":         "202",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "4",
			"QuestionID":       "28",
			"AnswerID":         "209",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "4",
			"QuestionID":       "29",
			"AnswerID":         "211",
		},
		map[string]string{
			"MatchingFormatID": "4",
			"UserID":           "4",
			"QuestionID":       "30",
			"AnswerID":         "218",
		},
		map[string]string{
			"MatchingFormatID": "4",
			"UserID":           "4",
			"QuestionID":       "31",
			"AnswerID":         "221",
		},
		map[string]string{
			"MatchingFormatID": "4",
			"UserID":           "4",
			"QuestionID":       "32",
			"AnswerID":         "226",
		},
		map[string]string{
			"MatchingFormatID": "4",
			"UserID":           "4",
			"QuestionID":       "33",
			"AnswerID":         "229",
		},
		map[string]string{
			"MatchingFormatID": "4",
			"UserID":           "4",
			"QuestionID":       "34",
			"AnswerID":         "232",
		},

		//userId 5
		map[string]string{
			"MatchingFormatID": "1",  //　質問ドキュメント参照
			"UserID":           "5",  //　ユーザ情報ドキュメント参照
			"QuestionID":       "1",  //　質問ドキュメント参照
			"AnswerID":         "11", //　回答ドキュメント参照
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "5",
			"QuestionID":       "2",
			"AnswerID":         "50",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "5",
			"QuestionID":       "3",
			"AnswerID":         "53",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "5",
			"QuestionID":       "4",
			"AnswerID":         "66",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "5",
			"QuestionID":       "5",
			"AnswerID":         "70",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "5",
			"QuestionID":       "6",
			"AnswerID":         "72",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "5",
			"QuestionID":       "7",
			"AnswerID":         "99",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "5",
			"QuestionID":       "8",
			"AnswerID":         "125",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "5",
			"QuestionID":       "9",
			"AnswerID":         "131",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "5",
			"QuestionID":       "10",
			"AnswerID":         "135",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "5",
			"QuestionID":       "11",
			"AnswerID":         "139",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "5",
			"QuestionID":       "12",
			"AnswerID":         "140",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "5",
			"QuestionID":       "13",
			"AnswerID":         "147",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "5",
			"QuestionID":       "14",
			"AnswerID":         "151",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "5",
			"QuestionID":       "15",
			"AnswerID":         "154",
		},
		map[string]string{
			"MatchingFormatID": "2",
			"UserID":           "5",
			"QuestionID":       "16",
			"AnswerID":         "158",
		},
		map[string]string{
			"MatchingFormatID": "2",
			"UserID":           "5",
			"QuestionID":       "17",
			"AnswerID":         "160",
		},
		map[string]string{
			"MatchingFormatID": "2",
			"UserID":           "5",
			"QuestionID":       "18",
			"AnswerID":         "165",
		},
		map[string]string{
			"MatchingFormatID": "2",
			"UserID":           "5",
			"QuestionID":       "19",
			"AnswerID":         "170",
		},
		map[string]string{
			"MatchingFormatID": "2",
			"UserID":           "5",
			"QuestionID":       "20",
			"AnswerID":         "175",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "5",
			"QuestionID":       "21",
			"AnswerID":         "177",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "5",
			"QuestionID":       "22",
			"AnswerID":         "179",
		},

		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "5",
			"QuestionID":       "23",
			"AnswerID":         "182",
		},

		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "5",
			"QuestionID":       "24",
			"AnswerID":         "188",
		},

		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "5",
			"QuestionID":       "25",
			"AnswerID":         "192",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "5",
			"QuestionID":       "26",
			"AnswerID":         "200",
		},

		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "5",
			"QuestionID":       "27",
			"AnswerID":         "203",
		},

		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "5",
			"QuestionID":       "28",
			"AnswerID":         "208",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "5",
			"QuestionID":       "29",
			"AnswerID":         "202",
		},

		map[string]string{
			"MatchingFormatID": "4",
			"UserID":           "5",
			"QuestionID":       "30",
			"AnswerID":         "217",
		},

		map[string]string{
			"MatchingFormatID": "4",
			"UserID":           "5",
			"QuestionID":       "31",
			"AnswerID":         "221",
		},
		map[string]string{
			"MatchingFormatID": "4",
			"UserID":           "5",
			"QuestionID":       "32",
			"AnswerID":         "227",
		},
		map[string]string{
			"MatchingFormatID": "4",
			"UserID":           "5",
			"QuestionID":       "33",
			"AnswerID":         "232",
		},

		map[string]string{
			"MatchingFormatID": "4",
			"UserID":           "5",
			"QuestionID":       "34",
			"AnswerID":         "234",
		},

		//userId 6
		map[string]string{
			"MatchingFormatID": "1",  //　質問ドキュメント参照
			"UserID":           "6",  //　ユーザ情報ドキュメント参照
			"QuestionID":       "1",  //　質問ドキュメント参照
			"AnswerID":         "12", //　回答ドキュメント参照
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "6",
			"QuestionID":       "2",
			"AnswerID":         "51",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "6",
			"QuestionID":       "3",
			"AnswerID":         "54",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "6",
			"QuestionID":       "4",
			"AnswerID":         "57",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "6",
			"QuestionID":       "5",
			"AnswerID":         "78",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "6",
			"QuestionID":       "6",
			"AnswerID":         "71",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "6",
			"QuestionID":       "7",
			"AnswerID":         "85",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "6",
			"QuestionID":       "8",
			"AnswerID":         "125",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "6",
			"QuestionID":       "9",
			"AnswerID":         "132",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "6",
			"QuestionID":       "10",
			"AnswerID":         "136",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "6",
			"QuestionID":       "11",
			"AnswerID":         "139",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "6",
			"QuestionID":       "12",
			"AnswerID":         "143",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "6",
			"QuestionID":       "13",
			"AnswerID":         "147",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "6",
			"QuestionID":       "14",
			"AnswerID":         "153",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "6",
			"QuestionID":       "15",
			"AnswerID":         "155",
		},
		map[string]string{
			"MatchingFormatID": "2",
			"UserID":           "6",
			"QuestionID":       "16",
			"AnswerID":         "158",
		},
		map[string]string{
			"MatchingFormatID": "2",
			"UserID":           "6",
			"QuestionID":       "17",
			"AnswerID":         "162",
		},
		map[string]string{
			"MatchingFormatID": "2",
			"UserID":           "6",
			"QuestionID":       "18",
			"AnswerID":         "165",
		},
		map[string]string{
			"MatchingFormatID": "2",
			"UserID":           "6",
			"QuestionID":       "19",
			"AnswerID":         "172",
		},
		map[string]string{
			"MatchingFormatID": "2",
			"UserID":           "6",
			"QuestionID":       "20",
			"AnswerID":         "175",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "6",
			"QuestionID":       "21",
			"AnswerID":         "177",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "6",
			"QuestionID":       "22",
			"AnswerID":         "178",
		},

		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "6",
			"QuestionID":       "23",
			"AnswerID":         "182",
		},

		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "6",
			"QuestionID":       "24",
			"AnswerID":         "187",
		},

		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "6",
			"QuestionID":       "25",
			"AnswerID":         "192",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "6",
			"QuestionID":       "26",
			"AnswerID":         "200",
		},

		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "6",
			"QuestionID":       "27",
			"AnswerID":         "203",
		},

		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "6",
			"QuestionID":       "28",
			"AnswerID":         "207",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "6",
			"QuestionID":       "29",
			"AnswerID":         "212",
		},

		map[string]string{
			"MatchingFormatID": "4",
			"UserID":           "6",
			"QuestionID":       "30",
			"AnswerID":         "218",
		},

		map[string]string{
			"MatchingFormatID": "4",
			"UserID":           "6",
			"QuestionID":       "31",
			"AnswerID":         "221",
		},
		map[string]string{
			"MatchingFormatID": "4",
			"UserID":           "6",
			"QuestionID":       "32",
			"AnswerID":         "227",
		},
		map[string]string{
			"MatchingFormatID": "4",
			"UserID":           "6",
			"QuestionID":       "33",
			"AnswerID":         "232",
		},

		map[string]string{
			"MatchingFormatID": "4",
			"UserID":           "6",
			"QuestionID":       "34",
			"AnswerID":         "234",
		},
		//userId 7
		map[string]string{
			"MatchingFormatID": "1",  //　質問ドキュメント参照
			"UserID":           "7",  //　ユーザ情報ドキュメント参照
			"QuestionID":       "1",  //　質問ドキュメント参照
			"AnswerID":         "11", //　回答ドキュメント参照
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "7",
			"QuestionID":       "2",
			"AnswerID":         "50",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "7",
			"QuestionID":       "3",
			"AnswerID":         "53",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "7",
			"QuestionID":       "4",
			"AnswerID":         "66",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "7",
			"QuestionID":       "5",
			"AnswerID":         "70",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "7",
			"QuestionID":       "6",
			"AnswerID":         "72",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "7",
			"QuestionID":       "7",
			"AnswerID":         "99",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "7",
			"QuestionID":       "8",
			"AnswerID":         "125",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "7",
			"QuestionID":       "9",
			"AnswerID":         "131",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "7",
			"QuestionID":       "10",
			"AnswerID":         "135",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "7",
			"QuestionID":       "11",
			"AnswerID":         "139",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "7",
			"QuestionID":       "12",
			"AnswerID":         "140",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "7",
			"QuestionID":       "13",
			"AnswerID":         "147",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "7",
			"QuestionID":       "14",
			"AnswerID":         "151",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "7",
			"QuestionID":       "15",
			"AnswerID":         "154",
		},
		map[string]string{
			"MatchingFormatID": "2",
			"UserID":           "7",
			"QuestionID":       "16",
			"AnswerID":         "158",
		},
		map[string]string{
			"MatchingFormatID": "2",
			"UserID":           "7",
			"QuestionID":       "17",
			"AnswerID":         "160",
		},
		map[string]string{
			"MatchingFormatID": "2",
			"UserID":           "7",
			"QuestionID":       "18",
			"AnswerID":         "165",
		},
		map[string]string{
			"MatchingFormatID": "2",
			"UserID":           "7",
			"QuestionID":       "19",
			"AnswerID":         "170",
		},
		map[string]string{
			"MatchingFormatID": "2",
			"UserID":           "7",
			"QuestionID":       "20",
			"AnswerID":         "175",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "7",
			"QuestionID":       "21",
			"AnswerID":         "177",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "7",
			"QuestionID":       "22",
			"AnswerID":         "179",
		},

		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "7",
			"QuestionID":       "23",
			"AnswerID":         "182",
		},

		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "7",
			"QuestionID":       "24",
			"AnswerID":         "188",
		},

		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "7",
			"QuestionID":       "25",
			"AnswerID":         "192",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "7",
			"QuestionID":       "26",
			"AnswerID":         "200",
		},

		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "7",
			"QuestionID":       "27",
			"AnswerID":         "203",
		},

		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "7",
			"QuestionID":       "28",
			"AnswerID":         "208",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "7",
			"QuestionID":       "29",
			"AnswerID":         "202",
		},

		map[string]string{
			"MatchingFormatID": "4",
			"UserID":           "7",
			"QuestionID":       "30",
			"AnswerID":         "217",
		},

		map[string]string{
			"MatchingFormatID": "4",
			"UserID":           "7",
			"QuestionID":       "31",
			"AnswerID":         "221",
		},
		map[string]string{
			"MatchingFormatID": "4",
			"UserID":           "7",
			"QuestionID":       "32",
			"AnswerID":         "227",
		},
		map[string]string{
			"MatchingFormatID": "4",
			"UserID":           "7",
			"QuestionID":       "33",
			"AnswerID":         "232",
		},

		map[string]string{
			"MatchingFormatID": "4",
			"UserID":           "7",
			"QuestionID":       "34",
			"AnswerID":         "234",
		},

		//userId 8
		map[string]string{
			"MatchingFormatID": "1", //　質問ドキュメント参照
			"UserID":           "8", //　ユーザ情報ドキュメント参照
			"QuestionID":       "1", //　質問ドキュメント参照
			"AnswerID":         "1", //　回答ドキュメント参照
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "8",
			"QuestionID":       "2",
			"AnswerID":         "49",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "8",
			"QuestionID":       "3",
			"AnswerID":         "53",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "8",
			"QuestionID":       "4",
			"AnswerID":         "62",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "8",
			"QuestionID":       "5",
			"AnswerID":         "69",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "8",
			"QuestionID":       "6",
			"AnswerID":         "71",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "8",
			"QuestionID":       "7",
			"AnswerID":         "89",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "8",
			"QuestionID":       "8",
			"AnswerID":         "125",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "8",
			"QuestionID":       "9",
			"AnswerID":         "130",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "8",
			"QuestionID":       "10",
			"AnswerID":         "136",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "8",
			"QuestionID":       "11",
			"AnswerID":         "139",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "8",
			"QuestionID":       "12",
			"AnswerID":         "140",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "8",
			"QuestionID":       "13",
			"AnswerID":         "147",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "8",
			"QuestionID":       "14",
			"AnswerID":         "150",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "8",
			"QuestionID":       "15",
			"AnswerID":         "154",
		},
		map[string]string{
			"MatchingFormatID": "2",
			"UserID":           "8",
			"QuestionID":       "16",
			"AnswerID":         "157",
		},
		map[string]string{
			"MatchingFormatID": "2",
			"UserID":           "8",
			"QuestionID":       "17",
			"AnswerID":         "159",
		},
		map[string]string{
			"MatchingFormatID": "2",
			"UserID":           "8",
			"QuestionID":       "18",
			"AnswerID":         "166",
		},
		map[string]string{
			"MatchingFormatID": "2",
			"UserID":           "8",
			"QuestionID":       "19",
			"AnswerID":         "170",
		},
		map[string]string{
			"MatchingFormatID": "2",
			"UserID":           "8",
			"QuestionID":       "20",
			"AnswerID":         "173",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "8",
			"QuestionID":       "21",
			"AnswerID":         "177",
		},

		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "8",
			"QuestionID":       "22",
			"AnswerID":         "178",
		},

		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "8",
			"QuestionID":       "23",
			"AnswerID":         "182",
		},

		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "8",
			"QuestionID":       "24",
			"AnswerID":         "186",
		},

		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "8",
			"QuestionID":       "25",
			"AnswerID":         "197",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "8",
			"QuestionID":       "26",
			"AnswerID":         "200",
		},

		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "8",
			"QuestionID":       "27",
			"AnswerID":         "203",
		},

		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "8",
			"QuestionID":       "28",
			"AnswerID":         "208",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "8",
			"QuestionID":       "29",
			"AnswerID":         "213",
		},

		map[string]string{
			"MatchingFormatID": "4",
			"UserID":           "8",
			"QuestionID":       "30",
			"AnswerID":         "216",
		},
		map[string]string{
			"MatchingFormatID": "4",
			"UserID":           "8",
			"QuestionID":       "31",
			"AnswerID":         "221",
		},
		map[string]string{
			"MatchingFormatID": "4",
			"UserID":           "8",
			"QuestionID":       "32",
			"AnswerID":         "226",
		},
		map[string]string{
			"MatchingFormatID": "4",
			"UserID":           "8",
			"QuestionID":       "33",
			"AnswerID":         "229",
		},

		map[string]string{
			"MatchingFormatID": "4",
			"UserID":           "8",
			"QuestionID":       "34",
			"AnswerID":         "232",
		},

		//userId 9
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "9",
			"QuestionID":       "1",
			"AnswerID":         "26",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "9",
			"QuestionID":       "2",
			"AnswerID":         "49",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "9",
			"QuestionID":       "3",
			"AnswerID":         "54",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "9",
			"QuestionID":       "4",
			"AnswerID":         "57",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "9",
			"QuestionID":       "5",
			"AnswerID":         "68",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "9",
			"QuestionID":       "6",
			"AnswerID":         "71",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "9",
			"QuestionID":       "7",
			"AnswerID":         "101",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "9",
			"QuestionID":       "8",
			"AnswerID":         "124",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "9",
			"QuestionID":       "9",
			"AnswerID":         "131",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "9",
			"QuestionID":       "10",
			"AnswerID":         "136",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "9",
			"QuestionID":       "11",
			"AnswerID":         "139",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "9",
			"QuestionID":       "12",
			"AnswerID":         "140",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "9",
			"QuestionID":       "13",
			"AnswerID":         "147",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "9",
			"QuestionID":       "14",
			"AnswerID":         "151",
		},
		map[string]string{
			"MatchingFormatID": "1",
			"UserID":           "9",
			"QuestionID":       "15",
			"AnswerID":         "154",
		},
		map[string]string{
			"MatchingFormatID": "2",
			"UserID":           "9",
			"QuestionID":       "16",
			"AnswerID":         "251",
		},
		map[string]string{
			"MatchingFormatID": "2",
			"UserID":           "9",
			"QuestionID":       "17",
			"AnswerID":         "252",
		},
		map[string]string{
			"MatchingFormatID": "2",
			"UserID":           "9",
			"QuestionID":       "18",
			"AnswerID":         "253",
		},
		map[string]string{
			"MatchingFormatID": "2",
			"UserID":           "9",
			"QuestionID":       "19",
			"AnswerID":         "254",
		},
		map[string]string{
			"MatchingFormatID": "2",
			"UserID":           "9",
			"QuestionID":       "20",
			"AnswerID":         "255",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "9",
			"QuestionID":       "21",
			"AnswerID":         "256",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "9",
			"QuestionID":       "22",
			"AnswerID":         "257",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "9",
			"QuestionID":       "23",
			"AnswerID":         "258",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "9",
			"QuestionID":       "24",
			"AnswerID":         "259",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "9",
			"QuestionID":       "25",
			"AnswerID":         "260",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "9",
			"QuestionID":       "26",
			"AnswerID":         "261",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "9",
			"QuestionID":       "27",
			"AnswerID":         "262",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "9",
			"QuestionID":       "28",
			"AnswerID":         "263",
		},
		map[string]string{
			"MatchingFormatID": "3",
			"UserID":           "9",
			"QuestionID":       "29",
			"AnswerID":         "264",
		},
		map[string]string{
			"MatchingFormatID": "4",
			"UserID":           "9",
			"QuestionID":       "30",
			"AnswerID":         "265",
		},
		map[string]string{
			"MatchingFormatID": "4",
			"UserID":           "9",
			"QuestionID":       "31",
			"AnswerID":         "266",
		},
		map[string]string{
			"MatchingFormatID": "4",
			"UserID":           "9",
			"QuestionID":       "32",
			"AnswerID":         "267",
		},
		map[string]string{
			"MatchingFormatID": "4",
			"UserID":           "9",
			"QuestionID":       "33",
			"AnswerID":         "268",
		},
		map[string]string{
			"MatchingFormatID": "4",
			"UserID":           "9",
			"QuestionID":       "34",
			"AnswerID":         "269",
		},
	}

	for _, info := range questionAndAnswer_infos {
		matchingFormatID, _ := strconv.Atoi(info["MatchingFormatID"])
		userID, _ := strconv.Atoi(info["UserID"])
		questionID, _ := strconv.Atoi(info["QuestionID"])
		answerID, _ := strconv.Atoi(info["AnswerID"])
		createUserQuestionAndAnswer(model.UserQuestionAndAnswer{
			MatchingFormatID: uint(matchingFormatID),
			UserID:           uint(userID),
			QuestionID:       uint(questionID),
			AnswerID:         uint(answerID),
		})
	}
}

func createUserQuestionAndAnswer(userQuestionAndAnswer model.UserQuestionAndAnswer) {
	userQuestionAndAnswer, err := model.CreateUserQuestionAndAnswer(userQuestionAndAnswer)
	if err != nil {
		panic(err)
	}
	fmt.Printf("userQuestionAndAnswer created\n")
}
