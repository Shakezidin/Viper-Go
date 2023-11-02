package initializer

import (
	"github.com/Shakezidin/Viper-go/config"
	"github.com/spf13/viper"
)

func Loadenv() (config.Config, error) {
	vp := viper.New()

	var config config.Config

	// Viper Configuration
	vp.AddConfigPath(".")
	vp.SetConfigName("test")
	vp.SetConfigType("json")
	if err := vp.ReadInConfig(); err != nil {
		return config, err
	}
	if err := vp.Unmarshal(&config); err != nil {
		return config, err
	}

	return config, nil
}
