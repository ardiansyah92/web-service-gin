package controllers

import (
	"example/web-service-gin/initializers"
	"example/web-service-gin/models"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// Secret key for signing JWT tokens
var jwtSecret = []byte("your-secret-key")

// GenerateJWT creates a JWT token for a user
func GenerateJWT(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(), // Token expires in 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    "401",
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}

		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		claims, err := ValidateJWT(tokenString)
		if err != nil || claims == nil {
			c.JSON(http.StatusUnauthorized, gin.H{"code": "401", "message": "Invalid Token"})
			c.Abort()
			return
		}

		fmt.Printf("Claims: %+v\n", claims)

		c.Set("username", (*claims)["username"])

		c.Next()
	}
}

// ValidateJWT parses and verifies a JWT token
func ValidateJWT(tokenString string) (*jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

func Register(c *gin.Context) {
	var request struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Role     bool   `json:"role"`
		Phone    string `jso:"number"`
		Email    string `json:email`
		Address  string `json:address`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid input",
			"code":    "400",
		})
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to hash password",
			"code":    "500",
		})
		return
	}

	// Check if user already exists
	var existingUser models.Users
	if err := models.DB.Where("username = ?", request.Username).First(&existingUser).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{
			"message": "User already exists",
			"code":    "409",
		})
		return
	}

	// Create user in database
	user := models.Users{Username: request.Username, Password: string(hashedPassword), Role: request.Role, Phone: request.Phone, Email: request.Email, Address: request.Address}
	result := models.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to register user",
			"code":    "500",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Success Register User",
		"code":    "200",
	})
}

// Login
func Login(c *gin.Context) {
	var request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Input",
			"code":    "404",
		})
	}
	var user models.Users
	result := models.DB.Where("username = ?", request.Username).First(&user)

	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "User not found",
			"code":    "404",
		})
		return
	}

	// Compare password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Incorrect password",
			"code":    "404",
		})
		return
	}

	// Prepare user response
	// datauserresponse := struct {
	// 	ID       uint   `json:"id_user"`
	// 	Username string `json:"username"`
	// 	Role     bool   `json:"role"`
	// 	Phone    string `json:"phone"`
	// 	Email    string `json:"email"`
	// 	Address  string `json:"address"`
	// }{
	// 	ID:       user.ID,
	// 	Username: user.Username,
	// 	Role:     user.Role,
	// 	Phone:    user.Phone,
	// 	Email:    user.Email,
	// 	Address:  user.Address,
	// }

	// Generate JWT token
	token, _ := GenerateJWT(user.Username)

	c.JSON(http.StatusOK, gin.H{
		// "data":    datauserresponse,
		"token":   token,
		"message": "Success Login",
		"code":    "200",
	})
}

// getAlbums responds with the list of all albums as JSON.
func GetAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Get Data Album",
		"data":    models.Albums,
		"code":    "200",
	})
}

// postAlbums adds an album from JSON received in the request body.
func PostAlbums(c *gin.Context) {
	var newAlbums models.Album

	// Call BindJson to bind the received JSON to New Albums

	if err := c.BindJSON(&newAlbums); err != nil {
		return
	}
	// Add the albums to the slice
	models.Albums = append(models.Albums, newAlbums)
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Get Data Album",
		"data":    newAlbums,
		"code":    "200",
	})

}

// Get Albums by ID
func GetAlbumsByID(c *gin.Context) {
	id := c.Param("id")

	//Loop Over the list of albums, looking for an albums whose ID value matches the parameter

	for _, a := range models.Albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, gin.H{
				"message": "Get Data By ID",
				"data":    a,
				"code":    "200",
			})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{
		"message": "albums not found",
		"code":    "404",
	})
}

// postDepartement adds an Departement from JSON received in the request body.
func PostDepartement(c *gin.Context) {
	var newDepartement models.Departements

	// Call BindJson to bind the received JSON to New Departement

	if err := c.BindJSON(&newDepartement); err != nil {
		return
	}

	// Insert into the database
	if err := initializers.DB.Create(&newDepartement).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert data"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Create Data Departemen",
		"data":    newDepartement,
		"code":    "200",
	})

}

