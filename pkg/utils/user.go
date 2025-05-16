package utils

import (
	"BlogHub/pkg/repo"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserID(c *gin.Context) uint {
	id, exists := c.Get("user_id")
	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return 0
	}

	userID, ok := id.(uint)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return 0
	}

	valid := repo.CheckForValidUser(userID)
	
	if !valid {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid or unauthorized user",
		})
		return 0
	}

	return userID
}
