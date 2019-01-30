package model

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/weet/service"
	"io/ioutil"
	"log"
)

// userBasicsTableにInsertする
func CreateUserBasics(userBasics UserBasics) (UserBasics, error) {
	err := db.Create(&userBasics).Error
	if err != nil {
		return UserBasics{}, err
	}
	return userBasics, nil
}

// userBasicsTableにUpdateする
func UpdateUserBasics(c *gin.Context, userId uint) error {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))

	var userBasics UserBasics
	err = json.Unmarshal(body, &userBasics)

	err = db.Model(&userBasics).Where("id = ?", userId).Updates(UserBasics{
		UserName: userBasics.UserName,
		Image1:   userBasics.Image1,
		Image2:   userBasics.Image2,
		Image3:   userBasics.Image3,
		Age:      userBasics.Age,
		Sex:      userBasics.Sex,
		Hitokoto: userBasics.Hitokoto,
		Comment:  userBasics.Comment,
	}).Error

	return err
}

func GetUserBasicsById(userId uint) (service.UserBasics, error) {
	userBasics := service.UserBasics{}
	err := db.Where("id = ?", userId).First(&userBasics).Error
	userBasics.MatchingFormatName = "基本情報"
	return userBasics, err
}

/*
//基本的な情報
type UserBasics struct {
	FormatName string `json:"format_name"`
	UserName   string `json:"user_name"`
	Image1     string `json:"image1"`
	Image2     string `json:"image2"`
	Image3     string `json:"image3"`
	Age        string `json:"age"`
	Hitokoto   string `json:"hitokoto"`
	Comment    string `json:"comment"`
}
*/
