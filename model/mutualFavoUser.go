package model

import (
	"github.com/weet/service"
)

// TableにInsertする
func CreateMutualFavoUser(mutualFavoUser MutualFavoUser) (MutualFavoUser, error) {
	err := db.Create(&mutualFavoUser).Error
	if err != nil {
		return MutualFavoUser{}, err
	}
	return mutualFavoUser, nil
}

func GetMutualFavoUsersById(userId uint) service.AllFavoUsers {
	serviceFavoUser := service.FavoUser{}
	friendFavoUsers := service.FriendFavoUsers{}
	loveFavoUsers := service.LoveFavoUsers{}
	marriageFavoUsers := service.MarriageFavoUsers{}
	roommateFavoUsers := service.RoommateFavoUsers{}
	allFavoUsers := service.AllFavoUsers{}
	userBasics := UserBasics{}
	mutualFavoUsers := []MutualFavoUser{}
	var prefectures string

	db.Where("user_id1 = ?", userId).Find(&mutualFavoUsers)

	for _, mutualFavoUser := range mutualFavoUsers {
		userBasics.ID = 0
		db.Where("id = ?", mutualFavoUser.UserID2).First(&userBasics)

		prefectures = GetPrefecturesNameById(userBasics.ResidenceID)

		serviceFavoUser.UserID = mutualFavoUser.UserID2
		serviceFavoUser.UserName = userBasics.UserName
		serviceFavoUser.Image1 = userBasics.Image1
		serviceFavoUser.Sex = userBasics.Sex
		serviceFavoUser.Age = userBasics.Age
		serviceFavoUser.Residence = prefectures
		serviceFavoUser.Hitokoto = userBasics.Hitokoto

		if mutualFavoUser.MatchingFormatID == 1 {
			friendFavoUsers = append(friendFavoUsers, serviceFavoUser)
		} else if mutualFavoUser.MatchingFormatID == 2 {
			loveFavoUsers = append(loveFavoUsers, serviceFavoUser)
		} else if mutualFavoUser.MatchingFormatID == 3 {
			marriageFavoUsers = append(marriageFavoUsers, serviceFavoUser)
		} else if mutualFavoUser.MatchingFormatID == 4 {
			roommateFavoUsers = append(roommateFavoUsers, serviceFavoUser)
		}
	}

	db.Where("user_id2 = ?", userId).Find(&mutualFavoUsers)

	for _, mutualFavoUser := range mutualFavoUsers {
		userBasics.ID = 0
		db.Where("id = ?", mutualFavoUser.UserID1).First(&userBasics)

		prefectures = GetPrefecturesNameById(userBasics.ResidenceID)

		serviceFavoUser.UserID = mutualFavoUser.UserID1
		serviceFavoUser.UserName = userBasics.UserName
		serviceFavoUser.Image1 = userBasics.Image1
		serviceFavoUser.Sex = userBasics.Sex
		serviceFavoUser.Age = userBasics.Age
		serviceFavoUser.Residence = prefectures
		serviceFavoUser.Hitokoto = userBasics.Hitokoto

		if mutualFavoUser.MatchingFormatID == 1 {
			friendFavoUsers = append(friendFavoUsers, serviceFavoUser)
		} else if mutualFavoUser.MatchingFormatID == 2 {
			loveFavoUsers = append(loveFavoUsers, serviceFavoUser)
		} else if mutualFavoUser.MatchingFormatID == 3 {
			marriageFavoUsers = append(marriageFavoUsers, serviceFavoUser)
		} else if mutualFavoUser.MatchingFormatID == 4 {
			roommateFavoUsers = append(roommateFavoUsers, serviceFavoUser)
		}
	}

	allFavoUsers.FriendFavoUsers = friendFavoUsers
	allFavoUsers.LoveFavoUsers = loveFavoUsers
	allFavoUsers.MarriageFavoUsers = marriageFavoUsers
	allFavoUsers.RoommateFavoUsers = roommateFavoUsers

	return allFavoUsers
}
