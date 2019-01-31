# WEET APIサーバー

## APIドキュメント

### ユーザ情報を全て取得する[GET]

#### Call
```
54.238.92.95:8080/api/v1/user/{user_id}
```

|Parameter|Value|type|
|:-:|:-:|:-:|
|user_id|1|int|

#### Response

+ Response 200 (application/json)

<details>
    <summary>Response Body</summary>
    <pre>
    <code>
    {
        "user_basics": {
            "matching_format_name": "基本情報",
            "user_name": "かまやん",
            "image1": "http://test-post.up.seesaa.net/image/weee.jpg-c97ff.jpg",
            "image2": "",
            "image3": "",
            "sex": "男性",
            "age": 22,
            "residence": "兵庫県",
            "hitokoto": "ガールフレンドを作りたいです！",
            "comment": "お酒を飲める人とバーや居酒屋に行きたいです。\n身長180cm以上でアクティブでポジティブなボーイッシュな人募集します！（周りからは理想高すぎっ！？と言われますが書いてみました。）"
        },
        "user_specials": [
            {
                "matching_format_name": "友達",
                "user_questions_and_answers": [
                    {
                        "question_id": 1,
                        "question_name": "出身地",
                        "answer_name": "兵庫県"
                    },
                    {
                        "question_id": 2,
                        "question_name": "血液型",
                        "answer_name": "A型"
                    },
                    {
                        "question_id": 3,
                        "question_name": "学歴",
                        "answer_name": "短大/専門学校卒"
                    },
                    {
                        "question_id": 4,
                        "question_name": "職種",
                        "answer_name": "IT・エンジニア系"
                    },
                    {
                        "question_id": 5,
                        "question_name": "休日の曜日",
                        "answer_name": "土日"
                    },
                    {
                        "question_id": 6,
                        "question_name": "休日の過ごし方",
                        "answer_name": "インドア派"
                    },
                    {
                        "question_id": 7,
                        "question_name": "身長",
                        "answer_name": "182cm"
                    },
                    {
                        "question_id": 8,
                        "question_name": "体型",
                        "answer_name": "細め"
                    },
                    {
                        "question_id": 9,
                        "question_name": "お酒",
                        "answer_name": "たまに飲む"
                    },
                    {
                        "question_id": 10,
                        "question_name": "煙草",
                        "answer_name": "吸わない"
                    },
                    {
                        "question_id": 11,
                        "question_name": "ギャンブル",
                        "answer_name": "しない"
                    },
                    {
                        "question_id": 12,
                        "question_name": "免許",
                        "answer_name": "ない"
                    },
                    {
                        "question_id": 13,
                        "question_name": "一ヶ月の娯楽費",
                        "answer_name": "20,000円〜30,000円"
                    },
                    {
                        "question_id": 14,
                        "question_name": "会話",
                        "answer_name": "話すのも聞くのも好き"
                    },
                    {
                        "question_id": 15,
                        "question_name": "出会うまでの希望",
                        "answer_name": "お互いの気が合えば"
                    }
                ]
            },
            {
                "matching_format_name": "恋愛",
                "user_questions_and_answers": [
                    {
                        "question_id": 16,
                        "question_name": "交際経験",
                        "answer_name": "あり"
                    },
                    {
                        "question_id": 17,
                        "question_name": "メッセージ交換の頻度",
                        "answer_name": "週2〜3回"
                    },
                    {
                        "question_id": 18,
                        "question_name": "デートの頻度",
                        "answer_name": "1週間に一回"
                    },
                    {
                        "question_id": 19,
                        "question_name": "初回デートの費用",
                        "answer_name": "男性が多めに払う"
                    },
                    {
                        "question_id": 20,
                        "question_name": "同棲の希望",
                        "answer_name": "交際してから考える"
                    }
                ]
            },
            {
                "matching_format_name": "婚活",
                "user_questions_and_answers": [
                    {
                        "question_id": 21,
                        "question_name": "結婚経験",
                        "answer_name": "なし"
                    },
                    {
                        "question_id": 22,
                        "question_name": "家事",
                        "answer_name": "余裕があれば参加したい"
                    },
                    {
                        "question_id": 23,
                        "question_name": "育児",
                        "answer_name": "積極的に参加したい"
                    },
                    {
                        "question_id": 24,
                        "question_name": "年収",
                        "answer_name": "200万円〜400万円"
                    },
                    {
                        "question_id": 25,
                        "question_name": "結婚式",
                        "answer_name": "現在こだわりはない"
                    },
                    {
                        "question_id": 26,
                        "question_name": "子供の有無",
                        "answer_name": "いない"
                    },
                    {
                        "question_id": 27,
                        "question_name": "子供はいつ欲しい？",
                        "answer_name": "2年目"
                    },
                    {
                        "question_id": 28,
                        "question_name": "子供は何人欲しい？",
                        "answer_name": "2人"
                    },
                    {
                        "question_id": 29,
                        "question_name": "親との同居",
                        "answer_name": "したいなら考える"
                    }
                ]
            },
            {
                "matching_format_name": "ルームメイト",
                "user_questions_and_answers": [
                    {
                        "question_id": 30,
                        "question_name": "一番の目的",
                        "answer_name": "金銭面の負担を減らしたい"
                    },
                    {
                        "question_id": 31,
                        "question_name": "期間",
                        "answer_name": "半年〜1年"
                    },
                    {
                        "question_id": 32,
                        "question_name": "部屋の貸し借り",
                        "answer_name": "誰かの部屋を借りたい"
                    },
                    {
                        "question_id": 33,
                        "question_name": "ペット",
                        "answer_name": "禁止"
                    },
                    {
                        "question_id": 34,
                        "question_name": "来客",
                        "answer_name": "態度や回数など節度を守れるなら"
                    }
                ]
            }
        ],
        "user_ideal_specials": [
            {
                "matching_format_name": "友達",
                "user_ideal_questions_and_answers": [
                    {
                        "ideal_question_id": 1,
                        "ideal_question_name": "出身地",
                        "ideal_answer_name": "大阪府,兵庫県"
                    },
                    {
                        "ideal_question_id": 2,
                        "ideal_question_name": "血液型",
                        "ideal_answer_name": "B型,O型"
                    },
                    {
                        "ideal_question_id": 3,
                        "ideal_question_name": "学歴",
                        "ideal_answer_name": "短大/専門学校卒,大学卒,大学院卒"
                    },
                    {
                        "ideal_question_id": 4,
                        "ideal_question_name": "職種",
                        "ideal_answer_name": "オフィス系,研究・開発系,IT・エンジニア系,医療系"
                    },
                    {
                        "ideal_question_id": 5,
                        "ideal_question_name": "休日の曜日",
                        "ideal_answer_name": "土日,平日"
                    },
                    {
                        "ideal_question_id": 6,
                        "ideal_question_name": "休日の過ごし方",
                        "ideal_answer_name": "インドア派,アウトドア派"
                    },
                    {
                        "ideal_question_id": 7,
                        "ideal_question_name": "身長",
                        "ideal_answer_name": "185cm,186cm,187cm"
                    },
                    {
                        "ideal_question_id": 8,
                        "ideal_question_name": "体型",
                        "ideal_answer_name": "普通"
                    },
                    {
                        "ideal_question_id": 9,
                        "ideal_question_name": "お酒",
                        "ideal_answer_name": "たまに飲む"
                    },
                    {
                        "ideal_question_id": 10,
                        "ideal_question_name": "煙草",
                        "ideal_answer_name": "たまに吸う,吸わない"
                    },
                    {
                        "ideal_question_id": 11,
                        "ideal_question_name": "ギャンブル",
                        "ideal_answer_name": "しない"
                    },
                    {
                        "ideal_question_id": 12,
                        "ideal_question_name": "免許",
                        "ideal_answer_name": "車,バイク"
                    },
                    {
                        "ideal_question_id": 13,
                        "ideal_question_name": "一ヶ月の娯楽費",
                        "ideal_answer_name": "10,000円〜20,000円,20,000円〜30,000円"
                    },
                    {
                        "ideal_question_id": 14,
                        "ideal_question_name": "会話",
                        "ideal_answer_name": "話すのも聞くのも好き"
                    },
                    {
                        "ideal_question_id": 15,
                        "ideal_question_name": "出会うまでの希望",
                        "ideal_answer_name": "お互いの気が合えば,とりあえず会って話したい"
                    }
                ]
            },
            {
                "matching_format_name": "恋愛",
                "user_ideal_questions_and_answers": [
                    {
                        "ideal_question_id": 16,
                        "ideal_question_name": "交際経験",
                        "ideal_answer_name": "あり"
                    },
                    {
                        "ideal_question_id": 17,
                        "ideal_question_name": "メッセージ交換の頻度",
                        "ideal_answer_name": "週2〜3回,週4〜6回"
                    },
                    {
                        "ideal_question_id": 18,
                        "ideal_question_name": "デートの頻度",
                        "ideal_answer_name": "2〜3日に一回,4〜6日に一回,1週間に一回"
                    },
                    {
                        "ideal_question_id": 19,
                        "ideal_question_name": "初回デートの費用",
                        "ideal_answer_name": "男性が多めに払う,割り勘"
                    },
                    {
                        "ideal_question_id": 20,
                        "ideal_question_name": "同棲の希望",
                        "ideal_answer_name": "交際してから考える"
                    }
                ]
            },
            {
                "matching_format_name": "婚活",
                "user_ideal_questions_and_answers": [
                    {
                        "ideal_question_id": 21,
                        "ideal_question_name": "結婚経験",
                        "ideal_answer_name": "なし"
                    },
                    {
                        "ideal_question_id": 22,
                        "ideal_question_name": "家事",
                        "ideal_answer_name": "積極的に参加したい,余裕があれば参加したい"
                    },
                    {
                        "ideal_question_id": 23,
                        "ideal_question_name": "育児",
                        "ideal_answer_name": "積極的に参加したい"
                    },
                    {
                        "ideal_question_id": 24,
                        "ideal_question_name": "年収",
                        "ideal_answer_name": "200万円〜400万円,400万円〜600万円,600万円〜800万円"
                    },
                    {
                        "ideal_question_id": 25,
                        "ideal_question_name": "結婚式",
                        "ideal_answer_name": "現在こだわりはない"
                    },
                    {
                        "ideal_question_id": 26,
                        "ideal_question_name": "子供の有無",
                        "ideal_answer_name": "いない"
                    },
                    {
                        "ideal_question_id": 27,
                        "ideal_question_name": "子供はいつ欲しい？",
                        "ideal_answer_name": "1年目,2年目,3年目"
                    },
                    {
                        "ideal_question_id": 28,
                        "ideal_question_name": "子供は何人欲しい？",
                        "ideal_answer_name": "2人,3人以上"
                    },
                    {
                        "ideal_question_id": 29,
                        "ideal_question_name": "親との同居",
                        "ideal_answer_name": "したいなら考える"
                    }
                ]
            },
            {
                "matching_format_name": "ルームメイト",
                "user_ideal_questions_and_answers": [
                    {
                        "ideal_question_id": 30,
                        "ideal_question_name": "一番の目的",
                        "ideal_answer_name": "金銭面の負担を減らしたい"
                    },
                    {
                        "ideal_question_id": 31,
                        "ideal_question_name": "期間",
                        "ideal_answer_name": "半年以下,半年〜1年,1年〜2年"
                    },
                    {
                        "ideal_question_id": 32,
                        "ideal_question_name": "部屋の貸し借り",
                        "ideal_answer_name": "自分の部屋を貸したい,新しい物件を二人で借りたい"
                    },
                    {
                        "ideal_question_id": 33,
                        "ideal_question_name": "ペット",
                        "ideal_answer_name": "禁止"
                    },
                    {
                        "ideal_question_id": 34,
                        "ideal_question_name": "来客",
                        "ideal_answer_name": "歓迎,態度や回数など節度を守れるなら"
                    }
                ]
            }
        ]
    }
    ```
    </code>
    </pre>
    </details>


