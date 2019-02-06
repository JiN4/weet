WEET Back-end
====

## Description
マッチングアプリ-WEETのAPIサーバー
- 仕様言語：Golang
- フレームワーク：Gin
- ORM：Gorm
- DateBase:Mysql

## Install
- 追加予定

## APIドキュメント

### ユーザ情報をIDで絞り込んで取得する[GET]

#### エンドポイント
```
GET http://localhost:8080/api/v1/user/{user_id}
```

#### Request Parameters
場所|パラメータ名|指定する値|説明|
|:-|:-|:-|:-|
|Path|user_id|数値|取得したいユーザーのID|

#### Example Response
+ Response 200 (application/json)

    <details>
    <summary>Body</summary>
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
    </code>
    </pre>
    </details>

---


### ユーザ情報をIDで絞り込んで取得する(AnswerNameが配列)[GET]

#### エンドポイント
```
GET http://localhost:8080/api/v2/user/{user_id}
```

#### Request Parameters
|場所|パラメータ名|指定する値|説明|
|:-|:-|:-|:-|
|Path|user_id|数値|取得したいユーザーのID|

#### Example Response
+ Response 200 (application/json)

    <details>
    <summary>Body</summary>
    <pre>
    <code>
    {
        "user_basics": {
            "matching_format_name": "基本情報",
            "user_id": 1,
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
                        "ideal_answer_name": [
                            "大阪府",
                            "兵庫県"
                        ]
                    },
                    {
                        "ideal_question_id": 2,
                        "ideal_question_name": "血液型",
                        "ideal_answer_name": [
                            "B型",
                            "O型"
                        ]
                    },
                    {
                        "ideal_question_id": 3,
                        "ideal_question_name": "学歴",
                        "ideal_answer_name": [
                            "短大/専門学校卒",
                            "大学卒",
                            "大学院卒"
                        ]
                    },
                    {
                        "ideal_question_id": 4,
                        "ideal_question_name": "職種",
                        "ideal_answer_name": [
                            "オフィス系",
                            "研究・開発系",
                            "IT・エンジニア系",
                            "医療系"
                        ]
                    },
                    {
                        "ideal_question_id": 5,
                        "ideal_question_name": "休日の曜日",
                        "ideal_answer_name": [
                            "土日",
                            "平日"
                        ]
                    },
                    {
                        "ideal_question_id": 6,
                        "ideal_question_name": "休日の過ごし方",
                        "ideal_answer_name": [
                            "インドア派",
                            "アウトドア派"
                        ]
                    },
                    {
                        "ideal_question_id": 7,
                        "ideal_question_name": "身長",
                        "ideal_answer_name": [
                            "185cm",
                            "186cm",
                            "187cm"
                        ]
                    },
                    {
                        "ideal_question_id": 8,
                        "ideal_question_name": "体型",
                        "ideal_answer_name": [
                            "普通"
                        ]
                    },
                    {
                        "ideal_question_id": 9,
                        "ideal_question_name": "お酒",
                        "ideal_answer_name": [
                            "たまに飲む"
                        ]
                    },
                    {
                        "ideal_question_id": 10,
                        "ideal_question_name": "煙草",
                        "ideal_answer_name": [
                            "たまに吸う",
                            "吸わない"
                        ]
                    },
                    {
                        "ideal_question_id": 11,
                        "ideal_question_name": "ギャンブル",
                        "ideal_answer_name": [
                            "しない"
                        ]
                    },
                    {
                        "ideal_question_id": 12,
                        "ideal_question_name": "免許",
                        "ideal_answer_name": [
                            "車",
                            "バイク"
                        ]
                    },
                    {
                        "ideal_question_id": 13,
                        "ideal_question_name": "一ヶ月の娯楽費",
                        "ideal_answer_name": [
                            "10,000円〜20,000円",
                            "20,000円〜30,000円"
                        ]
                    },
                    {
                        "ideal_question_id": 14,
                        "ideal_question_name": "会話",
                        "ideal_answer_name": [
                            "話すのも聞くのも好き"
                        ]
                    },
                    {
                        "ideal_question_id": 15,
                        "ideal_question_name": "出会うまでの希望",
                        "ideal_answer_name": [
                            "お互いの気が合えば",
                            "とりあえず会って話したい"
                        ]
                    }
                ]
            },
            {
                "matching_format_name": "恋愛",
                "user_ideal_questions_and_answers": [
                    {
                        "ideal_question_id": 16,
                        "ideal_question_name": "交際経験",
                        "ideal_answer_name": [
                            "あり"
                        ]
                    },
                    {
                        "ideal_question_id": 17,
                        "ideal_question_name": "メッセージ交換の頻度",
                        "ideal_answer_name": [
                            "週2〜3回",
                            "週4〜6回"
                        ]
                    },
                    {
                        "ideal_question_id": 18,
                        "ideal_question_name": "デートの頻度",
                        "ideal_answer_name": [
                            "2〜3日に一回",
                            "4〜6日に一回",
                            "1週間に一回"
                        ]
                    },
                    {
                        "ideal_question_id": 19,
                        "ideal_question_name": "初回デートの費用",
                        "ideal_answer_name": [
                            "男性が多めに払う",
                            "割り勘"
                        ]
                    },
                    {
                        "ideal_question_id": 20,
                        "ideal_question_name": "同棲の希望",
                        "ideal_answer_name": [
                            "交際してから考える"
                        ]
                    }
                ]
            },
            {
                "matching_format_name": "婚活",
                "user_ideal_questions_and_answers": [
                    {
                        "ideal_question_id": 21,
                        "ideal_question_name": "結婚経験",
                        "ideal_answer_name": [
                            "なし"
                        ]
                    },
                    {
                        "ideal_question_id": 22,
                        "ideal_question_name": "家事",
                        "ideal_answer_name": [
                            "積極的に参加したい",
                            "余裕があれば参加したい"
                        ]
                    },
                    {
                        "ideal_question_id": 23,
                        "ideal_question_name": "育児",
                        "ideal_answer_name": [
                            "積極的に参加したい"
                        ]
                    },
                    {
                        "ideal_question_id": 24,
                        "ideal_question_name": "年収",
                        "ideal_answer_name": [
                            "200万円〜400万円",
                            "400万円〜600万円",
                            "600万円〜800万円"
                        ]
                    },
                    {
                        "ideal_question_id": 25,
                        "ideal_question_name": "結婚式",
                        "ideal_answer_name": [
                            "現在こだわりはない"
                        ]
                    },
                    {
                        "ideal_question_id": 26,
                        "ideal_question_name": "子供の有無",
                        "ideal_answer_name": [
                            "いない"
                        ]
                    },
                    {
                        "ideal_question_id": 27,
                        "ideal_question_name": "子供はいつ欲しい？",
                        "ideal_answer_name": [
                            "1年目",
                            "2年目",
                            "3年目"
                        ]
                    },
                    {
                        "ideal_question_id": 28,
                        "ideal_question_name": "子供は何人欲しい？",
                        "ideal_answer_name": [
                            "2人",
                            "3人以上"
                        ]
                    },
                    {
                        "ideal_question_id": 29,
                        "ideal_question_name": "親との同居",
                        "ideal_answer_name": [
                            "したいなら考える"
                        ]
                    }
                ]
            },
            {
                "matching_format_name": "ルームメイト",
                "user_ideal_questions_and_answers": [
                    {
                        "ideal_question_id": 30,
                        "ideal_question_name": "一番の目的",
                        "ideal_answer_name": [
                            "金銭面の負担を減らしたい"
                        ]
                    },
                    {
                        "ideal_question_id": 31,
                        "ideal_question_name": "期間",
                        "ideal_answer_name": [
                            "半年以下",
                            "半年〜1年",
                            "1年〜2年"
                        ]
                    },
                    {
                        "ideal_question_id": 32,
                        "ideal_question_name": "部屋の貸し借り",
                        "ideal_answer_name": [
                            "自分の部屋を貸したい",
                            "新しい物件を二人で借りたい"
                        ]
                    },
                    {
                        "ideal_question_id": 33,
                        "ideal_question_name": "ペット",
                        "ideal_answer_name": [
                            "禁止"
                        ]
                    },
                    {
                        "ideal_question_id": 34,
                        "ideal_question_name": "来客",
                        "ideal_answer_name": [
                            "歓迎",
                            "態度や回数など節度を守れるなら"
                        ]
                    }
                ]
            }
        ]
    }
    </code>
    </pre>
    </details>

