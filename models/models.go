package models

import "gorm.io/gorm"

var DB *gorm.DB

// album represents data about a record album.
type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var Albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

type Departements struct {
	ID              string `json:"id"`
	DepartementName string `json:"departement_name"`
	Location        string `json:"location"`
}

type Users struct {
	ID       uint   `json:"id_user" gorm:"primaryKey;autoIncrement"`
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`
	Role     bool   `json:"role" gorm:"default:false"`
	Phone    string `json:"phone" gorm:"unique;not null"`
	Email    string `json:"email" gorm:"unique;not null"`
	Address  string `json:"address" gorm:"unique;not null"`
}
