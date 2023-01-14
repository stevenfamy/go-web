package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stevenfamy/go-web/helpers"
	"github.com/stevenfamy/go-web/models"
)

func AuthLogin(c *gin.Context) {
	var loginIdentification models.AuthLogin
	var userData models.User

	if err := c.ShouldBindJSON(&loginIdentification); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := models.DB.QueryRow("SELECT * FROM user WHERE email_address = ? or username = ?", loginIdentification.Identification, loginIdentification.Identification).Scan(&userData.ID, &userData.UserName, &userData.EmailAddress, &userData.Password, &userData.Status, &userData.CreatedAt, &userData.UpdatedAt)

	if err != nil {
		log.Printf(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email Address or Username not found!"})
		return
	}

	if userData.Status != 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User not active!"})
		return
	}

	if res := helpers.CheckPasswordHash(loginIdentification.Password, userData.Password); !res {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password not match!"})
		return
	}

	jwt, err := helpers.GenerateJWT(userData.ID)
	if err != nil {
		log.Printf(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"authToken": jwt})
}