---

### ユーザページの質疑応答の答えを全て取得する[GET]

#### エンドポイント
```
GET http://localhost:8080/api/v1/answers
```

#### Example Response

+ Response 200 (application/json)

    <details>
    <summary>Body</summary>
    <pre>
    <code>
    [
        {
            "question_id": 1,
            "candidate_answer": [
                {
                    "answer_id": 1,
                    "answer_name": "北海道"
                },
                {
                    "answer_id": 2,
                    "answer_name": "青森県"
                },
                {
                    "answer_id": 3,
                    "answer_name": "岩手県"
                },
                {
                    "answer_id": 4,
                    "answer_name": "宮城県"
                },
                {
                    "answer_id": 5,
                    "answer_name": "秋田県"
                },
                {
                    "answer_id": 6,
                    "answer_name": "山形県"
                },
                {
                    "answer_id": 7,
                    "answer_name": "福島県"
                },
                {
                    "answer_id": 8,
                    "answer_name": "茨城県"
                },
                {
                    "answer_id": 9,
                    "answer_name": "栃木県"
                },
                {
                    "answer_id": 10,
                    "answer_name": "群馬県"
                },
                {
                    "answer_id": 11,
                    "answer_name": "埼玉県"
                },
                {
                    "answer_id": 12,
                    "answer_name": "千葉県"
                },
                {
                    "answer_id": 13,
                    "answer_name": "東京都"
                },
                {
                    "answer_id": 14,
                    "answer_name": "神奈川県"
                },
                {
                    "answer_id": 15,
                    "answer_name": "新潟県"
                },
                {
                    "answer_id": 16,
                    "answer_name": "富山県"
                },
                {
                    "answer_id": 17,
                    "answer_name": "石川県"
                },
                {
                    "answer_id": 18,
                    "answer_name": "福井県"
                },
                {
                    "answer_id": 19,
                    "answer_name": "山梨県"
                },
                {
                    "answer_id": 20,
                    "answer_name": "長野県"
                },
                {
                    "answer_id": 21,
                    "answer_name": "岐阜県"
                },
                {
                    "answer_id": 22,
                    "answer_name": "静岡県"
                },
                {
                    "answer_id": 23,
                    "answer_name": "愛知県"
                },
                {
                    "answer_id": 24,
                    "answer_name": "三重県"
                },
                {
                    "answer_id": 25,
                    "answer_name": "滋賀県"
                },
                {
                    "answer_id": 26,
                    "answer_name": "京都府"
                },
                {
                    "answer_id": 27,
                    "answer_name": "大阪府"
                },
                {
                    "answer_id": 28,
                    "answer_name": "兵庫県"
                },
                {
                    "answer_id": 29,
                    "answer_name": "奈良県"
                },
                {
                    "answer_id": 30,
                    "answer_name": "和歌山県"
                },
                {
                    "answer_id": 31,
                    "answer_name": "鳥取県"
                },
                {
                    "answer_id": 32,
                    "answer_name": "島根県"
                },
                {
                    "answer_id": 33,
                    "answer_name": "岡山県"
                },
                {
                    "answer_id": 34,
                    "answer_name": "広島県"
                },
                {
                    "answer_id": 35,
                    "answer_name": "山口県"
                },
                {
                    "answer_id": 36,
                    "answer_name": "徳島県"
                },
                {
                    "answer_id": 37,
                    "answer_name": "香川県"
                },
                {
                    "answer_id": 38,
                    "answer_name": "愛媛県"
                },
                {
                    "answer_id": 39,
                    "answer_name": "高知県"
                },
                {
                    "answer_id": 40,
                    "answer_name": "福岡県"
                },
                {
                    "answer_id": 41,
                    "answer_name": "佐賀県"
                },
                {
                    "answer_id": 42,
                    "answer_name": "長崎県"
                },
                {
                    "answer_id": 43,
                    "answer_name": "熊本県"
                },
                {
                    "answer_id": 44,
                    "answer_name": "大分県"
                },
                {
                    "answer_id": 45,
                    "answer_name": "宮崎県"
                },
                {
                    "answer_id": 46,
                    "answer_name": "鹿児島県"
                },
                {
                    "answer_id": 47,
                    "answer_name": "沖縄県"
                }
            ]
        },
        {
            "question_id": 2,
            "candidate_answer": [
                {
                    "answer_id": 48,
                    "answer_name": "A型"
                },
                {
                    "answer_id": 49,
                    "answer_name": "B型"
                },
                {
                    "answer_id": 50,
                    "answer_name": "O型"
                },
                {
                    "answer_id": 51,
                    "answer_name": "AB型"
                }
            ]
        },
        {
            "question_id": 3,
            "candidate_answer": [
                {
                    "answer_id": 52,
                    "answer_name": "高校卒"
                },
                {
                    "answer_id": 53,
                    "answer_name": "短大/専門学校卒"
                },
                {
                    "answer_id": 54,
                    "answer_name": "大学卒"
                },
                {
                    "answer_id": 55,
                    "answer_name": "大学院卒"
                },
                {
                    "answer_id": 56,
                    "answer_name": "その他"
                }
            ]
        },
        {
            "question_id": 4,
            "candidate_answer": [
                {
                    "answer_id": 57,
                    "answer_name": "オフィス系"
                },
                {
                    "answer_id": 58,
                    "answer_name": "語学系"
                },
                {
                    "answer_id": 59,
                    "answer_name": "製造・物流・作業系"
                },
                {
                    "answer_id": 60,
                    "answer_name": "研究・開発系"
                },
                {
                    "answer_id": 61,
                    "answer_name": "IT・エンジニア系"
                },
                {
                    "answer_id": 62,
                    "answer_name": "接客・販売系"
                },
                {
                    "answer_id": 63,
                    "answer_name": "コールセンター系"
                },
                {
                    "answer_id": 64,
                    "answer_name": "営業系"
                },
                {
                    "answer_id": 65,
                    "answer_name": "金融系"
                },
                {
                    "answer_id": 66,
                    "answer_name": "医療系"
                },
                {
                    "answer_id": 67,
                    "answer_name": "その他"
                }
            ]
        },
        {
            "question_id": 5,
            "candidate_answer": [
                {
                    "answer_id": 68,
                    "answer_name": "土日"
                },
                {
                    "answer_id": 69,
                    "answer_name": "平日"
                },
                {
                    "answer_id": 70,
                    "answer_name": "不定期"
                }
            ]
        },
        {
            "question_id": 6,
            "candidate_answer": [
                {
                    "answer_id": 71,
                    "answer_name": "インドア派"
                },
                {
                    "answer_id": 72,
                    "answer_name": "アウトドア派"
                }
            ]
        },
        {
            "question_id": 7,
            "candidate_answer": [
                {
                    "answer_id": 73,
                    "answer_name": "140cm以下"
                },
                {
                    "answer_id": 74,
                    "answer_name": "141cm"
                },
                {
                    "answer_id": 75,
                    "answer_name": "142cm"
                },
                {
                    "answer_id": 76,
                    "answer_name": "143cm"
                },
                {
                    "answer_id": 77,
                    "answer_name": "144cm"
                },
                {
                    "answer_id": 78,
                    "answer_name": "145cm"
                },
                {
                    "answer_id": 79,
                    "answer_name": "146cm"
                },
                {
                    "answer_id": 80,
                    "answer_name": "147cm"
                },
                {
                    "answer_id": 81,
                    "answer_name": "148cm"
                },
                {
                    "answer_id": 82,
                    "answer_name": "149cm"
                },
                {
                    "answer_id": 83,
                    "answer_name": "150cm"
                },
                {
                    "answer_id": 84,
                    "answer_name": "151cm"
                },
                {
                    "answer_id": 85,
                    "answer_name": "152cm"
                },
                {
                    "answer_id": 86,
                    "answer_name": "153cm"
                },
                {
                    "answer_id": 87,
                    "answer_name": "154cm"
                },
                {
                    "answer_id": 88,
                    "answer_name": "155cm"
                },
                {
                    "answer_id": 89,
                    "answer_name": "156cm"
                },
                {
                    "answer_id": 90,
                    "answer_name": "157cm"
                },
                {
                    "answer_id": 91,
                    "answer_name": "158"
                },
                {
                    "answer_id": 92,
                    "answer_name": "159cm"
                },
                {
                    "answer_id": 93,
                    "answer_name": "160cm"
                },
                {
                    "answer_id": 94,
                    "answer_name": "161cm"
                },
                {
                    "answer_id": 95,
                    "answer_name": "162cm"
                },
                {
                    "answer_id": 96,
                    "answer_name": "163cm"
                },
                {
                    "answer_id": 97,
                    "answer_name": "164cm"
                },
                {
                    "answer_id": 98,
                    "answer_name": "165cm"
                },
                {
                    "answer_id": 99,
                    "answer_name": "166cm"
                },
                {
                    "answer_id": 100,
                    "answer_name": "167cm"
                },
                {
                    "answer_id": 101,
                    "answer_name": "168cm"
                },
                {
                    "answer_id": 102,
                    "answer_name": "169cm"
                },
                {
                    "answer_id": 103,
                    "answer_name": "170cm"
                },
                {
                    "answer_id": 104,
                    "answer_name": "171cm"
                },
                {
                    "answer_id": 105,
                    "answer_name": "172cm"
                },
                {
                    "answer_id": 106,
                    "answer_name": "173cm"
                },
                {
                    "answer_id": 107,
                    "answer_name": "174cm"
                },
                {
                    "answer_id": 108,
                    "answer_name": "175cm"
                },
                {
                    "answer_id": 109,
                    "answer_name": "176cm"
                },
                {
                    "answer_id": 110,
                    "answer_name": "177cm"
                },
                {
                    "answer_id": 111,
                    "answer_name": "178cm"
                },
                {
                    "answer_id": 112,
                    "answer_name": "179cm"
                },
                {
                    "answer_id": 113,
                    "answer_name": "180cm"
                },
                {
                    "answer_id": 114,
                    "answer_name": "181cm"
                },
                {
                    "answer_id": 115,
                    "answer_name": "182cm"
                },
                {
                    "answer_id": 116,
                    "answer_name": "183cm"
                },
                {
                    "answer_id": 117,
                    "answer_name": "184cm"
                },
                {
                    "answer_id": 118,
                    "answer_name": "185cm"
                },
                {
                    "answer_id": 119,
                    "answer_name": "186cm"
                },
                {
                    "answer_id": 120,
                    "answer_name": "187cm"
                },
                {
                    "answer_id": 121,
                    "answer_name": "188cm"
                },
                {
                    "answer_id": 122,
                    "answer_name": "189cm"
                },
                {
                    "answer_id": 123,
                    "answer_name": "190cm以上"
                }
            ]
        },
        {
            "question_id": 8,
            "candidate_answer": [
                {
                    "answer_id": 124,
                    "answer_name": "細め"
                },
                {
                    "answer_id": 125,
                    "answer_name": "スレンダー"
                },
                {
                    "answer_id": 126,
                    "answer_name": "普通"
                },
                {
                    "answer_id": 127,
                    "answer_name": "ぽっちゃり"
                },
                {
                    "answer_id": 128,
                    "answer_name": "グラマー"
                },
                {
                    "answer_id": 129,
                    "answer_name": "筋肉質"
                }
            ]
        },
        {
            "question_id": 9,
            "candidate_answer": [
                {
                    "answer_id": 130,
                    "answer_name": "飲む"
                },
                {
                    "answer_id": 131,
                    "answer_name": "たまに飲む"
                },
                {
                    "answer_id": 132,
                    "answer_name": "あまり飲まない"
                },
                {
                    "answer_id": 133,
                    "answer_name": "飲まない"
                }
            ]
        },
        {
            "question_id": 10,
            "candidate_answer": [
                {
                    "answer_id": 134,
                    "answer_name": "吸う"
                },
                {
                    "answer_id": 135,
                    "answer_name": "たまに吸う"
                },
                {
                    "answer_id": 136,
                    "answer_name": "吸わない"
                }
            ]
        },
        {
            "question_id": 11,
            "candidate_answer": [
                {
                    "answer_id": 137,
                    "answer_name": "する"
                },
                {
                    "answer_id": 138,
                    "answer_name": "たまにする"
                },
                {
                    "answer_id": 139,
                    "answer_name": "しない"
                }
            ]
        },
        {
            "question_id": 12,
            "candidate_answer": [
                {
                    "answer_id": 140,
                    "answer_name": "車"
                },
                {
                    "answer_id": 141,
                    "answer_name": "バイク"
                },
                {
                    "answer_id": 142,
                    "answer_name": "車・バイク"
                },
                {
                    "answer_id": 143,
                    "answer_name": "ない"
                }
            ]
        },
        {
            "question_id": 13,
            "candidate_answer": [
                {
                    "answer_id": 144,
                    "answer_name": "10,000円以下"
                },
                {
                    "answer_id": 145,
                    "answer_name": "10,000円〜20,000円"
                },
                {
                    "answer_id": 146,
                    "answer_name": "20,000円〜30,000円"
                },
                {
                    "answer_id": 147,
                    "answer_name": "30,000円〜40,000円"
                },
                {
                    "answer_id": 148,
                    "answer_name": "40,000円〜50,000円"
                },
                {
                    "answer_id": 149,
                    "answer_name": "50,000円以上"
                }
            ]
        },
        {
            "question_id": 14,
            "candidate_answer": [
                {
                    "answer_id": 150,
                    "answer_name": "話す方が好き"
                },
                {
                    "answer_id": 151,
                    "answer_name": "聞く方が好き"
                },
                {
                    "answer_id": 152,
                    "answer_name": "話すのも聞くのも好き"
                },
                {
                    "answer_id": 153,
                    "answer_name": "会話は得意ではない"
                }
            ]
        },
        {
            "question_id": 15,
            "candidate_answer": [
                {
                    "answer_id": 154,
                    "answer_name": "ある程度のメッセージ交換を重ねてから"
                },
                {
                    "answer_id": 155,
                    "answer_name": "お互いの気が合えば"
                },
                {
                    "answer_id": 156,
                    "answer_name": "とりあえず会って話したい"
                }
            ]
        },
        {
            "question_id": 16,
            "candidate_answer": [
                {
                    "answer_id": 157,
                    "answer_name": "あり"
                },
                {
                    "answer_id": 158,
                    "answer_name": "なし"
                }
            ]
        },
        {
            "question_id": 17,
            "candidate_answer": [
                {
                    "answer_id": 159,
                    "answer_name": "毎日"
                },
                {
                    "answer_id": 160,
                    "answer_name": "週2〜3回"
                },
                {
                    "answer_id": 161,
                    "answer_name": "週4〜6回"
                },
                {
                    "answer_id": 162,
                    "answer_name": "その他"
                }
            ]
        },
        {
            "question_id": 18,
            "candidate_answer": [
                {
                    "answer_id": 163,
                    "answer_name": "2〜3日に一回"
                },
                {
                    "answer_id": 164,
                    "answer_name": "4〜6日に一回"
                },
                {
                    "answer_id": 165,
                    "answer_name": "1週間に一回"
                },
                {
                    "answer_id": 166,
                    "answer_name": "2週間に一回"
                },
                {
                    "answer_id": 167,
                    "answer_name": "その他"
                }
            ]
        },
        {
            "question_id": 19,
            "candidate_answer": [
                {
                    "answer_id": 168,
                    "answer_name": "男性が全額払う"
                },
                {
                    "answer_id": 169,
                    "answer_name": "男性が多めに払う"
                },
                {
                    "answer_id": 170,
                    "answer_name": "割り勘"
                },
                {
                    "answer_id": 171,
                    "answer_name": "持っている方が払う"
                },
                {
                    "answer_id": 172,
                    "answer_name": "場所のことも考え、話し合って決める"
                }
            ]
        },
        {
            "question_id": 20,
            "candidate_answer": [
                {
                    "answer_id": 173,
                    "answer_name": "したい"
                },
                {
                    "answer_id": 174,
                    "answer_name": "したくない"
                },
                {
                    "answer_id": 175,
                    "answer_name": "交際してから考える"
                }
            ]
        },
        {
            "question_id": 21,
            "candidate_answer": [
                {
                    "answer_id": 176,
                    "answer_name": "あり"
                },
                {
                    "answer_id": 177,
                    "answer_name": "なし"
                }
            ]
        },
        {
            "question_id": 22,
            "candidate_answer": [
                {
                    "answer_id": 178,
                    "answer_name": "積極的に参加したい"
                },
                {
                    "answer_id": 179,
                    "answer_name": "余裕があれば参加したい"
                },
                {
                    "answer_id": 180,
                    "answer_name": "状況を見て参加したい"
                },
                {
                    "answer_id": 181,
                    "answer_name": "相手に任せたい"
                }
            ]
        },
        {
            "question_id": 23,
            "candidate_answer": [
                {
                    "answer_id": 182,
                    "answer_name": "積極的に参加したい"
                },
                {
                    "answer_id": 183,
                    "answer_name": "余裕があれば参加したい"
                },
                {
                    "answer_id": 184,
                    "answer_name": "状況を見て参加したい"
                },
                {
                    "answer_id": 185,
                    "answer_name": "相手に任せたい"
                }
            ]
        },
        {
            "question_id": 24,
            "candidate_answer": [
                {
                    "answer_id": 186,
                    "answer_name": "200万円未満"
                },
                {
                    "answer_id": 187,
                    "answer_name": "200万円〜400万円"
                },
                {
                    "answer_id": 188,
                    "answer_name": "400万円〜600万円"
                },
                {
                    "answer_id": 189,
                    "answer_name": "600万円〜800万円"
                },
                {
                    "answer_id": 190,
                    "answer_name": "800万円〜1000万円"
                },
                {
                    "answer_id": 191,
                    "answer_name": "1,000万円以上"
                }
            ]
        },
        {
            "question_id": 25,
            "candidate_answer": [
                {
                    "answer_id": 192,
                    "answer_name": "教会式"
                },
                {
                    "answer_id": 193,
                    "answer_name": "キリスト教式"
                },
                {
                    "answer_id": 194,
                    "answer_name": "神前式"
                },
                {
                    "answer_id": 195,
                    "answer_name": "仏前式"
                },
                {
                    "answer_id": 196,
                    "answer_name": "人前式"
                },
                {
                    "answer_id": 197,
                    "answer_name": "現在こだわりはない"
                },
                {
                    "answer_id": 198,
                    "answer_name": "しない"
                }
            ]
        },
        {
            "question_id": 26,
            "candidate_answer": [
                {
                    "answer_id": 199,
                    "answer_name": "いる"
                },
                {
                    "answer_id": 200,
                    "answer_name": "いない"
                }
            ]
        },
        {
            "question_id": 27,
            "candidate_answer": [
                {
                    "answer_id": 201,
                    "answer_name": "1年未満"
                },
                {
                    "answer_id": 202,
                    "answer_name": "1年目"
                },
                {
                    "answer_id": 203,
                    "answer_name": "2年目"
                },
                {
                    "answer_id": 204,
                    "answer_name": "3年目"
                },
                {
                    "answer_id": 205,
                    "answer_name": "3年目以降"
                },
                {
                    "answer_id": 206,
                    "answer_name": "いらない"
                }
            ]
        },
        {
            "question_id": 28,
            "candidate_answer": [
                {
                    "answer_id": 207,
                    "answer_name": "1人"
                },
                {
                    "answer_id": 208,
                    "answer_name": "2人"
                },
                {
                    "answer_id": 209,
                    "answer_name": "3人以上"
                },
                {
                    "answer_id": 210,
                    "answer_name": "いらない"
                }
            ]
        },
        {
            "question_id": 29,
            "candidate_answer": [
                {
                    "answer_id": 211,
                    "answer_name": "あり"
                },
                {
                    "answer_id": 212,
                    "answer_name": "したいなら考える"
                },
                {
                    "answer_id": 213,
                    "answer_name": "なし"
                }
            ]
        },
        {
            "question_id": 30,
            "candidate_answer": [
                {
                    "answer_id": 214,
                    "answer_name": "金銭面の負担を減らしたい"
                },
                {
                    "answer_id": 215,
                    "answer_name": "家事の負担を減らしたい"
                },
                {
                    "answer_id": 216,
                    "answer_name": "誰かがいる安心感を得たいから"
                },
                {
                    "answer_id": 217,
                    "answer_name": "相談・愚痴が言える人と住みたい"
                },
                {
                    "answer_id": 218,
                    "answer_name": "きまぐれ"
                },
                {
                    "answer_id": 219,
                    "answer_name": "その他"
                }
            ]
        },
        {
            "question_id": 31,
            "candidate_answer": [
                {
                    "answer_id": 220,
                    "answer_name": "半年以下"
                },
                {
                    "answer_id": 221,
                    "answer_name": "半年〜1年"
                },
                {
                    "answer_id": 222,
                    "answer_name": "1年〜2年"
                },
                {
                    "answer_id": 223,
                    "answer_name": "2年以上"
                },
                {
                    "answer_id": 224,
                    "answer_name": "未定"
                }
            ]
        },
        {
            "question_id": 32,
            "candidate_answer": [
                {
                    "answer_id": 225,
                    "answer_name": "自分の部屋を貸したい"
                },
                {
                    "answer_id": 226,
                    "answer_name": "誰かの部屋を借りたい"
                },
                {
                    "answer_id": 227,
                    "answer_name": "新しい物件を二人で借りたい"
                },
                {
                    "answer_id": 228,
                    "answer_name": "その他"
                }
            ]
        },
        {
            "question_id": 33,
            "candidate_answer": [
                {
                    "answer_id": 229,
                    "answer_name": "大丈夫"
                },
                {
                    "answer_id": 230,
                    "answer_name": "躾されているなら大丈夫"
                },
                {
                    "answer_id": 231,
                    "answer_name": "禁止"
                }
            ]
        },
        {
            "question_id": 34,
            "candidate_answer": [
                {
                    "answer_id": 232,
                    "answer_name": "歓迎"
                },
                {
                    "answer_id": 233,
                    "answer_name": "態度や回数など節度を守れるなら"
                },
                {
                    "answer_id": 234,
                    "answer_name": "あまり好まない"
                },
                {
                    "answer_id": 235,
                    "answer_name": "禁止"
                }
            ]
        },
        {
            "question_id": 1,
            "candidate_answer": [
                {
                    "answer_id": 236,
                    "answer_name": "-"
                }
            ]
        },
        {
            "question_id": 2,
            "candidate_answer": [
                {
                    "answer_id": 237,
                    "answer_name": "-"
                }
            ]
        },
        {
            "question_id": 3,
            "candidate_answer": [
                {
                    "answer_id": 238,
                    "answer_name": "-"
                }
            ]
        },
        {
            "question_id": 4,
            "candidate_answer": [
                {
                    "answer_id": 239,
                    "answer_name": "-"
                }
            ]
        },
        {
            "question_id": 5,
            "candidate_answer": [
                {
                    "answer_id": 240,
                    "answer_name": "-"
                }
            ]
        },
        {
            "question_id": 6,
            "candidate_answer": [
                {
                    "answer_id": 241,
                    "answer_name": "-"
                }
            ]
        },
        {
            "question_id": 7,
            "candidate_answer": [
                {
                    "answer_id": 242,
                    "answer_name": "-"
                }
            ]
        },
        {
            "question_id": 8,
            "candidate_answer": [
                {
                    "answer_id": 243,
                    "answer_name": "-"
                }
            ]
        },
        {
            "question_id": 9,
            "candidate_answer": [
                {
                    "answer_id": 244,
                    "answer_name": "-"
                }
            ]
        },
        {
            "question_id": 10,
            "candidate_answer": [
                {
                    "answer_id": 245,
                    "answer_name": "-"
                }
            ]
        },
        {
            "question_id": 11,
            "candidate_answer": [
                {
                    "answer_id": 246,
                    "answer_name": "-"
                }
            ]
        },
        {
            "question_id": 12,
            "candidate_answer": [
                {
                    "answer_id": 247,
                    "answer_name": "-"
                }
            ]
        },
        {
            "question_id": 13,
            "candidate_answer": [
                {
                    "answer_id": 248,
                    "answer_name": "-"
                }
            ]
        },
        {
            "question_id": 14,
            "candidate_answer": [
                {
                    "answer_id": 249,
                    "answer_name": "-"
                }
            ]
        },
        {
            "question_id": 15,
            "candidate_answer": [
                {
                    "answer_id": 250,
                    "answer_name": "-"
                }
            ]
        },
        {
            "question_id": 16,
            "candidate_answer": [
                {
                    "answer_id": 251,
                    "answer_name": "-"
                }
            ]
        },
        {
            "question_id": 17,
            "candidate_answer": [
                {
                    "answer_id": 252,
                    "answer_name": "-"
                }
            ]
        },
        {
            "question_id": 18,
            "candidate_answer": [
                {
                    "answer_id": 253,
                    "answer_name": "-"
                }
            ]
        },
        {
            "question_id": 19,
            "candidate_answer": [
                {
                    "answer_id": 254,
                    "answer_name": "-"
                }
            ]
        },
        {
            "question_id": 20,
            "candidate_answer": [
                {
                    "answer_id": 255,
                    "answer_name": "-"
                }
            ]
        },
        {
            "question_id": 21,
            "candidate_answer": [
                {
                    "answer_id": 256,
                    "answer_name": "-"
                }
            ]
        },
        {
            "question_id": 22,
            "candidate_answer": [
                {
                    "answer_id": 257,
                    "answer_name": "-"
                }
            ]
        },
        {
            "question_id": 23,
            "candidate_answer": [
                {
                    "answer_id": 258,
                    "answer_name": "-"
                }
            ]
        },
        {
            "question_id": 24,
            "candidate_answer": [
                {
                    "answer_id": 259,
                    "answer_name": "-"
                }
            ]
        },
        {
            "question_id": 25,
            "candidate_answer": [
                {
                    "answer_id": 260,
                    "answer_name": "-"
                }
            ]
        },
        {
            "question_id": 26,
            "candidate_answer": [
                {
                    "answer_id": 261,
                    "answer_name": "-"
                }
            ]
        },
        {
            "question_id": 27,
            "candidate_answer": [
                {
                    "answer_id": 262,
                    "answer_name": "-"
                }
            ]
        },
        {
            "question_id": 28,
            "candidate_answer": [
                {
                    "answer_id": 263,
                    "answer_name": "-"
                }
            ]
        },
        {
            "question_id": 29,
            "candidate_answer": [
                {
                    "answer_id": 264,
                    "answer_name": "-"
                }
            ]
        },
        {
            "question_id": 30,
            "candidate_answer": [
                {
                    "answer_id": 265,
                    "answer_name": "-"
                }
            ]
        },
        {
            "question_id": 31,
            "candidate_answer": [
                {
                    "answer_id": 266,
                    "answer_name": "-"
                }
            ]
        },
        {
            "question_id": 32,
            "candidate_answer": [
                {
                    "answer_id": 267,
                    "answer_name": "-"
                }
            ]
        },
        {
            "question_id": 33,
            "candidate_answer": [
                {
                    "answer_id": 268,
                    "answer_name": "-"
                }
            ]
        },
        {
            "question_id": 34,
            "candidate_answer": [
                {
                    "answer_id": 269,
                    "answer_name": "-"
                }
            ]
        }
    ]
    </code>
    </pre>
    </details>
