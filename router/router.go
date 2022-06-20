package router

import (
	"privyr/api"
	"privyr/api/v1/public/leads"
	"privyr/api/v1/public/onboarding"
	"privyr/api/v1/public/webhook"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func Setup() {
	Router = gin.New()

	Router.Use(gin.Logger())
	Router.Use(gin.Recovery())

	Router.GET("/ping", api.Status)

	v1 := Router.Group("v1")
	{
		public := v1.Group("")
		{
			onboarding.AddRoutes(public)
			webhook.AddRoutes(public)
			leads.AddRoutes(public)
		}
	}
}
