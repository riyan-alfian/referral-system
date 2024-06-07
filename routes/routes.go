package routes

import (
	"net/http"
	"referral-system-2/controllers"
	"referral-system-2/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Public routes
	application := "referral system 1.0"
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusCreated, gin.H{"application": application})

	})

	public := router.Group("/api")
	{
		public.POST("/register", controllers.Register)
		public.POST("/login", controllers.Login)
	}

	// Authenticated routes
	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/contribute/:referralLink", controllers.CreateContribution)
		protected.POST("/generate_new_link", controllers.GenerateNewLink)
	}

	return router
}
