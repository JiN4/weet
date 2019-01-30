package model

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/weet/service"
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
		UserName:    userBasics.UserName,
		Image1:      userBasics.Image1,
		Image2:      userBasics.Image2,
		Image3:      userBasics.Image3,
		Age:         userBasics.Age,
		Sex:         userBasics.Sex,
		ResidenceID: userBasics.ResidenceID,
		Hitokoto:    userBasics.Hitokoto,
		Comment:     userBasics.Comment,
	}).Error

	return err
}

func GetUserBasicsById(userId uint) (service.UserBasics, error) {
	serviceUserBasics := service.UserBasics{}
	userBasics := UserBasics{}
	err := db.Where("id = ?", userId).First(&userBasics).Error
	serviceUserBasics.MatchingFormatName = "基本情報"
	prefectures := GetPrefecturesNameById(userBasics.ResidenceID)

	serviceUserBasics.MatchingFormatName = "基本情報"
	serviceUserBasics.UserName = userBasics.UserName
	serviceUserBasics.Image1 = userBasics.Image1
	serviceUserBasics.Image2 = userBasics.Image2
	serviceUserBasics.Image3 = userBasics.Image3
	serviceUserBasics.Sex = userBasics.Sex
	serviceUserBasics.Age = userBasics.Age
	serviceUserBasics.Residence = prefectures
	serviceUserBasics.Hitokoto = userBasics.Hitokoto
	serviceUserBasics.Comment = userBasics.Comment

	return serviceUserBasics, err
}

func GetCandidateMatchingUserId(userId uint, matchingPrefectures []uint, matchingSex uint, matchingFirstAge uint, matchingLastAge uint, matchingFormatId uint) (uint, error) {
	userBasics := []UserBasics{}
	var err error

	if matchingFormatId == 1 {
		if matchingSex == 1 {
			err = db.Raw("select id from user_basics where Sex = '男性' and id != ? and residence_id IN (?) and age between ? and ? order by rand() LIMIT 10", userId, matchingPrefectures, matchingFirstAge, matchingLastAge).Scan(&userBasics).Error
		} else if matchingSex == 2 {
			err = db.Raw("select id from user_basics where Sex = '女性' and id != ? and residence_id IN (?) and age between ? and ? order by rand() LIMIT 10", userId, matchingPrefectures, matchingFirstAge, matchingLastAge).Scan(&userBasics).Error
		} else {
			err = db.Raw("select id from user_basics where id != ? and residence_id IN (?) and age between ? and ? order by rand() LIMIT 10", userId, matchingPrefectures, matchingFirstAge, matchingLastAge).Scan(&userBasics).Error
		}
	} else if matchingFormatId == 2 {
		if matchingSex == 1 {
			err = db.Raw("select id from user_basics inner join matching_format_choices on user_basics.id = matching_format_choices.user_id where matching_format_choices.love = 1 and Sex = '男性' and id != ? and residence_id IN (?) and age between ? and ? order by rand() LIMIT 10", userId, matchingPrefectures, matchingFirstAge, matchingLastAge).Scan(&userBasics).Error
		} else if matchingSex == 2 {
			err = db.Raw("select id from user_basics inner join matching_format_choices on user_basics.id = matching_format_choices.user_id where matching_format_choices.love = 1 and Sex = '女性' and id != ? and residence_id IN (?) and age between ? and ? order by rand() LIMIT 10", userId, matchingPrefectures, matchingFirstAge, matchingLastAge).Scan(&userBasics).Error
		} else {
			err = db.Raw("select id from user_basics inner join matching_format_choices on user_basics.id = matching_format_choices.user_id where matching_format_choices.love = 1 and id != ? and residence_id IN (?) and age between ? and ? order by rand() LIMIT 10", userId, matchingPrefectures, matchingFirstAge, matchingLastAge).Scan(&userBasics).Error
		}
	} else if matchingFormatId == 3 {
		if matchingSex == 1 {
			err = db.Raw("select id from user_basics inner join matching_format_choices on user_basics.id = matching_format_choices.user_id where matching_format_choices.marriage = 1 and Sex = '男性' and id != ? and residence_id IN (?) and age between ? and ? order by rand() LIMIT 10", userId, matchingPrefectures, matchingFirstAge, matchingLastAge).Scan(&userBasics).Error
		} else if matchingSex == 2 {
			err = db.Raw("select id from user_basics inner join matching_format_choices on user_basics.id = matching_format_choices.user_id where matching_format_choices.marriage = 1 and Sex = '女性' and id != ? and residence_id IN (?) and age between ? and ? order by rand() LIMIT 10", userId, matchingPrefectures, matchingFirstAge, matchingLastAge).Scan(&userBasics).Error
		} else {
			err = db.Raw("select id from user_basics inner join matching_format_choices on user_basics.id = matching_format_choices.user_id where matching_format_choices.marriage = 1 and id != ? and residence_id IN (?) and age between ? and ? order by rand() LIMIT 10", userId, matchingPrefectures, matchingFirstAge, matchingLastAge).Scan(&userBasics).Error
		}
	} else if matchingFormatId == 4 {
		if matchingSex == 1 {
			err = db.Raw("select id from user_basics inner join matching_format_choices on user_basics.id = matching_format_choices.user_id where matching_format_choices.roommate = 1 and Sex = '男性' and id != ? and residence_id IN (?) and age between ? and ? order by rand() LIMIT 10", userId, matchingPrefectures, matchingFirstAge, matchingLastAge).Scan(&userBasics).Error
		} else if matchingSex == 2 {
			err = db.Raw("select id from user_basics inner join matching_format_choices on user_basics.id = matching_format_choices.user_id where matching_format_choices.roommate = 1 and Sex = '女性' and id != ? and residence_id IN (?) and age between ? and ? order by rand() LIMIT 10", userId, matchingPrefectures, matchingFirstAge, matchingLastAge).Scan(&userBasics).Error
		} else {
			err = db.Raw("select id from user_basics inner join matching_format_choices on user_basics.id = matching_format_choices.user_id where matching_format_choices.roommate = 1 and id != ? and residence_id IN (?) and age between ? and ? order by rand() LIMIT 10", userId, matchingPrefectures, matchingFirstAge, matchingLastAge).Scan(&userBasics).Error
		}
	}

	fmt.Println(userBasics)
	return userBasics[0].ID, err
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
