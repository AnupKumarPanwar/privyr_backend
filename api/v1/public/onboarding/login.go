package onboarding

import (
	"net/http"
	"privyr/database"
	"privyr/database/models"
	"privyr/requests"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func Login(c *gin.Context) {
	tx := database.Db.Begin()
	var request requests.Login
	err := c.ShouldBindJSON(&request)

	if err != nil {
		tx.Rollback()
		zap.L().Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Email: request.Email,
	}

	err = tx.Where(user).First(&user).Error

	if err == nil {
		c.JSON(http.StatusOK, gin.H{"data": gin.H{"id": user.ID, "email": user.Email}})
		return
	}

	user = models.User{
		Email: request.Email,
	}

	err = tx.Create(&user).Error

	if err != nil {
		zap.L().Error(err.Error())
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tx.Commit()
	c.JSON(http.StatusOK, gin.H{"data": gin.H{"id": user.ID, "email": user.Email}})
}
