package model

import (
	"time"
	_ "time"
)

// //model内で共通のfieldをここに抜き出す
//type //model struct {

//}

// 自分のユーザー情報
type Sample struct {
	//Model
	ID        uint `gorm:"primary_key" json:"id"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `gorm:"not null"`
	Comment   string `gorm:"not null"`
}
