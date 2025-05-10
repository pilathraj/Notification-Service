package tests

import (
	"net/http"
	"net/http/httptest"
	"notification-service/src/consumer"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHandleNotifications(t *testing.T) {
	// Create a new NotificationStore
	store := &consumer.NotificationStore{
		Data: make(consumer.UserNotifications),
	}

	// Add test data to the store
	userID := "user123"
	notification := consumer.EventNotification{"event": "notification"}
	store.Add(userID, notification)

	// Create a new Gin context for testing
	router := gin.Default()
	router.GET("/notifications/:userID", func(ctx *gin.Context) {
		consumer.HandleNotifications(ctx, store)
	})

	// Test case 1: User with notifications
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/notifications/user123", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"messageId"`)
	assert.Contains(t, w.Body.String(), `"message"`)
	assert.Contains(t, w.Body.String(), `"events"`)

	// Test case 2: User without notifications
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/notifications/user456", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"messageId"`)
	assert.Contains(t, w.Body.String(), `"message"`)
	assert.Contains(t, w.Body.String(), `"events"`)
	assert.NotContains(t, w.Body.String(), `"event"`)
}
func TestHandleAllNotifications(t *testing.T) {
	// Create a new NotificationStore
	store := &consumer.NotificationStore{
		Data: make(consumer.UserNotifications),
	}

	// Add test data to the store
	userID := "user123"
	notification := consumer.EventNotification{"event": "notification"}
	store.Add(userID, notification)

	// Create a new Gin context for testing
	router := gin.Default()
	router.GET("/notifications", func(ctx *gin.Context) {
		consumer.HandleAllNotifications(ctx, store)
	})

	// Test case 1: No notifications in the store
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/notifications", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), `"messageId"`)
	assert.Contains(t, w.Body.String(), `"message"`)
	assert.Contains(t, w.Body.String(), `"events"`)

}
