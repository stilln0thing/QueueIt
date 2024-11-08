package controllers

import(
	"fmt"
	"log"
	"net/http"
	"time"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	helper "QueueIt/helpers"
	
	"QueueIt/models"
	"QueueIt/database"
	"golang.org/x/crypto/bcrypt"
)
var db = database.DB 
var validate = validator.New()
func HashPassword()

func VerifyPassword()

func Signup()

func Login()

func GetUsers()

func GetUser() gin.HandlerFunc{
	return func(c *gin.Context){
		userId := c.Param("user_id")

		if err := helper.MatchUserTypeToUid(c, userId); err !=nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var user models.User
		user
	}
}