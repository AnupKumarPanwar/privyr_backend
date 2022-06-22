package webhook

import (
	"net/http"
	"privyr/database"
	"privyr/database/models"
	"privyr/requests"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Create(c *gin.Context) {
	tx := database.Db.Begin()
	var request requests.Webhook
	err := c.ShouldBindJSON(&request)

	if err != nil {
		tx.Rollback()
		zap.L().Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	webhook := models.Webhook{
		Name:   request.Name,
		UserID: request.UserID,
	}

	err = tx.Create(&webhook).Error

	if err != nil {
		zap.L().Error(err.Error())
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"data": gin.H{"id": webhook.ID, "name": webhook.Name}})
}
