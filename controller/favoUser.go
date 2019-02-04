package controller

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/weet/model"
	"github.com/weet/service"
)

func GetFavoUsersById(c *gin.Context) {
	allFavoUsers := service.AllFavoUsers{}

	userId, err := GetUint(c, "user_id")
	allFavoUsers = model.GetFavoUsersById(userId)
	if err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, allFavoUsers)
}

func PostFavoUsers(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))

	var favoUser model.FavoUser
	err = json.Unmarshal(body, &favoUser)

	favoUser, err = model.CreateFavoUser(favoUser)
	if err != nil {
		panic(err)
	}
}

func DeleteFavoUsers(c *gin.Context) {
	var err error
	var playerUserId, favoUserId, matchingFormatId uint

	playerUserId, err = GetUint(c, "player_user_id")
	if err != nil {
		log.Println(err)
	}
	favoUserId, err = GetUint(c, "favo_user_id")
	if err != nil {
		log.Println(err)
	}
	matchingFormatId, err = GetUint(c, "matching_format_id")
	if err != nil {
		log.Println(err)
	}

	err = model.DeleteFavoUsers(playerUserId, favoUserId, matchingFormatId)
}
