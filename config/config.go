package config

type DbConfig struct {
	Host     string `mapstructure:"HOST"`
	User     string `mapstructure:"USER"`
	Password string `mapstructure:"PASSWORD"`
	Database string `mapstructure:"DATABASE"`
	Port     string `mapstructure:"PORT"`
	Sslmode  string `mapstructure:"SSLMODE"`
}

type Config struct {
	Db DbConfig `mapstructure:"db"`
}