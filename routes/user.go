package routes

import (
	"github.com/Shakezidin/Viper-go/controllers"
	"github.com/gin-gonic/gin"
)

func User(c *gin.Engine) {
	c.POST("user/signup", controllers.UserSignup)
	c.POST("user/login", controllers.UserLogin)
}
