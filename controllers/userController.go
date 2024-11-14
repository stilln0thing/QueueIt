package controllers

import (
	database "QueueIt/database"
	helper "QueueIt/helpers"
	"QueueIt/models"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

// Use the initialized DB from the database package


var validate = validator.New()

// HashPassword hashes a given password using bcrypt.
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Panic(err)
	}
	return string(hashedPassword), nil
}

// VerifyPassword compares a given password with a hashed password.
func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}


// Signup handles user registration.
func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		
		var user models.User
        
		// Bind JSON to the user struct and validate
		err := c.ShouldBindJSON(&user)
		fmt.Println(err)	
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
																	
		validationErr := validate.Struct(&user)
		if validationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
			log.Println(err)
			return
		}
		token, _ := helper.GenerateAllTokens(user.Email, user.FirstName, user.LastName)
		user.Token = token
		// Hash the password
		hashedPassword, err := HashPassword(user.Password)
		if err != nil {	
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
			return
		}
		user.Password = hashedPassword
	
		// Save the user in the database
		if result := database.DB.Create(&user); result.Error != nil {
			fmt.Println(result.Error)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
			return
		}
		
		c.JSON(http.StatusOK, gin.H{"message": "User created successfully"})
	}
}

// Login handles user authentication.
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody struct {
			Email    string ` binding:"required,email"`
			Password string ` binding:"required"`
		}
		fmt.Println("1")
		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("11")
		var user models.User
		if result := database.DB.Where("Email = ?", requestBody.Email).First(&user); result.Error != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid EmailID"})
			return
		}
		fmt.Println("111")
		if err := VerifyPassword(requestBody.Password, user.Password); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Password"})
			return
		}
		fmt.Println("1111")
		c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
	}
}

// GetUsers retrieves all users from the database.
func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var users []models.User

		if result := database.DB.Find(&users); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
			return
		}

		c.JSON(http.StatusOK, users)
	}
}

// GetUser retrieves a specific user based on the user ID.
func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("user_id")

		if err := helper.MatchUserTypeToUid(c, userId); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var user models.User
		if result := database.DB.Where("id = ?", userId).First(&user); result.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}

		c.JSON(http.StatusOK, user)
	}
}
