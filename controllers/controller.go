package controllers

import (
	"net/http"

	"github.com/Shakezidin/Viper-go/initializer"
	"github.com/Shakezidin/Viper-go/models"
	"github.com/gin-gonic/gin"
)

// UserSignup function
func UserSignup(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Binding error",
		})
		return
	}
	if err := initializer.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "error while creating user"})
	}
	c.JSON(http.StatusOK, gin.H{"status": "user creation success"})
}

// User login Function
func UserLogin(c *gin.Context) {
	var userLogin models.Login
	var user models.User
	if err := c.ShouldBindJSON(&userLogin); err != nil {
		c.JSON(400, gin.H{"error": "binding error"})
		return
	}

	if err := initializer.DB.Where("email = ?", userLogin.Email).First(&user).Error; err != nil {
		c.JSON(400, gin.H{"Error": "user not found"})
		return
	}

	if userLogin.Password != user.Password {
		c.JSON(400, gin.H{"error": "password incorrect"})
		return
	}

	c.JSON(200, gin.H{"status ": "Login success"})
}
