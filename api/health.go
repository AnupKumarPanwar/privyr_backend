package api

import (
	"net/http"
	"privyr/database"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Status(c *gin.Context) {
	db := database.Db

	conn, err := db.DB()
	if err != nil {
		zap.L().Error(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	err = conn.Ping()
	if err != nil {
		zap.L().Error(err.Error())
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.String(http.StatusOK, "pong")
}
