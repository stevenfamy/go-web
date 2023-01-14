package helpers

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stevenfamy/go-web/config"
	"golang.org/x/crypto/bcrypt"
)

type MyClaims struct {
	jwt.RegisteredClaims
	Selector string
	Token    string
	UserId   string
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateJWT(userId string) (string, error) {
	log.Println(userId)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyClaims{
		RegisteredClaims: jwt.RegisteredClaims{},
		Selector:         randomHex(10),
		Token:            randomHex(20),
		UserId:           userId,
	})

	tokenString, err := token.SignedString([]byte(config.Config("JWTSECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(tokenString string, c *gin.Context) (MyClaims, error) {
	var claim MyClaims

	token, err := jwt.ParseWithClaims(tokenString, &claim, func(t *jwt.Token) (interface{}, error) {
		return []byte(config.Config("JWTSECRET")), nil
	})
	if err != nil {
		return claim, errors.New("Invalid token: " + err.Error())
	}

	if !token.Valid {
		return claim, errors.New("Invalid token: " + err.Error())
	}

	return claim, nil

}

func randomHex(n int) string {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return ""
	}

	return hex.EncodeToString(bytes)
}