---

### マッチング相手を性別・居住地・年齢で絞り込みランダムで取得[GET]

#### エンドポイント
```
GET http://localhost:8080/api/v1/matching/player/{user_id}/matching-format/{matching_format_id}
```
#### Request Parameters
|場所|パラメータ名|指定する値|説明|
|:-|:-|:-|:-|
|Path|user_id|数値|プレイヤーID|
|Path|matching_format_id|数値|マッチング形式ID|

#### Example Response

+ Response 200 (application/json)

    <details>
    <summary>Body</summary>
    <pre>
    <code>
    {
        "user_basics": {
            "matching_format_name": "基本情報",
            "user_name": "七瀬 奏",
            "image1": "https://www.pakutaso.com/shared/img/thumb/YUKI151117060I9A1562_TP_V.jpg",
            "image2": "",
            "image3": "",
            "sex": "女性",
            "age": 21,
            "residence": "兵庫県",
            "hitokoto": "音楽が趣味です",
            "comment": "高校の時、バイオリンをしてました。\nそのときからコンサートにいくのが好きで、一緒に聴きにいける人と知り合えたらいいなと思ってます！"
        },
        "user_specials": [
            {
                "matching_format_name": "友達",
                "user_questions_and_answers": [
                    {
                        "question_id": 1,
                        "question_name": "出身地",
                        "answer_name": "千葉県"
                    },
                    {
                        "question_id": 2,
                        "question_name": "血液型",
                        "answer_name": "AB型"
                    },
                    {
                        "question_id": 3,
                        "question_name": "学歴",
                        "answer_name": "大学卒"
                    },
                    {
                        "question_id": 4,
                        "question_name": "職種",
                        "answer_name": "オフィス系"
                    },
                    {
                        "question_id": 5,
                        "question_name": "休日の曜日",
                        "answer_name": "145cm"
                    },
                    {
                        "question_id": 6,
                        "question_name": "休日の過ごし方",
                        "answer_name": "インドア派"
                    },
                    {
                        "question_id": 7,
                        "question_name": "身長",
                        "answer_name": "152cm"
                    },
                    {
                        "question_id": 8,
                        "question_name": "体型",
                        "answer_name": "スレンダー"
                    },
                    {
                        "question_id": 9,
                        "question_name": "お酒",
                        "answer_name": "あまり飲まない"
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
                        "answer_name": "30,000円〜40,000円"
                    },
                    {
                        "question_id": 14,
                        "question_name": "会話",
                        "answer_name": "会話は得意ではない"
                    },
                    {
                        "question_id": 15,
                        "question_name": "出会うまでの希望",
                        "answer_name": "お互いの気が合えば"
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
                        "ideal_answer_name": "神奈川県"
                    },
                    {
                        "ideal_question_id": 2,
                        "ideal_question_name": "血液型",
                        "ideal_answer_name": "O型"
                    },
                    {
                        "ideal_question_id": 3,
                        "ideal_question_name": "学歴",
                        "ideal_answer_name": "大学卒"
                    },
                    {
                        "ideal_question_id": 4,
                        "ideal_question_name": "職種",
                        "ideal_answer_name": "オフィス系"
                    },
                    {
                        "ideal_question_id": 5,
                        "ideal_question_name": "休日の曜日",
                        "ideal_answer_name": "145cm"
                    },
                    {
                        "ideal_question_id": 6,
                        "ideal_question_name": "休日の過ごし方",
                        "ideal_answer_name": "インドア派"
                    },
                    {
                        "ideal_question_id": 7,
                        "ideal_question_name": "身長",
                        "ideal_answer_name": "175cm"
                    },
                    {
                        "ideal_question_id": 8,
                        "ideal_question_name": "体型",
                        "ideal_answer_name": "普通"
                    },
                    {
                        "ideal_question_id": 9,
                        "ideal_question_name": "お酒",
                        "ideal_answer_name": "あまり飲まない"
                    },
                    {
                        "ideal_question_id": 10,
                        "ideal_question_name": "煙草",
                        "ideal_answer_name": "吸わない"
                    },
                    {
                        "ideal_question_id": 11,
                        "ideal_question_name": "ギャンブル",
                        "ideal_answer_name": "しない"
                    },
                    {
                        "ideal_question_id": 12,
                        "ideal_question_name": "免許",
                        "ideal_answer_name": "車"
                    },
                    {
                        "ideal_question_id": 13,
                        "ideal_question_name": "一ヶ月の娯楽費",
                        "ideal_answer_name": "20,000円〜30,000円"
                    },
                    {
                        "ideal_question_id": 14,
                        "ideal_question_name": "会話",
                        "ideal_answer_name": "話すのも聞くのも好き"
                    },
                    {
                        "ideal_question_id": 15,
                        "ideal_question_name": "出会うまでの希望",
                        "ideal_answer_name": "お互いの気が合えば"
                    }
                ]
            }
        ],
        "matching_questions": [
            3,
            4,
            6,
            10,
            11,
            15
        ],
        "matching_ideal_questions": [
            6,
            10,
            11,
            13,
            14,
            15
        ]
    }
    </code>
    </pre>
    </details>

