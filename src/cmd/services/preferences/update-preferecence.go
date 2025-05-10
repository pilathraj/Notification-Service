package preferences

import (
	"net/http"
	"notification-service/pkg/models"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UpdateUserPreferences(ctx *gin.Context, db *gorm.DB) {
	userID := ctx.Param("userId")
	var input models.UserPreferences
	var preference models.UserPreferences
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := db.Where("user_id = ?", userID).First(&preference).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			// Create a new record if it doesn't exist
			preference = models.UserPreferences{
				UserID:             userID,
				DueReminders:       input.DueReminders,
				OverdueNotices:     input.OverdueNotices,
				ReservationNotices: input.ReservationNotices,
				FineNotices:        input.FineNotices,
				PreferredChannels:  input.PreferredChannels,
				UpdatedAt:          time.Now(),
			}
			if err := db.Create(&preference).Error; err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create preferences"})
				return
			}
			ctx.JSON(http.StatusCreated, gin.H{"message": "Preferences created successfully"})
			return
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		}
		return
	}

	// Update preferences in the database
	if err := db.Model(&models.UserPreferences{}).Where("user_id = ?", userID).Updates(input).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update preferences"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Preferences updated successfully"})
}
