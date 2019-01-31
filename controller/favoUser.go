package controller

import (
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