---

### プレイヤーにいいね！を送ってきたユーザを全て取得する[GET]

#### エンドポイント
```
GET http://localhost:8080/api/v1/favo/user/{user_id}
```

#### Request Parameters
|場所|パラメータ名|指定する値|説明|
|:-|:-|:-|:-|
|Path|user_id|数値|プレイヤーID|

#### Example Response

+ Response 200 (application/json)

    <details>
    <summary>Body</summary>
    <pre>
    <code>
    {
        "friend_favo_users": [
            {
                "user_id": 9,
                "user_name": "小鳥遊 詩織",
                "image1": "https://www.pakutaso.com/shared/img/thumb/YUKA150912596015_TP_V.jpg",
                "sex": "女性",
                "age": 24,
                "residence": "京都府",
                "hitokoto": "本が好きです。"
            }
        ],
        "love_favo_users": [
            {
                "user_id": 7,
                "user_name": "立花 彩香",
                "image1": "https://www.pakutaso.com/shared/img/thumb/TSURU1891A041_TP_V.jpg",
                "sex": "女性",
                "age": 27,
                "residence": "兵庫県",
                "hitokoto": "誰かー！かまってくれー！"
            }
        ],
        "marriage_favo_users": [],
        "roommate_favo_users": []
    }
    </code>
    </pre>
    </details>

