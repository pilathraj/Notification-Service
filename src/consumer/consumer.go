package consumer

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/IBM/sarama"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var KafkaServerAddress = os.Getenv("KAFKA_BROKERS")
var ConsumerTopic = os.Getenv("KAFKA_TOPIC")
var ConsumerGroup = os.Getenv("KAFKA_CONSUMER_GROUP")

// ============== HELPER FUNCTIONS ==============
var ErrNoMessagesFound = errors.New("no messages found")

func getUserIDFromRequest(ctx *gin.Context) (string, error) {
	userID := ctx.Param("userID")
	if userID == "" {
		return "", ErrNoMessagesFound
	}
	return userID, nil
}

type EventNotification map[string]interface{}

// ====== NOTIFICATION STORAGE ======
type UserNotifications map[string][]EventNotification

type NotificationStore struct {
	Data UserNotifications
	Mu   sync.RWMutex
}

func (ns *NotificationStore) Add(userID string,
	notification EventNotification) {
	ns.Mu.Lock()
	defer ns.Mu.Unlock()
	ns.Data[userID] = append(ns.Data[userID], notification)
}
func (ns *NotificationStore) GetAll() UserNotifications {
	ns.Mu.RLock()
	defer ns.Mu.RUnlock()
	return ns.Data
}

func (ns *NotificationStore) Get(userID string) []EventNotification {
	ns.Mu.RLock()
	defer ns.Mu.RUnlock()
	return ns.Data[userID]
}

// ============== KAFKA RELATED FUNCTIONS ==============
type Consumer struct {
	store *NotificationStore
}

func (*Consumer) Setup(sarama.ConsumerGroupSession) error   { return nil }
func (*Consumer) Cleanup(sarama.ConsumerGroupSession) error { return nil }

func (consumer *Consumer) ConsumeClaim(
	sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		userID := string(msg.Key)
		var notification EventNotification
		err := json.Unmarshal(msg.Value, &notification)
		if err != nil {
			log.Printf("failed to unmarshal notification: %v", err)
			continue
		}
		consumer.store.Add(userID, notification)
		sess.MarkMessage(msg, "")
	}
	return nil
}

func initializeConsumerGroup() (sarama.ConsumerGroup, error) {
	config := sarama.NewConfig()

	consumerGroup, err := sarama.NewConsumerGroup(
		[]string{KafkaServerAddress}, ConsumerGroup, config)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize consumer group: %w", err)
	}

	return consumerGroup, nil
}

func SetupConsumerGroup(ctx context.Context, store *NotificationStore) {
	consumerGroup, err := initializeConsumerGroup()
	if err != nil {
		log.Printf("initialization error: %v", err)
	}
	defer consumerGroup.Close()

	consumer := &Consumer{
		store: store,
	}

	for {
		err = consumerGroup.Consume(ctx, []string{ConsumerTopic}, consumer)
		if err != nil {
			log.Printf("error from consumer: %v", err)
		}
		if ctx.Err() != nil {
			return
		}
	}
}

func HandleAllNotifications(ctx *gin.Context, store *NotificationStore) {

	uuid := uuid.New()
	events := store.GetAll()
	if len(events) == 0 {
		ctx.JSON(http.StatusOK,
			gin.H{
				"messageId": uuid.String(),
				"message":   "No events found",
				"events":    []EventNotification{},
			})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"messageId": uuid.String(), "message": "Email has been sent successfully.", "events": events})
}

func HandleNotifications(ctx *gin.Context, store *NotificationStore) {
	userID, err := getUserIDFromRequest(ctx)
	uuid := uuid.New()
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"messageId": uuid.String(), "message": err.Error()})
		return
	}

	events := store.Get(userID)
	if len(events) == 0 {
		ctx.JSON(http.StatusOK,
			gin.H{
				"messageId": uuid.String(),
				"message":   "No events found",
				"events":    []EventNotification{},
			})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"messageId": uuid.String(), "message": "Email has been sent successfully.", "events": events})
}
