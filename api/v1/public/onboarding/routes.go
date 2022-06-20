package onboarding

import "github.com/gin-gonic/gin"

func AddRoutes(route *gin.RouterGroup) {
	routeGroup := route.Group("login")
	{
		routeGroup.POST("/", Login)
	}
}
