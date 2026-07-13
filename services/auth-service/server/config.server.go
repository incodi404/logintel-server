package server

import (
	"auth-service/middlewares"
	"auth-service/routes"
	"auth-service/utils"
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitializeServer() error {
	// getting env values
	corsOrigin := utils.GetenvWithDefaultValue("CORS_ORIGIN", "*")
	port := utils.GetenvWithDefaultValue("PORT", "7000")

	// initalize logger
	InitializeLogger()

	router := gin.Default()

	// cors config
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{corsOrigin},
		AllowCredentials: true,
	}))

	// error handler
	router.Use(middlewares.ErrorHandler())

	// register router
	routes.SetupRouter(router)

	// health
	router.GET("/hi", func(ctx *gin.Context) {
		ctx.String(200, "Hello!")
	})

	fmt.Println("[INFO] Server is listening at ", port)

	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		return fmt.Errorf("[ERROR] Error occurred while starting the server: %w", err)
	}

	return nil
}
