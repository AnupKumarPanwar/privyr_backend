package router

import (
	"privyr/api"

	"github.com/gin-gonic/gin"
)

var Router *gin.Engine

func Setup() {
	Router = gin.New()

	Router.Use(gin.Logger())
	Router.Use(gin.Recovery())

	Router.GET("/ping", api.Status)
}
