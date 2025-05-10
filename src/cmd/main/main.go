package main

import (
	"context"
	"fmt"
	"log"

	consumer_service "notification-service/cmd/consumer"
	producer_service "notification-service/cmd/producer"
	"notification-service/pkg/models"

	"github.com/gin-gonic/gin"
)

const (
	LaunchPort = ":8081"
)

func main() {

	// Sample users for testing
	users := []models.User{
		{ID: 1, Name: "Pilathraj"},
		{ID: 2, Name: "Prakash"},
		{ID: 3, Name: "Gopika"},
		{ID: 4, Name: "Manasa"},
		{ID: 5, Name: "Himanshu"},
	}

	store := &consumer_service.NotificationStore{
		Data: make(consumer_service.UserNotifications),
	}

	producer, err := producer_service.SetupProducer()
	if err != nil {
		log.Fatalf("failed to initialize producer: %v", err)
	}
	defer producer.Close()

	ctx, cancel := context.WithCancel(context.Background())
	go consumer_service.SetupConsumerGroup(ctx, store)
	defer cancel()

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.POST("/send", producer_service.SendMessageHandler(producer, users))
	router.GET("/notifications/:userID", func(ctx *gin.Context) {
		consumer_service.HandleNotifications(ctx, store)
	})

	fmt.Printf("Nofification service 📨 started at http://localhost%s\n",
		LaunchPort)

	if err := router.Run(LaunchPort); err != nil {
		log.Printf("failed to run the server: %v", err)
	}
}
