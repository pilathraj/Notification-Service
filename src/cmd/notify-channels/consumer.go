package notifychannels

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"notification-service/pkg/models"
	"os"
	"strings"
	"time"

	"github.com/segmentio/kafka-go"
)

var KafkaBroker = os.Getenv("KAFKA_BROKERS")
var KafkaTopic = os.Getenv("KAFKA_TOPIC")
var KafkaGroupID = os.Getenv("KAFKA_CONSUMER_GROUP")

func SetupConsumerGroup(ctx context.Context) {

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{KafkaBroker},
		Topic:   KafkaTopic,
		//GroupID:  KafkaGroupID,
		MaxBytes: 100,
	})
	defer reader.Close()

	readDeadline, _ := context.WithDeadline(context.Background(),
		time.Now().Add(5*time.Second))
	var notification models.Notification

	for {
		m, err := reader.ReadMessage(readDeadline)

		if err != nil {
			log.Printf("error reading message: %v", err)
			break
		}
		log.Printf("received message: %s", string(m.Value))
		// Parse the message
		err = json.Unmarshal(m.Value, &notification)
		if err != nil {
			fmt.Println("Error unmarshalling message:", err)
		}

		channel, message := notification.Channel, notification.Content

		channel = strings.ToLower(channel)

		log.Printf("channel: %s , message: %s", channel, message)

		// Handle the notification based on the channel
		switch channel {
		case "email":
			SendEmailNotification(message)
		case "sms":
			SendSMSNotification(message)
		case "push":
			SendPushNotification(message)
		default:
			log.Printf("unknown notification channel: %s", channel)
		}
	}
}
