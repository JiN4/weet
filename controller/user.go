package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/weet/model"
	"github.com/weet/service"
)

func GetUserById(c *gin.Context) {
	userId, err := GetUint(c, "user_id")

	user := service.User{}

	user.UserBasics, err = model.GetUserBasicsById(userId)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	user.UserSpecials, err = model.GetUserSpecialsById(userId)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	user.UserIdealSpecials, err = model.GetUserIdealSpecialsById(userId)
	if err != nil {
		log.Println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.JSON(http.StatusOK, user)
}

func UpdateUserBasics(c *gin.Context) {
	userId, err := GetUint(c, "user_id")
	err = model.UpdateUserBasics(c, userId)
	if err != nil {
		log.Println(err)
	}
}
