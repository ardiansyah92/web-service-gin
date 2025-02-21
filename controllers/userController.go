package controllers

import (
	"example/web-service-gin/initializers"
	"example/web-service-gin/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

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
