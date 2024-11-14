package helpers

import (

	"log"
	"os"
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
	"QueueIt/database"
)

type SignedDetails struct {
	Email     string
	First_name string
	Last_name  string
	Uid        string
	
	jwt.StandardClaims
}

var SECRET_KEY = os.Getenv("SECRET_KEY")
var db *gorm.DB = database.DB

func GenerateAllTokens(email string, firstName string, lastName string) (signedToken string, err error) {
	claims := &SignedDetails{
		Email:      email,
		First_name: firstName,
		Last_name:  lastName,
		
	}

	// Generate token using the secret key
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
		return
	}

	return token, err
}

// ValidateToken verifies the given JWT token
func ValidateToken(signedToken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&SignedDetails{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)

	if err != nil {
		msg = err.Error()
		return
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = "the token is invalid"
		return
	}


	return claims, msg
}
