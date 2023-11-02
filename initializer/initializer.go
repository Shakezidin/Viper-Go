package initializer

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

func Loadenv() {
	vp := viper.New()

	//viper Configuration
	vp.AddConfigPath(".")
	vp.SetConfigName("test")
	vp.SetConfigType("json")
	if err := vp.ReadInConfig(); err != nil {
		log.Fatal("error readinmg json env file", err)
	}

	fmt.Println(vp.GetString("MESSAGE"))
}