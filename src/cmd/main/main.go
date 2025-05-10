package main

import (
	"context"
	"fmt"
	"log"

	notifications "notification-service/src/cmd/services/notifications"
	notifychannels "notification-service/src/cmd/services/notify-channels"
	preferences "notification-service/src/cmd/services/preferences"
	consumer_service "notification-service/src/consumer"

	"notification-service/src/pkg/models"

	"os"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/kafka-go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	LaunchPort = ":8083"
)

var KafkaBroker = os.Getenv("KAFKA_BROKERS")
var KafkaTopic = os.Getenv("KAFKA_TOPIC")

var db *gorm.DB

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {

	var err error
	dsn := os.Getenv("POSTGRES_DATABASE_DSN")
	//dsn := "host=postgres-service user=root password=postgres dbname=notifications_db port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	if dsn == "" {
		log.Fatalf("Missing POSTGRES_DATABASE_DSN environment variable")
	}
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Auto-migrate the Notification and Preference models
	if err := db.AutoMigrate(&models.Notification{}, &models.UserPreferences{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	// Initialize Kafka producer
	producer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{KafkaBroker},
		Topic:   KafkaTopic,
	})
	defer producer.Close()

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(CORSMiddleware())

	ctx, cancel := context.WithCancel(context.Background())

	store := &consumer_service.NotificationStore{
		Data: make(consumer_service.UserNotifications),
	}
	go consumer_service.SetupConsumerGroup(ctx, store)
	defer cancel()

	// Notification Management
	router.POST("/api/notifications", func(ctx *gin.Context) {
		notifications.CreateNotification(ctx, producer, db)
	})
	router.GET("/api/notifications/user/:userId", func(ctx *gin.Context) {
		notifications.GetUserNotifications(ctx, db)
	})
	router.PUT("/api/notifications/:notificationId/read", func(ctx *gin.Context) {

		notifications.MarkNotificationAsRead(ctx, db)
	})
	router.DELETE("/api/notifications/:notificationId", func(ctx *gin.Context) {
		notifications.DeleteNotification(ctx, db)
	})

	// Preferences Management
	router.GET("/api/notifications/preferences/:userId", func(ctx *gin.Context) {
		preferences.GetUserPreferences(ctx, db)
	})
	router.PUT("/api/notifications/preferences/:userId", func(ctx *gin.Context) {
		preferences.UpdateUserPreferences(ctx, db)
	})

	router.GET("/api/notifications/notify", func(ctx *gin.Context) {
		notifychannels.SetupConsumerGroup(ctx)
	})

	router.GET("/api/notifications/consume/:userID", func(ctx *gin.Context) {
		consumer_service.HandleNotifications(ctx, store)
	})

	router.GET("/api/notifications/consume", func(ctx *gin.Context) {
		consumer_service.HandleAllNotifications(ctx, store)
	})

	fmt.Printf("Nofification service 📨 started at http://localhost%s\n",
		LaunchPort)

	if err := router.Run(LaunchPort); err != nil {
		log.Printf("failed to run the server: %v", err)
	}
}
