package main

import (
	"example/web-service-gin/controllers"
	"example/web-service-gin/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	// Initialize the database connection
	initializers.ConnectToDB()
}

func main() {
	router := gin.Default()
	router.GET("/albums", controllers.GetAlbums)
	router.POST("/albums", controllers.PostAlbums)
	router.GET("/albums/:id", controllers.GetAlbumsByID)
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)
	// Protected Routes (JWT Middleware)

	auth := router.Group("/departemen")
	auth.Use(controllers.JWTAuthMiddleware())

	{
		auth.POST("/", controllers.PostDepartement)
		auth.GET("/", controllers.GetDepartement)
		auth.GET("/:id", controllers.GetDepartementId)
		auth.PUT("/:id", controllers.PutDepartementId)
		auth.DELETE("/:id", controllers.DeleteDepartement)
	}

	router.Run()
}
