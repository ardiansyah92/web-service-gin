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

	auth := router.Group("/")
	auth.Use(controllers.JWTAuthMiddleware())

	{
		auth.POST("departemen/", controllers.PostDepartement)
		auth.GET("departemen/", controllers.GetDepartement)
		auth.GET("departemen/:id", controllers.GetDepartementId)
		auth.PUT("departemen/:id", controllers.PutDepartementId)
		auth.DELETE("departemen/:id", controllers.DeleteDepartement)
		auth.GET("users/", controllers.GetUser)
		auth.GET("me/", controllers.GetProfile)
		auth.POST("loan/", controllers.PostLoan)
		auth.GET("loan/", controllers.GetLoan)
		auth.GET("loanview/", controllers.GetLoanUser)
		auth.POST("uploadfile/", controllers.UploadFile)
	}

	router.Run()
}
