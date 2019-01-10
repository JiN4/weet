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
	UserName string `gorm:"not null"`
	Image1   string `gorm:"not null"`
	Image2   string `gorm:"not null"`
	Image3   string `gorm:"not null"`
	Age      string `gorm:"not null"`
	Hitokoto string `gorm:"not null"`
	Comment  string `gorm:"not null"`
}

//ユーザごとの質疑応答
type UserQuestionAndAnswer struct {
	gorm.Model
	UserID     uint     `gorm:"not null"`
	QuestionID uint     `gorm:"not null"`
	AnswerID   uint     `gorm:"not null"`
	Question   Question `gorm:"foreignkey:QuestionID"`
	Answer     Answer   `gorm:"foreignkey:AnswerID"`
}

//マイページの質問
type Question struct {
	ID       uint   `gorm:"primary_key"`
	FormatID uint   `gorm:"not null"`
	Name     string `gorm:"not null"`
}

//マイページの質問の答えの選択肢
type Answer struct {
	ID         uint   `gorm:"primary_key"`
	QuestionID uint   `gorm:"not null"`
	Name       string `gorm:"not null"`
}

//マッチング形式
type Format struct {
	ID   uint   `gorm:"primary_key"`
	Name string `gorm:"not null"`
}
