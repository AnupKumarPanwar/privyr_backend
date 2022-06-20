package webhook

import (
	"net/http"
	"privyr/database"
	"privyr/database/models"
	"privyr/requests"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func Listener(c *gin.Context) {
	tx := database.Db.Begin()

	webhookIdStr := c.Param("webhook_id")

	webhookId, err := uuid.Parse(webhookIdStr)

	var request requests.Lead
	err = c.ShouldBindJSON(&request)

	if err != nil {
		tx.Rollback()
		zap.L().Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	lead := models.Lead{
		Name:        request.Name,
		Email:       request.Email,
		Phone:       request.Phone,
		OtherFields: request.OtherFields,
		WebhookID:   webhookId,
	}

	err = tx.Create(&lead).Error

	if err != nil {
		zap.L().Error(err.Error())
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"data": gin.H{"id": lead.ID}})
}
