package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/stevenfamy/go-web/controllers"
	"github.com/stevenfamy/go-web/middleware"
)

func TestRoutes(api *gin.RouterGroup) {
	const routesGroup string = "/test"
	api.GET(routesGroup+"/helloworld", middleware.AuthorizationMiddleware(), controllers.HelloWorld)
}
