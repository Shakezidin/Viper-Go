package main

import (
	"fmt"

	"github.com/Shakezidin/Viper-go/initializer"
)

func main() {
	config, err := initializer.Loadenv()
	if err != nil {
		fmt.Println(err)
		return
	}
	initializer.DatabaseConnection(config)
}
