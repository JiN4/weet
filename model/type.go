package model

import (
	"github.com/jinzhu/gorm"
)

//type //model struct {
//    ID      uint `gorm:"primary_key"`
//    CreatedAt time.Time
//    UpdatedAt time.Time
//}

// Sample情報
type Sample struct {
	gorm.Model
	Name    string `gorm:"not null"`
	Comment string `gorm:"not null"`
}

//基本的なユーザ情報
type UserBasics struct {
	gorm.Model
	UserName    string `gorm:"not null"`
	Image1      string `gorm:"not null"`
	Image2      string
	Image3      string
	Age         uint   `gorm:"not null"`
	Sex         string `gorm:"not null"`
	ResidenceID uint
	Hitokoto    string
	Comment     string
}

//ユーザごとの質疑応答
type UserQuestionAndAnswer struct {
	gorm.Model
	MatchingFormatID uint `gorm:"not null"`
	UserID           uint `gorm:"not null"`
	QuestionID       uint `gorm:"not null"`
	AnswerID         uint
	Question         Question
	Answer           Answer
}

//ユーザごとの質疑応答(理想像)
type UserIdealQuestionAndAnswer struct {
	gorm.Model
	MatchingFormatID uint `gorm:"not null"`
	UserID           uint `gorm:"not null"`
	QuestionID       uint `gorm:"not null"`
	AnswerID         uint
	Question         Question
	Answer           Answer
}

//マイページの質問
type Question struct {
	gorm.Model
	Name             string `gorm:"not null"`
	MatchingFormatID uint   `gorm:"not null"`
	Answer           []Answer
}

//マイページの質問の答えの選択肢
type Answer struct {
	gorm.Model
	Name       string `gorm:"not null"`
	QuestionID uint   `gorm:"not null"`
}

//マッチング形式
type MatchingFormat struct {
	gorm.Model
	Name string `gorm:"not null"`
}

//性別
type UserSex struct {
	UserID uint `sql:"type:int" gorm:"primary_key"`
	SexID  uint
}

//都道府県
type Prefectures struct {
	ID   uint `gorm:"primary_key"`
	Name string
}

//マッチングしたい都道府県
type MatchingPrefectures struct {
	gorm.Model
	UserID        uint
	PrefecturesID string
}

//マッチングしたい性別
type MatchingSex struct {
	UserID uint `sql:"type:int" gorm:"primary_key"`
	SexID  uint //男性：１ 女性：２ どちらも：３
}

//マッチングしたい年齢
type MatchingAge struct {
	UserID   uint `sql:"type:int" gorm:"primary_key"`
	FirstAge uint
	LastAge  uint
}

//マッチングされたいマッチング形式
type MatchingFormatChoice struct {
	UserID   uint `sql:"type:int" gorm:"primary_key"`
	Love     uint
	Marriage uint
	Roommate uint
}

//プレイヤーにいいねしたユーザ
type FavoUser struct {
	gorm.Model
	MatchingFormatID uint
	PlayerUserID     uint
	FavoUserID       uint
}

//お互いにいいねしたユーザ
type MutualFavoUser struct {
	gorm.Model
	MatchingFormatID uint
	UserID1          uint
	UserID2          uint
}

//メッセージしてるユーザ
type MessageUser struct {
	gorm.Model
	MatchingFormatID uint
	UserID1          uint
	UserID2          uint
}
