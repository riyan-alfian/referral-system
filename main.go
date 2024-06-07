package main

import (
	"referral-system-2/config"
	"referral-system-2/routes"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Referral System API
// @version 1.0
// @description This is a referral system server.

// @host localhost:8080
// @BasePath /api

func main() {
	config.InitDB()
	defer config.CloseDB()

	router := routes.SetupRouter()

	// Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.Run(":8080")

}
