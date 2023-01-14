package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/stevenfamy/go-web/controllers"
)

func AuthRoutes(api *gin.RouterGroup) {
	const routesGroup string = "/auth"
	api.POST(routesGroup+"/login", controllers.AuthLogin)
}