---

### プレイヤーとメッセージ可能なユーザを全て取得する[GET]

#### エンドポイント
```
GET http://localhost:8080/api/v1/mutual-favo/user/{user_id}
```

#### Request Parameters
|場所|パラメータ名|指定する値|説明|
|:-|:-|:-|:-|
|Path|user_id|数値|プレイヤーのID|

#### Example Response

+ Response 200 (application/json)

    <details>
    <summary>Body</summary>
    <pre>
    <code>
    {
        "friend_favo_users": [
            {
                "user_id": 5,
                "user_name": "球投 一球",
                "image1": "https://images.unsplash.com/photo-1489460427746-b6296f4bc3f5?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=634&q=80",
                "sex": "男性",
                "age": 26,
                "residence": "大阪府",
                "hitokoto": "ストライク王に俺はなる！"
            }
        ],
        "love_favo_users": [
            {
                "user_id": 6,
                "user_name": "七瀬 奏",
                "image1": "https://www.pakutaso.com/shared/img/thumb/YUKI151117060I9A1562_TP_V.jpg",
                "sex": "女性",
                "age": 21,
                "residence": "兵庫県",
                "hitokoto": "音楽が趣味です"
            }
        ],
        "marriage_favo_users": [],
        "roommate_favo_users": []
    }
    </code>
    </pre>
    </details>

