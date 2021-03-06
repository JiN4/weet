package model

import (
	"github.com/weet/service"
)

// TableにInsertする
func CreateFavoUser(favoUser FavoUser) (FavoUser, error) {
	err := db.Create(&favoUser).Error
	if err != nil {
		return FavoUser{}, err
	}
	return favoUser, nil
}

func GetFavoUsersById(userId uint) service.AllFavoUsers {
	serviceFavoUser := service.FavoUser{}
	friendFavoUsers := service.FriendFavoUsers{}
	loveFavoUsers := service.LoveFavoUsers{}
	marriageFavoUsers := service.MarriageFavoUsers{}
	roommateFavoUsers := service.RoommateFavoUsers{}
	allFavoUsers := service.AllFavoUsers{}
	userBasics := UserBasics{}
	favoUsers := []FavoUser{}
	var prefectures string

	db.Where("player_User_id = ?", userId).Find(&favoUsers)

	for _, favoUser := range favoUsers {
		userBasics.ID = 0
		db.Where("id = ?", favoUser.FavoUserID).First(&userBasics)

		prefectures = GetPrefecturesNameById(userBasics.ResidenceID)

		serviceFavoUser.UserID = favoUser.FavoUserID
		serviceFavoUser.UserName = userBasics.UserName
		serviceFavoUser.Image1 = userBasics.Image1
		serviceFavoUser.Sex = userBasics.Sex
		serviceFavoUser.Age = userBasics.Age
		serviceFavoUser.Residence = prefectures
		serviceFavoUser.Hitokoto = userBasics.Hitokoto

		if favoUser.MatchingFormatID == 1 {
			friendFavoUsers = append(friendFavoUsers, serviceFavoUser)
		} else if favoUser.MatchingFormatID == 2 {
			loveFavoUsers = append(loveFavoUsers, serviceFavoUser)
		} else if favoUser.MatchingFormatID == 3 {
			marriageFavoUsers = append(marriageFavoUsers, serviceFavoUser)
		} else if favoUser.MatchingFormatID == 4 {
			roommateFavoUsers = append(roommateFavoUsers, serviceFavoUser)
		}
	}
	allFavoUsers.FriendFavoUsers = friendFavoUsers
	allFavoUsers.LoveFavoUsers = loveFavoUsers
	allFavoUsers.MarriageFavoUsers = marriageFavoUsers
	allFavoUsers.RoommateFavoUsers = roommateFavoUsers

	return allFavoUsers
}

func DeleteFavoUsers(playerUserId uint, favoUserId uint, matchingFormatId uint) error {
	err := db.Where("player_user_id = ? and favo_user_id = ? and matching_format_id = ?", playerUserId, favoUserId, matchingFormatId).Delete(FavoUser{}).Error

	return err
}