func GetDepartement(c *gin.Context) {

	var getDepartements []models.Departements

	// Fetch data from the database
	if err := initializers.DB.Find(&getDepartements).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to fetch data",
			"message": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Get Data Departement",
		"data":    getDepartements,
		"code":    "200",
	})
}

// GetProfile
func GetProfile(c *gin.Context) {
	username, exists := c.Get("username")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Unauthorized",
			"code":    "401",
		})
		return
	}
	var user models.Users
	result := models.DB.Where("username = ?", username).First(&user)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Profile not found",
			"code":    "404",
		})
		return
	}
	// Prepare user response
	getMeprofile := struct {
		ID       uint   `json:"id_user"`
		Username string `json:"username"`
		Role     bool   `json:"role"`
		Phone    string `json:"phone"`
		Email    string `json:"email"`
		Address  string `json:"address"`
	}{
		ID:       user.ID,
		Username: user.Username,
		Role:     user.Role,
		Phone:    user.Phone,
		Email:    user.Email,
		Address:  user.Address,
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Get Profil Me",
		"data":    getMeprofile,
		"code":    "200",
	})
}

func GetUser(c *gin.Context) {

	// Fetch users from the database
	var getUsers []models.Users
	if err := initializers.DB.Find(&getUsers).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to fetch data",
			"message": err.Error(),
		})
		return
	}

	// Prepare response without exposing passwords
	var usersResponse []struct {
		ID       uint   `json:"id_user"`
		Username string `json:"username"`
		Role     bool   `json:"role"`
		Phone    string `json:"phone"`
		Email    string `json:"email"`
		Address  string `json:"address"`
	}

	for _, user := range getUsers {
		usersResponse = append(usersResponse, struct {
			ID       uint   `json:"id_user"`
			Username string `json:"username"`
			Role     bool   `json:"role"`
			Phone    string `json:"phone"`
			Email    string `json:"email"`
			Address  string `json:"address"`
		}{
			ID:       user.ID,
			Username: user.Username,
			Role:     user.Role,
			Phone:    user.Phone,
			Email:    user.Email,
			Address:  user.Address,
		})
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Get Data User",
		"data":    usersResponse,
		"code":    "200",
	})
}

func GetDepartementId(c *gin.Context) {
	id := c.Param("id")

	var dept models.Departements // Use the correct struct name

	// Query database for department by ID
	if err := models.DB.First(&dept, "id = ?", id).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "Department not found",
			"code":    404,
		})
		return
	}

	// Return department if found
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Get Data Departemen By ID",
		"data":    dept,
		"code":    200,
	})
}

// PutDepartementId updates a department by ID in the database
func PutDepartementId(c *gin.Context) {
	id := c.Param("id")
	var dept models.Departements

	// Bind JSON body to the department struct
	if err := c.BindJSON(&dept); err != nil {
		return
	}
	if err := models.DB.Model(&models.Departements{}).Where("id = ?", id).Updates(models.Departements{
		DepartementName: dept.DepartementName,
		Location:        dept.Location,
	}).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": "Error updating department",
			"code":    500,
		})

		return
	}

	// Return updated department
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Department updated successfully",
		"data":    dept,
		"code":    200,
	})
}

func DeleteDepartement(c *gin.Context) {
	// Get the ID from the URL params
	idStr := c.Param("id")

	// Convert id to integer (if your ID is numeric)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"message": "Invalid ID format",
			"code":    400,
		})
		return
	}

	var dept models.Departements
	// Check if the department exists
	if err := models.DB.First(&dept, id).Error; err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{
			"message": "Department not found",
			"code":    404,
		})
		return
	}

	// Delete the department
	if err := models.DB.Delete(&dept).Error; err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{
			"message": "Error deleting department",
			"code":    500,
		})
		return
	}

	// Return success response
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Department deleted successfully",
		"code":    200,
	})
}
