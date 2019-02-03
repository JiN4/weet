package router

import (
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	api1 := r.Group("/api/v1")
	api2 := r.Group("/api/v2")
	apiRouter(api1, api2)

	return r
}
