package main

import (
	"fmt"

	"github.com/Shakezidin/Viper-go/initializer"
	"github.com/Shakezidin/Viper-go/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	config, err := initializer.Loadenv()
	if err != nil {
		fmt.Println(err)
		return
	}
	initializer.DatabaseConnection(config)

	r := gin.Default()
	routes.User(r)

	r.Run("localhost:3000")
}
