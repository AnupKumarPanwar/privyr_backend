package leads

import "github.com/gin-gonic/gin"

func AddRoutes(route *gin.RouterGroup) {
	routeGroup := route.Group("leads")
	{
		routeGroup.GET("/:webhook_id", All)
	}
}
