package initializer

import (
	"fmt"

	"github.com/Shakezidin/Viper-go/config"
	"github.com/Shakezidin/Viper-go/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseConnection(config config.Config) *gorm.DB {
	host := config.Db.Host
	user := config.Db.User
	password := config.Db.Password
	dbname := config.Db.Database
	port := config.Db.Port
	sslmode := config.Db.Sslmode
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbname, port, sslmode)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Connection to the database failed:", err)
	}

	DB.AutoMigrate(models.User{})
	return DB
}
