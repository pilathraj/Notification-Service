package preferences

import (
	"net/http"
	"notification-service/src/pkg/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetUserPreferences(ctx *gin.Context, db *gorm.DB) {
	userID := ctx.Param("userId")
	var preference models.UserPreferences
	if err := db.Where("user_id = ?", userID).First(&preference).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "preferences not found"})
		return
	}
	ctx.JSON(http.StatusOK, preference)
}