---

### プレイヤーがいいねしたユーザを登録する[POST]

#### エンドポイント
```
POST http://localhost:8080/api/v1/favo/user
```

#### Request Parameters
|場所|パラメータ名|指定する値|説明|
|:-|:-|:-|:-|
|Body|MatchingFormatID|数値|マッチングしたマッチング形式ID|
|Body|PlayerUserID|数値|いいね！する側のユーザーID|
|Body|FavoUserID|数値|いいね！される側のユーザーID|

#### Example Request


```
curl http://localhost:8080/api/v1/favo/user -X POST -H "Content-Type: application/json" -d '{"MatchingFormatID": 1,"PlayerUserID": 1,"FavoUserID": 2}'
```
---

### プレイヤーの相互いいね！ユーザを登録する[POST]

#### エンドポイント
```
POST http://localhost:8080/api/v1/mutual-favo/user
```

#### Request Parameters
|場所|パラメータ名|指定する値|説明|
|:-|:-|:-|:-|
|Body|MatchingFormatID|数値|マッチングしたマッチング形式ID|
|Body|UserID1|数値|いいね！申請の受信ユーザーID|
|Body|UserID2|数値|いいね！申請の送信ユーザーID|

#### Example Request
```
curl http://localhost:8080/api/v1/mutual-favo/user -X POST -H "Content-Type: application/json" -d '{"MatchingFormatID": 1,"UserID1": 1,"UserID2": 2}'

```
---
### ユーザの基本情報を変更する[PUT]

