package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/stevenfamy/go-web/controllers"
)

const routesGroup string = "/auth"

func AuthRoutes(api *gin.RouterGroup) {
	api.POST(routesGroup+"/login", controllers.AuthLogin)
}
