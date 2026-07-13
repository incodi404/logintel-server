package routes

import "github.com/gin-gonic/gin"

func SetupRouter(router *gin.Engine) {
	// api
	api := router.Group("/api")

	// version 1
	v1 := api.Group("/v1")
	RegisterV1AdminRoutes(v1)
}
