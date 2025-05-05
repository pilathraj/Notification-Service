package notifications

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"notification-service/pkg/models"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"
)

func CreateNotification(ctx *gin.Context, producer *kafka.Writer, db *gorm.DB) {
	var notification models.Notification
	if err := ctx.ShouldBindJSON(&notification); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := notification.UserID

	notificationJSON, err := json.Marshal(notification)
	if err != nil {
		fmt.Println("Error marshalling notification:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to marshal notification"})
		return
	}

	// Send notification to Kafka
	message := kafka.Message{
		Key:   []byte(userID),
		Value: []byte(notificationJSON),
	}
	if err := producer.WriteMessages(context.Background(), message); err != nil {
		fmt.Println("Failed to send notification:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to send notification to Kafka"})
		return
	}

	// Save notification to the database
	if err := db.Create(&notification).Error; err != nil {
		fmt.Println("Failed to save notification:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save notification to database"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "notification created successfully"})
}
