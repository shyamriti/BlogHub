package controllers

import (
	"BlogHub/pkg/dto"
	"BlogHub/pkg/repo"
	"BlogHub/pkg/services"
	"BlogHub/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var reqUser dto.RegisterUserRequest
	if err := c.ShouldBindJSON(&reqUser); err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "Invalid data", "details": err.Error()})
		return
	}

	exist, _ := repo.IsUserExist(reqUser.Email)

	if exist {
		c.AbortWithStatusJSON(500, gin.H{"error": "User already exist"})
		return
	}

	hashedPassword, err := utils.HashPassword(reqUser.Password)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"error": "Password encryption failed"})
		return
	}
	reqUser.Password = hashedPassword

	respUser, err := services.RegisterUserService(reqUser)

	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"dberror": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created",
		"details": respUser,
	})
}

func Login(c *gin.Context) {
	var reqUser dto.LoginUserRequest

	if err := c.ShouldBindJSON(&reqUser); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Invalid data", "details": err.Error()})
		return
	}

	token, err := services.LoginUserService(reqUser)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}

func Logout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Logout successful.",
	})
}
