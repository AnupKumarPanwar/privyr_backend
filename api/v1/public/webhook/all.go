package webhook

import (
	"net/http"
	"privyr/database"
	"privyr/database/models"
	"privyr/response"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func All(c *gin.Context) {
	db := database.Db
	userIdStr := c.Param("user_id")

	userId, err := uuid.Parse(userIdStr)

	if err != nil {
		zap.L().Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var webhooks []response.Webhook

	webhook := models.Webhook{
		UserID: userId,
	}

	db.Where(webhook).Find(&webhooks)

	c.JSON(http.StatusOK, gin.H{"data": webhooks})
}
