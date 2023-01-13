package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stevenfamy/go-web/models"
	"github.com/stevenfamy/go-web/routes"
)

const portNumber = ":8001"

func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	models.ConnectDatabase()

	api := r.Group("/api")
	{
		api.GET("/helloworld", helloWorld)
		routes.AuthRoutes(api)
	}

	if err := r.Run(portNumber); err != nil {
		log.Fatal(err)
	}
}

func helloWorld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld, server is active")
}
