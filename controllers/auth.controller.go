package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stevenfamy/go-web/models"
)

func AuthLogin(c *gin.Context) {
	var loginIdentification models.AuthLoginModel
	var userData models.User

	if err := c.ShouldBindJSON(&loginIdentification); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := models.DB.QueryRow("SELECT * FROM user WHERE email_address = ? or username = ?", loginIdentification.Identification, loginIdentification.Identification).Scan(&userData.ID, &userData.UserName, &userData.EmailAddress, &userData.Password, &userData.CreatedAt, &userData.UpdatedAt)

	if err != nil {
		log.Printf(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email Address or Username not found!"})
		return
	}

	log.Println(userData.ID)

	// for results.Next() {
	// 	err = results.Scan(userData.ID, userData.UserName, userData.EmailAddress, userData.Password)
	// 	if err != nil {
	// 		panic(err.Error())
	// 	}

	// 	log.Printf(userData.ID)
	// }

	// log.Printf(string(result))

	c.JSON(http.StatusOK, "ok")
}