#### エンドポイント
```
PUT http://54.238.92.95:8080/api/v1/user/:user_id/update/basics
```

#### Request Parameters
|場所|パラメータ名|指定する値|説明|
|:-|:-|:-|:-|
|Path|user_id|数値|変更したいユーザーのID|
|Body|Image1|文字列|新しいプロフィール画像１つ目のURL|
|Body|Image2|文字列|新しいプロフィール画像２つ目のURL|
|Body|Image3|文字列|新しいプロフィール画像３つ目のURL|
|Body|Hitokoto|文字列|新しい一言の文字列|
|Body|Comment|文字列|新しい自由記入欄の文字列|

#### Example Request
```
curl http://localhost:8080/api/v1/user/:user_id/update/basics -X PUT -H "Content-Type: application/json" -d '{"Hitokoto": "よろしくお願いします！"}'

```
---

### ユーザの質疑応答の自分の情報を変更する[PUT]

#### エンドポイント
```
PUT http://54.238.92.95:8080/api/v1/user/:user_id/update/specials
```

#### Request Parameters
|場所|パラメータ名|指定する値|説明|
|:-|:-|:-|:-|
|Path|user_id|数値|変更したいユーザーのID|
|Body|question_id|数値|更新したい質問のID|
|Body|answer_id|数値|新しい自分の回答|

