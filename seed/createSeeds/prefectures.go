package createSeeds

import (
	"fmt"

	"github.com/weet/model"
)

func CreateSeedPrefectures() {

	prefectures_infos := []map[string]string{
		map[string]string{
			"Name": "北海道",
		},
		map[string]string{
			"Name": "青森県",
		},
		map[string]string{
			"Name": "岩手県",
		},
		map[string]string{
			"Name": "宮城県",
		},
		map[string]string{
			"Name": "秋田県",
		},
		map[string]string{
			"Name": "山形県",
		},
		map[string]string{
			"Name": "福島県",
		},
		map[string]string{
			"Name": "茨城県",
		},
		map[string]string{
			"Name": "栃木県",
		},
		map[string]string{
			"Name": "群馬県",
		},
		map[string]string{
			"Name": "埼玉県",
		},
		map[string]string{
			"Name": "千葉県",
		},
		map[string]string{
			"Name": "東京都",
		},
		map[string]string{
			"Name": "神奈川県",
		},
		map[string]string{
			"Name": "新潟県",
		},
		map[string]string{
			"Name": "富山県",
		},
		map[string]string{
			"Name": "石川県",
		},
		map[string]string{
			"Name": "福井県",
		},
		map[string]string{
			"Name": "山梨県",
		},
		map[string]string{
			"Name": "長野県",
		},
		map[string]string{
			"Name": "岐阜県",
		},
		map[string]string{
			"Name": "静岡県",
		},
		map[string]string{
			"Name": "愛知県",
		},
		map[string]string{
			"Name": "三重県",
		},
		map[string]string{
			"Name": "滋賀県",
		},
		map[string]string{ //26
			"Name": "京都府",
		},
		map[string]string{ //27
			"Name": "大阪府",
		},
		map[string]string{ //28
			"Name": "兵庫県",
		},
		map[string]string{
			"Name": "奈良県",
		},
		map[string]string{
			"Name": "和歌山県",
		},
		map[string]string{
			"Name": "鳥取県",
		},
		map[string]string{
			"Name": "島根県",
		},
		map[string]string{
			"Name": "岡山県",
		},
		map[string]string{
			"Name": "広島県",
		},
		map[string]string{
			"Name": "山口県",
		},
		map[string]string{
			"Name": "徳島県",
		},
		map[string]string{
			"Name": "香川県",
		},
		map[string]string{
			"Name": "愛媛県",
		},
		map[string]string{
			"Name": "高知県",
		},
		map[string]string{
			"Name": "福岡県",
		},
		map[string]string{
			"Name": "佐賀県",
		},
		map[string]string{
			"Name": "長崎県",
		},
		map[string]string{
			"Name": "熊本県",
		},
		map[string]string{
			"Name": "大分県",
		},
		map[string]string{
			"Name": "宮崎県",
		},
		map[string]string{
			"Name": "鹿児島県",
		},
		map[string]string{
			"Name": "沖縄県",
		},
	}

	for _, info := range prefectures_infos {
		createPrefectures(model.Prefectures{
			Name: info["Name"],
		})
	}
}

func createPrefectures(prefectures model.Prefectures) {
	prefectures, err := model.CreatePrefectures(prefectures)
	if err != nil {
		panic(err)
	}
	fmt.Printf("prefectures created\n")
}
