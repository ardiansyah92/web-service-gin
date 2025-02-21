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
	router.POST("/departemen", controllers.PostDepartement)
	router.GET("/departemen", controllers.GetDepartement)
	router.GET("/departemen/:id", controllers.GetDepartementId)
	router.PUT("/departemenid/:id", controllers.PutDepartementId)
	router.DELETE("/departemen/:id", controllers.DeleteDepartement)

	router.Run()
}
