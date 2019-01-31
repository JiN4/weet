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

func GetMutualFavoUsersById(c *gin.Context) {
	allFavoUsers := service.AllFavoUsers{}

	userId, err := GetUint(c, "user_id")
	allFavoUsers = model.GetMutualFavoUsersById(userId)
	if err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, allFavoUsers)
}

func PostMutualFavoUsers(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(body))

	var mutualFavoUser model.MutualFavoUser
	err = json.Unmarshal(body, &mutualFavoUser)

	mutualFavoUser, err = model.CreateMutualFavoUser(mutualFavoUser)
	if err != nil {
		panic(err)
	}

}