#### Example Request
```
curl -w '\n' 'http://localhost:8080/api/v1/user/1/update/specials' --data 'question_id=1&answer_id=27' -XPUT
```
---
### ユーザの質疑応答の理想像を変更する[PUT]

#### エンドポイント
```
PUT http://localhost:8080/api/v1/user/:user_id/update/ideal-specials
```

#### Request Parameters
|場所|パラメータ名|指定する値|説明|
|:-|:-|:-|:-|
|Path|user_id|数値|変更したいユーザーのID|
|Body|question_id|文字列|更新したい質問のID|
|Body|answer_ids|文字列|新しい自分の回答群|

#### Example Request
```
curl -w '\n' 'http://localhost:8080/api/v1/user/1/update/ideal-specials' --data 'matching_format_id=1&question_id=1&answer_ids=3,4,47' -XPUT
```
---
### マッチング対象に絞り込む居住地を変更する[PUT]

#### エンドポイント
```
PUT http://localhost:8080/api/v1/matching-prefectures/:user_id
```

#### Request Parameters
|場所|パラメータ名|指定する値|説明|
|:-|:-|:-|:-|
|Path|user_id|数値|変更したいユーザーのID|
|Body|PrefecturesID|文字列|絞り込みたい居住地のID群（1:北海道から47:沖縄県まで連番）|

#### Example Request
```
curl http://localhost:8080/api/v1/matching-prefectures/1 -X PUT -H "Content-Type: application/json" -d '{"PrefecturesID": "1"}'

```
---
### マッチング対象に絞り込む性別を変更する[PUT]

#### エンドポイント
```
PUT http://localhost:8080/api/v1/matching-sexes/:user_id
```

#### Request Parameters
|場所|パラメータ名|指定する値|説明|
|:-|:-|:-|:-|
|Path|user_id|数値|変更したいユーザーのID|
|Body|SexID|数値|絞り込みたい性別の種類（1:男性 2:女性 3:両方）|

#### Example Request
```
curl http://localhost:8080/api/v1/matching-sexes/1 -X PUT -H "Content-Type: application/json" -d '{"SexID": 3}'

```
---
### マッチング対象に絞り込む年齢を変更する[PUT]

#### エンドポイント
```
PUT http://localhost:8080/api/v1/matching-ages/:user_id
```

#### Request Parameters
|場所|パラメータ名|指定する値|説明|
|:-|:-|:-|:-|
|Path|user_id|数値|変更したいユーザーのID|
|Body|FirstAge|数値|絞り込みたい年齢の下限の数値|
|Body|LastAge|数値|絞り込みたい年齢の上限の数値|

#### Example Request
```
curl http://localhost:8080/api/v1/matching-ages/1 -X PUT -H "Content-Type: application/json" -d '{"FirstAge": 20, "LastAge": 30}'

```
---
### 誰かのマッチングの対象にされる形式の種類を変更する[PUT]

#### エンドポイント
```
PUT http://localhost:8080/api/v1/matching-format-choices/:user_id
```

#### Request Parameters
|場所|パラメータ名|指定する値|説明|
|:-|:-|:-|:-|
|Path|user_id|数値|変更したいユーザーのID|
|Body|Love|数値|変更しない場合は現在、変更する場合は変更後の値（0:マッチング対象にならない/1:マッチング対象になる）|
|Body|Marriage|数値|変更しない場合は現在、変更する場合は変更後の値（0:マッチング対象にならない/1:マッチング対象になる）|
|Body|Roommate|数値|変更しない場合は現在、変更する場合は変更後の値（0:マッチング対象にならない/1:マッチング対象になる）|

#### Example Request
```
curl http://localhost:8080/api/v1/matching-sexes/1 -X PUT -H "Content-Type: application/json" -d '{"Love": 1,"Marriage": 0,"Roommate": 0}'
```
---
### プレイヤーをいいね！しているユーザを削除[DELETE]

#### エンドポイント
```
DELETE http://localhost:8080/api/v1/favo/player/:player_user_id/favo-user/:favo_user_id/matching-format/:matching_format_id
```

#### Request Parameters
|場所|パラメータ名|指定する値|説明|
|:-|:-|:-|:-|
|Path|player_user_id|数値|変更したいユーザーのID|
|Path|favo_user_id|数値|いいね！された一覧から削除したいユーザーのID|
|Path|matching_format_id|数値|いいね！してきた削除したいユーザーのマッチング形式のID|

#### Example Request
```
curl http://http://54.238.92.95:8080/api/v1/favo/player/1/favo-user/7/matching-format/2 -X DELETE -H "Content-Type: application/json"

```
---