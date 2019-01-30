package model

// func GetCandidateMatchingUserId(matchingPrefectures string, matchingSex uint, matchingFirstAge uint, matchingLastAge uint)(uint, error){

// 		err := db.Raw("select id AS user_id from user_basics where residence = ? and age between ? and ? ", matchingPrefectures, matchingFirstAge, matchingLastAge).Scan(&userQuestionAndAnswers).Error

// }

// //基本的なユーザ情報
// type UserBasics struct {
// 	gorm.Model
// 	UserName  string `gorm:"not null"`
// 	Image1    string `gorm:"not null"`
// 	Image2    string
// 	Image3    string
// 	Age       uint   `gorm:"not null"`
// 	Sex       string `gorm:"not null"`
// 	Residence string
// 	Hitokoto  string
// 	Comment   string
// }
