package controllers

import (

	database "QueueIt/database"
	"fmt"
	"log"
	"net/http"
	"QueueIt/models"
	"github.com/gin-gonic/gin"
)

func CreateBusiness() gin.HandlerFunc {
	return func(c *gin.Context) {

		var business models.Business

		err := c.ShouldBindJSON(&business)
		fmt.Println(err)	
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		validationErr := validate.Struct(&business)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			log.Println(err)
			return
		}

		if result := database.DB.Create(&business); result.Error != nil {
			fmt.Println(result.Error)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create business"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Business created successfully"})
	}
}

func GetAllBusinesses() gin.HandlerFunc {
	return func(c *gin.Context) {
		var businesses []models.Business

		if result := database.DB.Find(&businesses); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve businesses"})
			return
		}

		c.JSON(http.StatusOK, businesses)
	}
}

func GetBusiness() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("business_id")

		var business models.Business
		if result := database.DB.Where("id = ?", userId).First(&business); result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Business not found"})
			return
		}

		c.JSON(http.StatusOK, business)
	}
}

// Do I need to make an api for Updating BUSINESS and deleting BUSINESS ??


