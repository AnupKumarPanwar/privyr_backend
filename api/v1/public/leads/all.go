package leads

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
	webhookIdStr := c.Param("webhook_id")

	webhookId, err := uuid.Parse(webhookIdStr)

	if err != nil {
		zap.L().Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var leads []response.Lead

	webhook := models.Lead{
		WebhookID: webhookId,
	}

	db.Where(webhook).Find(&leads)

	c.JSON(http.StatusOK, gin.H{"data": leads})
}
