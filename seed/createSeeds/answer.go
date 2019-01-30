package createSeeds

import (
	"fmt"
	"strconv"

	"github.com/weet/model"
)

func CreateSeedAnswers() {

	answers_infos := []map[string]string{
		map[string]string{
			"QuestionID": "1",
			"Name":       "北海道",
		},
		map[string]string{
			"QuestionID": "1",
			"Name":       "青森県",
		},
		map[string]string{
			"QuestionID": "1",
			"Name":       "岩手県",
		},
		map[string]string{
			"QuestionID": "1",
			"Name":       "宮城県",
		},
		map[string]string{ //5
			"QuestionID": "1",
			"Name":       "秋田県",
		},
		map[string]string{
			"QuestionID": "1",
			"Name":       "山形県",
		},
		map[string]string{
			"QuestionID": "1",
			"Name":       "福島県",
		},
		map[string]string{
			"QuestionID": "1",
			"Name":       "茨城県",
		},
		map[string]string{
			"QuestionID": "1",
			"Name":       "栃木県",
		},
		map[string]string{ //10
			"QuestionID": "1",
			"Name":       "群馬県",
		},
		map[string]string{
			"QuestionID": "1",
			"Name":       "埼玉県",
		},
		map[string]string{
			"QuestionID": "1",
			"Name":       "千葉県",
		},
		map[string]string{
			"QuestionID": "1",
			"Name":       "東京都",
		},
		map[string]string{
			"QuestionID": "1",
			"Name":       "神奈川県",
		},
		map[string]string{ //15
			"QuestionID": "1",
			"Name":       "新潟県",
		},
		map[string]string{
			"QuestionID": "1",
			"Name":       "富山県",
		},
		map[string]string{
			"QuestionID": "1",
			"Name":       "石川県",
		},
		map[string]string{
			"QuestionID": "1",
			"Name":       "福井県",
		},
		map[string]string{
			"QuestionID": "1",
			"Name":       "山梨県",
		},
		map[string]string{ //20
			"QuestionID": "1",
			"Name":       "長野県",
		},
		map[string]string{
			"QuestionID": "1",
			"Name":       "岐阜県",
		},
		map[string]string{
			"QuestionID": "1",
			"Name":       "静岡県",
		},
		map[string]string{
			"QuestionID": "1",
			"Name":       "愛知県",
		},
		map[string]string{
			"QuestionID": "1",
			"Name":       "三重県",
		},
		map[string]string{ //15
			"QuestionID": "1",
			"Name":       "滋賀県",
		},
		map[string]string{ //26
			"QuestionID": "1",
			"Name":       "京都府",
		},
		map[string]string{ //27
			"QuestionID": "1",
			"Name":       "大阪府",
		},
		map[string]string{ //28
			"QuestionID": "1",
			"Name":       "兵庫県",
		},
		map[string]string{
			"QuestionID": "1",
			"Name":       "奈良県",
		},
		map[string]string{ //30
			"QuestionID": "1",
			"Name":       "和歌山県",
		},
		map[string]string{
			"QuestionID": "1",
			"Name":       "鳥取県",
		},
		map[string]string{
			"QuestionID": "1",
			"Name":       "島根県",
		},
		map[string]string{
			"QuestionID": "1",
			"Name":       "岡山県",
		},
		map[string]string{
			"QuestionID": "1",
			"Name":       "広島県",
		},
		map[string]string{ //35
			"QuestionID": "1",
			"Name":       "山口県",
		},
		map[string]string{
			"QuestionID": "1",
			"Name":       "徳島県",
		},
		map[string]string{
			"QuestionID": "1",
			"Name":       "香川県",
		},
		map[string]string{
			"QuestionID": "1",
			"Name":       "愛媛県",
		},
		map[string]string{
			"QuestionID": "1",
			"Name":       "高知県",
		},
		map[string]string{ //40
			"QuestionID": "1",
			"Name":       "福岡県",
		},
		map[string]string{
			"QuestionID": "1",
			"Name":       "佐賀県",
		},
		map[string]string{
			"QuestionID": "1",
			"Name":       "長崎県",
		},
		map[string]string{
			"QuestionID": "1",
			"Name":       "熊本県",
		},
		map[string]string{
			"QuestionID": "1",
			"Name":       "大分県",
		},
		map[string]string{ //45
			"QuestionID": "1",
			"Name":       "宮崎県",
		},
		map[string]string{
			"QuestionID": "1",
			"Name":       "鹿児島県",
		},
		map[string]string{
			"QuestionID": "1",
			"Name":       "沖縄県",
		},
		map[string]string{
			"QuestionID": "2",
			"Name":       "A型",
		},
		map[string]string{
			"QuestionID": "2",
			"Name":       "B型",
		},
		map[string]string{ //50
			"QuestionID": "2",
			"Name":       "O型",
		},
		map[string]string{
			"QuestionID": "2",
			"Name":       "AB型",
		},
		map[string]string{
			"QuestionID": "3",
			"Name":       "高校卒",
		},
		map[string]string{
			"QuestionID": "3",
			"Name":       "短大/専門学校卒",
		},
		map[string]string{
			"QuestionID": "3",
			"Name":       "大学卒",
		},
		map[string]string{ //55
			"QuestionID": "3",
			"Name":       "大学院卒",
		},
		map[string]string{
			"QuestionID": "3",
			"Name":       "その他",
		},
		map[string]string{
			"QuestionID": "4",
			"Name":       "オフィス系",
		},
		map[string]string{
			"QuestionID": "4",
			"Name":       "語学系",
		},
		map[string]string{
			"QuestionID": "4",
			"Name":       "製造・物流・作業系",
		},
		map[string]string{ //60
			"QuestionID": "4",
			"Name":       "研究・開発系",
		},
		map[string]string{
			"QuestionID": "4",
			"Name":       "IT・エンジニア系",
		},
		map[string]string{
			"QuestionID": "4",
			"Name":       "接客・販売系",
		},
		map[string]string{
			"QuestionID": "4",
			"Name":       "コールセンター系",
		},
		map[string]string{
			"QuestionID": "4",
			"Name":       "営業系",
		},
		map[string]string{ //65
			"QuestionID": "4",
			"Name":       "金融系",
		},
		map[string]string{
			"QuestionID": "4",
			"Name":       "医療系",
		},
		map[string]string{
			"QuestionID": "4",
			"Name":       "その他",
		},
		map[string]string{
			"QuestionID": "5",
			"Name":       "土日",
		},
		map[string]string{
			"QuestionID": "5",
			"Name":       "平日",
		},
		map[string]string{ //70
			"QuestionID": "5",
			"Name":       "不定期",
		},
		map[string]string{
			"QuestionID": "6",
			"Name":       "インドア派",
		},
		map[string]string{
			"QuestionID": "6",
			"Name":       "アウトドア派",
		},
		map[string]string{
			"QuestionID": "7",
			"Name":       "140cm以下",
		},
		map[string]string{
			"QuestionID": "7",
			"Name":       "141cm",
		},
		map[string]string{ //75
			"QuestionID": "7",
			"Name":       "142cm",
		},
		map[string]string{
			"QuestionID": "7",
			"Name":       "143cm",
		},
		map[string]string{
			"QuestionID": "7",
			"Name":       "144cm",
		},
		map[string]string{
			"QuestionID": "7",
			"Name":       "145cm",
		},
		map[string]string{
			"QuestionID": "7",
			"Name":       "146cm",
		},
		map[string]string{ //80
			"QuestionID": "7",
			"Name":       "147cm",
		},
		map[string]string{
			"QuestionID": "7",
			"Name":       "148cm",
		},
		map[string]string{
			"QuestionID": "7",
			"Name":       "149cm",
		},
		map[string]string{
			"QuestionID": "7",
			"Name":       "150cm",
		},
		map[string]string{
			"QuestionID": "7",
			"Name":       "151cm",
		},
		map[string]string{ //85
			"QuestionID": "7",
			"Name":       "152cm",
		},
		map[string]string{
			"QuestionID": "7",
			"Name":       "153cm",
		},
		map[string]string{
			"QuestionID": "7",
			"Name":       "154cm",
		},
		map[string]string{
			"QuestionID": "7",
			"Name":       "155cm",
		},
		map[string]string{
			"QuestionID": "7",
			"Name":       "156cm",
		},
		map[string]string{ //90
			"QuestionID": "7",
			"Name":       "157cm",
		},
		map[string]string{
			"QuestionID": "7",
			"Name":       "158",
		},
		map[string]string{
			"QuestionID": "7",
			"Name":       "159cm",
		},
		map[string]string{
			"QuestionID": "7",
			"Name":       "160cm",
		},
		map[string]string{
			"QuestionID": "7",
			"Name":       "161cm",
		},
		map[string]string{ //95
			"QuestionID": "7",
			"Name":       "162cm",
		},
		map[string]string{
			"QuestionID": "7",
			"Name":       "163cm",
		},
		map[string]string{
			"QuestionID": "7",
			"Name":       "164cm",
		},
		map[string]string{
			"QuestionID": "7",
			"Name":       "165cm",
		},
		map[string]string{
			"QuestionID": "7",
			"Name":       "166cm",
		},
		map[string]string{ //100
			"QuestionID": "7",
			"Name":       "167cm",
		},
		map[string]string{
			"QuestionID": "7",
			"Name":       "168cm",
		},
		map[string]string{
			"QuestionID": "7",
			"Name":       "169cm",
		},
		map[string]string{
			"QuestionID": "7",
			"Name":       "170cm",
		},
		map[string]string{
			"QuestionID": "7",
			"Name":       "171cm",
		},
		map[string]string{ //105
			"QuestionID": "7",
			"Name":       "172cm",
		},
		map[string]string{
			"QuestionID": "7",
			"Name":       "173cm",
		},
		map[string]string{
			"QuestionID": "7",
			"Name":       "174cm",
		},
		map[string]string{
			"QuestionID": "7",
			"Name":       "175cm",
		},
		map[string]string{
			"QuestionID": "7",
			"Name":       "176cm",
		},
		map[string]string{ //110
			"QuestionID": "7",
			"Name":       "177cm",
		},
		map[string]string{
			"QuestionID": "7",
			"Name":       "178cm",
		},
		map[string]string{
			"QuestionID": "7",
			"Name":       "179cm",
		},
		map[string]string{
			"QuestionID": "7",
			"Name":       "180cm",
		},
		map[string]string{
			"QuestionID": "7",
			"Name":       "181cm",
		},
		map[string]string{ //115
			"QuestionID": "7",
			"Name":       "182cm",
		},
		map[string]string{
			"QuestionID": "7",
			"Name":       "183cm",
		},
		map[string]string{
			"QuestionID": "7",
			"Name":       "184cm",
		},
		map[string]string{
			"QuestionID": "7",
			"Name":       "185cm",
		},
		map[string]string{
			"QuestionID": "7",
			"Name":       "186cm",
		},
		map[string]string{ //120
			"QuestionID": "7",
			"Name":       "187cm",
		},
		map[string]string{
			"QuestionID": "7",
			"Name":       "188cm",
		},
		map[string]string{
			"QuestionID": "7",
			"Name":       "189cm",
		},
		map[string]string{
			"QuestionID": "7",
			"Name":       "190cm以上",
		},
		map[string]string{
			"QuestionID": "8",
			"Name":       "細め",
		},
		map[string]string{ //125
			"QuestionID": "8",
			"Name":       "スレンダー",
		},
		map[string]string{
			"QuestionID": "8",
			"Name":       "普通",
		},
		map[string]string{
			"QuestionID": "8",
			"Name":       "ぽっちゃり",
		},
		map[string]string{
			"QuestionID": "8",
			"Name":       "グラマー",
		},
		map[string]string{
			"QuestionID": "8",
			"Name":       "筋肉質",
		},
		map[string]string{ //130
			"QuestionID": "9",
			"Name":       "飲む",
		},
		map[string]string{
			"QuestionID": "9",
			"Name":       "たまに飲む",
		},
		map[string]string{
			"QuestionID": "9",
			"Name":       "あまり飲まない",
		},
		map[string]string{
			"QuestionID": "9",
			"Name":       "飲まない",
		},
		map[string]string{
			"QuestionID": "10",
			"Name":       "吸う",
		},
		map[string]string{ //135
			"QuestionID": "10",
			"Name":       "たまに吸う",
		},
		map[string]string{
			"QuestionID": "10",
			"Name":       "吸わない",
		},
		map[string]string{
			"QuestionID": "11",
			"Name":       "する",
		},
		map[string]string{
			"QuestionID": "11",
			"Name":       "たまにする",
		},
		map[string]string{
			"QuestionID": "11",
			"Name":       "しない",
		},
		map[string]string{ //140
			"QuestionID": "12",
			"Name":       "車",
		},
		map[string]string{
			"QuestionID": "12",
			"Name":       "バイク",
		},
		map[string]string{
			"QuestionID": "12",
			"Name":       "車・バイク",
		},
		map[string]string{
			"QuestionID": "12",
			"Name":       "ない",
		},
		map[string]string{
			"QuestionID": "13",
			"Name":       "10,000円以下",
		},
		map[string]string{ //145
			"QuestionID": "13",
			"Name":       "10,000円〜20,000円",
		},
		map[string]string{
			"QuestionID": "13",
			"Name":       "20,000円〜30,000円",
		},
		map[string]string{
			"QuestionID": "13",
			"Name":       "30,000円〜40,000円",
		},
		map[string]string{
			"QuestionID": "13",
			"Name":       "40,000円〜50,000円",
		},
		map[string]string{
			"QuestionID": "13",
			"Name":       "50,000円以上",
		},
		map[string]string{ //150
			"QuestionID": "14",
			"Name":       "話す方が好き",
		},
		map[string]string{
			"QuestionID": "14",
			"Name":       "聞く方が好き",
		},
		map[string]string{
			"QuestionID": "14",
			"Name":       "話すのも聞くのも好き",
		},
		map[string]string{
			"QuestionID": "14",
			"Name":       "会話は得意ではない",
		},
		map[string]string{
			"QuestionID": "15",
			"Name":       "ある程度のメッセージ交換を重ねてから",
		},
		map[string]string{ //155
			"QuestionID": "15",
			"Name":       "お互いの気が合えば",
		},
		map[string]string{
			"QuestionID": "15",
			"Name":       "とりあえず会って話したい",
		},
		map[string]string{
			"QuestionID": "16",
			"Name":       "あり",
		},
		map[string]string{
			"QuestionID": "16",
			"Name":       "なし",
		},
		map[string]string{
			"QuestionID": "17",
			"Name":       "毎日",
		},
		map[string]string{ //160
			"QuestionID": "17",
			"Name":       "週2〜3回",
		},
		map[string]string{
			"QuestionID": "17",
			"Name":       "週4〜6回",
		},
		map[string]string{
			"QuestionID": "17",
			"Name":       "その他",
		},
		map[string]string{
			"QuestionID": "18",
			"Name":       "2〜3日に一回",
		},
		map[string]string{
			"QuestionID": "18",
			"Name":       "4〜6回に一回",
		},
		map[string]string{ //165
			"QuestionID": "18",
			"Name":       "1週間に一回",
		},
		map[string]string{
			"QuestionID": "18",
			"Name":       "2週間に一回",
		},
		map[string]string{
			"QuestionID": "18",
			"Name":       "その他",
		},
		map[string]string{
			"QuestionID": "19",
			"Name":       "男性が全額払う",
		},
		map[string]string{
			"QuestionID": "19",
			"Name":       "男性が多めに払う",
		},
		map[string]string{ //170
			"QuestionID": "19",
			"Name":       "割り勘",
		},
		map[string]string{
			"QuestionID": "19",
			"Name":       "持っている方が払う",
		},
		map[string]string{
			"QuestionID": "19",
			"Name":       "場所のことも考え、話し合って決める",
		},
		map[string]string{
			"QuestionID": "20",
			"Name":       "したい",
		},
		map[string]string{
			"QuestionID": "20",
			"Name":       "したくない",
		},
		map[string]string{ //175
			"QuestionID": "20",
			"Name":       "交際してから考える",
		},
		map[string]string{
			"QuestionID": "21",
			"Name":       "あり",
		},
		map[string]string{
			"QuestionID": "21",
			"Name":       "なし",
		},
		map[string]string{
			"QuestionID": "22",
			"Name":       "積極的に参加したい",
		},
		map[string]string{
			"QuestionID": "22",
			"Name":       "余裕があれば参加したい",
		},
		map[string]string{ //180
			"QuestionID": "22",
			"Name":       "状況を見て参加したい",
		},
		map[string]string{
			"QuestionID": "22",
			"Name":       "相手に任せたい",
		},
		map[string]string{
			"QuestionID": "23",
			"Name":       "積極的に参加したい",
		},
		map[string]string{
			"QuestionID": "23",
			"Name":       "余裕があれば参加したい",
		},
		map[string]string{
			"QuestionID": "23",
			"Name":       "状況を見て参加したい",
		},
		map[string]string{ //185
			"QuestionID": "23",
			"Name":       "相手に任せたい",
		},
		map[string]string{
			"QuestionID": "24",
			"Name":       "200万円未満",
		},
		map[string]string{
			"QuestionID": "24",
			"Name":       "200万円〜400万円",
		},
		map[string]string{
			"QuestionID": "24",
			"Name":       "400万円〜600万円",
		},
		map[string]string{
			"QuestionID": "24",
			"Name":       "600万円〜800万円",
		},
		map[string]string{ //190
			"QuestionID": "24",
			"Name":       "800万円〜1000万円",
		},
		map[string]string{
			"QuestionID": "24",
			"Name":       "1,000万円以上",
		},
		map[string]string{
			"QuestionID": "25",
			"Name":       "教会式",
		},
		map[string]string{
			"QuestionID": "25",
			"Name":       "キリスト教式",
		},
		map[string]string{
			"QuestionID": "25",
			"Name":       "神前式",
		},
		map[string]string{ //195
			"QuestionID": "25",
			"Name":       "仏前式",
		},
		map[string]string{
			"QuestionID": "25",
			"Name":       "人前式",
		},
		map[string]string{
			"QuestionID": "25",
			"Name":       "現在こだわりはない",
		},
		map[string]string{
			"QuestionID": "25",
			"Name":       "しない",
		},
		map[string]string{
			"QuestionID": "26",
			"Name":       "いる",
		},
		map[string]string{ //200
			"QuestionID": "26",
			"Name":       "いない",
		},
		map[string]string{
			"QuestionID": "27",
			"Name":       "1年未満",
		},
		map[string]string{
			"QuestionID": "27",
			"Name":       "1年目",
		},
		map[string]string{
			"QuestionID": "27",
			"Name":       "2年目",
		},
		map[string]string{
			"QuestionID": "27",
			"Name":       "3年目",
		},
		map[string]string{ //205
			"QuestionID": "27",
			"Name":       "3年目以降",
		},
		map[string]string{
			"QuestionID": "27",
			"Name":       "いらない",
		},
		map[string]string{
			"QuestionID": "28",
			"Name":       "1人",
		},
		map[string]string{
			"QuestionID": "28",
			"Name":       "2人",
		},
		map[string]string{
			"QuestionID": "28",
			"Name":       "3人以上",
		},
		map[string]string{ //210
			"QuestionID": "28",
			"Name":       "いらない",
		},
		map[string]string{
			"QuestionID": "29",
			"Name":       "あり",
		},
		map[string]string{
			"QuestionID": "29",
			"Name":       "したいなら考える",
		},
		map[string]string{
			"QuestionID": "29",
			"Name":       "なし",
		},
		map[string]string{
			"QuestionID": "30",
			"Name":       "金銭面の負担を減らしたい",
		},
		map[string]string{ //215
			"QuestionID": "30",
			"Name":       "家事の負担を減らしたい",
		},
		map[string]string{
			"QuestionID": "30",
			"Name":       "誰かがいる安心感を得たいから",
		},
		map[string]string{
			"QuestionID": "30",
			"Name":       "相談・愚痴が言える人と住みたい",
		},
		map[string]string{
			"QuestionID": "30",
			"Name":       "きまぐれ",
		},
		map[string]string{
			"QuestionID": "30",
			"Name":       "その他",
		},
		map[string]string{ //220
			"QuestionID": "31",
			"Name":       "半年以下",
		},
		map[string]string{
			"QuestionID": "31",
			"Name":       "半年〜1年",
		},
		map[string]string{
			"QuestionID": "31",
			"Name":       "1年〜2年",
		},
		map[string]string{
			"QuestionID": "31",
			"Name":       "2年以上",
		},
		map[string]string{
			"QuestionID": "31",
			"Name":       "未定",
		},
		map[string]string{ //225
			"QuestionID": "32",
			"Name":       "自分の部屋を貸したい",
		},
		map[string]string{
			"QuestionID": "32",
			"Name":       "誰かの部屋を借りたい",
		},
		map[string]string{
			"QuestionID": "32",
			"Name":       "新しい物件を二人で借りたい",
		},
		map[string]string{
			"QuestionID": "32",
			"Name":       "その他",
		},
		map[string]string{
			"QuestionID": "33",
			"Name":       "大丈夫",
		},
		map[string]string{ //230
			"QuestionID": "33",
			"Name":       "躾されているなら大丈夫",
		},
		map[string]string{
			"QuestionID": "33",
			"Name":       "禁止",
		},
		map[string]string{
			"QuestionID": "34",
			"Name":       "歓迎",
		},
		map[string]string{
			"QuestionID": "34",
			"Name":       "態度や回数など節度を守れるなら",
		},
		map[string]string{
			"QuestionID": "34",
			"Name":       "あまり好まない",
		},
		map[string]string{ //235
			"QuestionID": "34",
			"Name":       "禁止",
		},
		map[string]string{ //236
			"QuestionID": "1",
			"Name":       "-",
		},
		map[string]string{ //237
			"QuestionID": "2",
			"Name":       "-",
		},
		map[string]string{ //238
			"QuestionID": "3",
			"Name":       "-",
		},
		map[string]string{ //239
			"QuestionID": "4",
			"Name":       "-",
		},
		map[string]string{ //240
			"QuestionID": "5",
			"Name":       "-",
		},
		map[string]string{ //241
			"QuestionID": "6",
			"Name":       "-",
		},
		map[string]string{ //242
			"QuestionID": "7",
			"Name":       "-",
		},
		map[string]string{ //243
			"QuestionID": "8",
			"Name":       "-",
		},
		map[string]string{ //244
			"QuestionID": "9",
			"Name":       "-",
		},
		map[string]string{ //245
			"QuestionID": "10",
			"Name":       "-",
		},
		map[string]string{ //246
			"QuestionID": "11",
			"Name":       "-",
		},
		map[string]string{ //247
			"QuestionID": "12",
			"Name":       "-",
		},
		map[string]string{ //248
			"QuestionID": "13",
			"Name":       "-",
		},
		map[string]string{ //249
			"QuestionID": "14",
			"Name":       "-",
		},
		map[string]string{ //250
			"QuestionID": "15",
			"Name":       "-",
		},
		map[string]string{ //251
			"QuestionID": "16",
			"Name":       "-",
		},
		map[string]string{ //252
			"QuestionID": "17",
			"Name":       "-",
		},
		map[string]string{ //253
			"QuestionID": "18",
			"Name":       "-",
		},
		map[string]string{ //254
			"QuestionID": "19",
			"Name":       "-",
		},
		map[string]string{ //255
			"QuestionID": "20",
			"Name":       "-",
		},
		map[string]string{ //256
			"QuestionID": "21",
			"Name":       "-",
		},
		map[string]string{ //257
			"QuestionID": "22",
			"Name":       "-",
		},
		map[string]string{ //258
			"QuestionID": "23",
			"Name":       "-",
		},
		map[string]string{ //259
			"QuestionID": "24",
			"Name":       "-",
		},
		map[string]string{ //260
			"QuestionID": "25",
			"Name":       "-",
		},
		map[string]string{ //261
			"QuestionID": "26",
			"Name":       "-",
		},
		map[string]string{ //262
			"QuestionID": "27",
			"Name":       "-",
		},
		map[string]string{ //263
			"QuestionID": "28",
			"Name":       "-",
		},
		map[string]string{ //264
			"QuestionID": "29",
			"Name":       "-",
		},
		map[string]string{ //265
			"QuestionID": "30",
			"Name":       "-",
		},
		map[string]string{ //266
			"QuestionID": "31",
			"Name":       "-",
		},
		map[string]string{ //267
			"QuestionID": "32",
			"Name":       "-",
		},
		map[string]string{ //268
			"QuestionID": "33",
			"Name":       "-",
		},
		map[string]string{ //269
			"QuestionID": "34",
			"Name":       "-",
		},
	}

	for _, info := range answers_infos {
		questionID, _ := strconv.Atoi(info["QuestionID"])
		createAnswer(model.Answer{
			QuestionID: uint(questionID),
			Name:       info["Name"],
		})
	}
}

func createAnswer(answer model.Answer) {
	answer, err := model.CreateAnswer(answer)
	if err != nil {
		panic(err)
	}
	fmt.Printf("answer created\n")
}
