package notifications

import (
	"net/http"
	"notification-service/src/pkg/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUserNotifications(ctx *gin.Context, db *gorm.DB) {
	userID := ctx.Param("userId")
	var notifications []models.Notification
	if err := db.Where("user_id = ?", userID).Find(&notifications).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch notifications"})
		return
	}
	if len(notifications) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	ctx.JSON(http.StatusOK, notifications)
}
