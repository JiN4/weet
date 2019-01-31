package createSeeds

import (
	"fmt"
	"strconv"

	"github.com/weet/model"
)

// // 自分のユーザー情報
// type User struct {
// 	//Model
// 	gorm.Model
// 	Uid   string
// 	Name  string
// 	Email string
// 	//        val githubStatus : GitHubStatus,// ログインで得られるGitHubの情報によってデータクラスを作る
// }

func CreateSeedUserBasics() {

	userBasics_infos := []map[string]string{
		map[string]string{
			"UserName":    "かまやん",
			"Image1":      "http://test-post.up.n.seesaa.net/test-post/image/weee.jpg-c97ff-thumbnail2.jpg?d=a1",
			"Image2":      "",
			"Image3":      "",
			"Age":         "22",
			"Sex":         "男性",
			"ResidenceID": "28",
			"Hitokoto":    "ガールフレンドを作りたいです！",
			"Comment":     "お酒を飲める人とバーや居酒屋に行きたいです。\n身長180cm以上でアクティブでポジティブなボーイッシュな人募集します！（周りからは理想高すぎっ！？と言われますが書いてみました。）",
		},
		map[string]string{
			"UserName":    "田中 修",
			"Image1":      "https://www.pakutaso.com/shared/img/thumb/YUSEI1030IMGL7452_TP_V.jpg",
			"Image2":      "",
			"Image3":      "",
			"Age":         "34",
			"Sex":         "男性",
			"ResidenceID": "28",
			"Hitokoto":    "雑学マスターと呼んでください",
			"Comment":     "この34年、この世のためにならない知識を蓄えに蓄えてきました。世界各地の雑学を知っています。興味のある方、話しませんか？",
		},
		map[string]string{
			"UserName":    "荻野 秀平",
			"Image1":      "https://www.pakutaso.com/shared/img/thumb/00_PP05_PP_TP_V.jpg",
			"Image2":      "",
			"Image3":      "",
			"Age":         "26",
			"Sex":         "男性",
			"ResidenceID": "27",
			"Hitokoto":    "おしゃれなカフェ知ってます！",
			"Comment":     "都会に憧れ10年、とりあえず近場のおしゃれスポットを探し、標準語の練習をする日々を送っています。",
		},
		map[string]string{
			"UserName":    "球蹴 蹴斗",
			"Image1":      "https://www.pakutaso.com/shared/img/thumb/KENTA8452dds_TP_V.jpg",
			"Image2":      "",
			"Image3":      "",
			"Age":         "20",
			"Sex":         "男性",
			"ResidenceID": "28",
			"Hitokoto":    "サッカーできる人探してます！",
			"Comment":     "一緒にサッカーできる友達を探しにきました！\n休日とかに僕の知り合いと一緒にサッカーしませんか！？",
		},
		map[string]string{
			"UserName":    "球投 一球",
			"Image1":      "https://images.unsplash.com/photo-1489460427746-b6296f4bc3f5?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=634&q=80",
			"Image2":      "",
			"Image3":      "",
			"Age":         "26",
			"Sex":         "男性",
			"ResidenceID": "27",
			"Hitokoto":    "ストライク王に俺はなる！",
			"Comment":     "全ての打者から３振をとる、すなわちそれ最強。",
		},
		map[string]string{
			"UserName":    "七瀬 奏",
			"Image1":      "https://www.pakutaso.com/shared/img/thumb/YUKI151117060I9A1562_TP_V.jpg",
			"Image2":      "",
			"Image3":      "",
			"Age":         "21",
			"Sex":         "女性",
			"ResidenceID": "28",
			"Hitokoto":    "音楽が趣味です",
			"Comment":     "高校の時、バイオリンをしてました。\nそのときからコンサートにいくのが好きで、一緒に聴きにいける人と知り合えたらいいなと思ってます！",
		},
		map[string]string{
			"UserName":    "立花 彩香",
			"Image1":      "https://www.pakutaso.com/shared/img/thumb/TSURU1891A041_TP_V.jpg",
			"Image2":      "",
			"Image3":      "",
			"Age":         "27",
			"Sex":         "女性",
			"ResidenceID": "28",
			"Hitokoto":    "誰かー！かまってくれー！",
			"Comment":     "かまちょかまちょかまちょかまちょかまちょかまちょ",
		},
		map[string]string{
			"UserName":    "藤宮 紗代",
			"Image1":      "https://www.pakutaso.com/shared/img/thumb/kangoIMGL7876_TP_V.jpg",
			"Image2":      "",
			"Image3":      "",
			"Age":         "26",
			"Sex":         "女性",
			"ResidenceID": "26",
			"Hitokoto":    "結婚を考えている男性探してます",
			"Comment":     "そろそろ結婚も考えているのですが出会いの場も少なく、このアプリを入れてみることにしました。結婚を考えたお付き合いをしてくれる人、声をかけてくれると嬉しいです。",
		},
		map[string]string{
			"UserName":    "小鳥遊 詩織",
			"Image1":      "https://www.pakutaso.com/shared/img/thumb/YUKA150912596015_TP_V.jpg",
			"Image2":      "",
			"Image3":      "",
			"Age":         "24",
			"Sex":         "女性",
			"ResidenceID": "26",
			"Hitokoto":    "本が好きです。",
			"Comment":     "気に入った本の感想を言い合える友達を探しています。\n好きな本を教えてくれると読んでみるかもしれません。",
		},
	}

	for _, info := range userBasics_infos {
		age, _ := strconv.Atoi(info["Age"])
		residenceID, _ := strconv.Atoi(info["ResidenceID"])
		createUserBasics(model.UserBasics{
			UserName:    info["UserName"],
			Image1:      info["Image1"],
			Image2:      info["Image2"],
			Image3:      info["Image3"],
			Age:         uint(age),
			ResidenceID: uint(residenceID),
			Sex:         info["Sex"],
			Hitokoto:    info["Hitokoto"],
			Comment:     info["Comment"],
		})
	}
}

func createUserBasics(userBasics model.UserBasics) {
	userBasics, err := model.CreateUserBasics(userBasics)
	if err != nil {
		panic(err)
	}
	fmt.Printf("userBasics created\n")
}
