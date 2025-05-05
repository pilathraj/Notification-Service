package models

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

// Custom type for PreferredChannels
type StringArray []string

// Implement the `Value` method to convert the Go type to a JSON string for the database
func (s StringArray) Value() (driver.Value, error) {
	return json.Marshal(s)
}

// Implement the `Scan` method to convert a JSON string from the database to the Go type
func (s *StringArray) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(bytes, s)
}

type Notification struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    string    `json:"userId"`
	Subject   string    `json:"subject"`
	Content   string    `json:"content"`
	Type      string    `json:"type"`
	Channel   string    `json:"channel"`
	Email     string    `json:"email" option:"email"`
	MobileNo  string    `json:"mobileno" option:"mobileno"`
	IsRead    bool      `json:"isRead"`
	CreatedAt time.Time `json:"createdAt"`
	ReadAt    time.Time `json:"readAt"`
}

type UserPreferences struct {
	ID                 uint        `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID             string      `json:"userId"`
	DueReminders       bool        `json:"dueReminders"`
	OverdueNotices     bool        `json:"overdueNotices"`
	ReservationNotices bool        `json:"reservationNotices"`
	FineNotices        bool        `json:"fineNotices"`
	PreferredChannels  StringArray `gorm:"type:json" json:"preferredChannels"`
	UpdatedAt          time.Time   `json:"updatedAt"`
}

// type Preference struct {
// 	ID     uint `gorm:"primaryKey" json:"id"`
// 	UserID uint `json:"user_id"`
// 	Email  bool `json:"email"`
// 	SMS    bool `json:"sms"`
// 	Push   bool `json:"push"`
// }
