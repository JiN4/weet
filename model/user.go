package model

import "github.com/weet/service"

func GetUserById(userId uint) (service.User, error) {
	user := service.User{}
	var err error
	user.UserBasics, err = GetUserBasicsById(userId)
	user.UserSpecials = GetUserSpecialsById(userId)
	user.UserIdealSpecials = GetUserIdealSpecialsById(userId)
	return user, err
}

/*
// ユーザ情報-----！
type User struct {
	UserBasicInfo    UserBasicInfo    `json:"User_basic_info"`
	UserSpecialInfos UserSpecialInfos `json:"User_pecial_infos"`
}
*/
