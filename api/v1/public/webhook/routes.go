package webhook

import "github.com/gin-gonic/gin"

func AddRoutes(route *gin.RouterGroup) {
	routeGroup := route.Group("webhook")
	{
		routeGroup.POST("/", Create)
	}
}