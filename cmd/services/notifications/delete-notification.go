package notifications

import (
	"net/http"
	"notification-service/pkg/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func DeleteNotification(ctx *gin.Context, db *gorm.DB) {
	notificationID := ctx.Param("notificationId")
	if err := db.Delete(&models.Notification{}, notificationID).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete notification"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "notification deleted successfully"})
}
