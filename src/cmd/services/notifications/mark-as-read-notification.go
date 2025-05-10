package notifications

import (
	"net/http"
	"notification-service/src/pkg/models"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func MarkNotificationAsRead(ctx *gin.Context, db *gorm.DB) {
	notificationID := ctx.Param("notificationId")
	if err := db.Model(&models.Notification{}).Where("id = ?", notificationID).Updates(&models.Notification{IsRead: true, ReadAt: time.Now()}).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to mark notification as read"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "notification marked as read"})
}
