package initializers

import (
	"example/web-service-gin/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	//dsn := os.Getenv("db")
	dsn := "root@tcp(127.0.0.1)/db_golang?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed connect to db")
	}
	models.DB = DB                         // Assign the database connection globally
	DB.AutoMigrate(&models.Departements{}) // Ensure the table exists
	DB.AutoMigrate(&models.User{})
}
