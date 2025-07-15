package helpers

import (
	"ecommerce-ums/internal/models"
	"fmt"
	"log"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetupMySQL() {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		GetEnv("DB_HOST", "127.0.0.1"),
		GetEnv("DB_USER", ""),
		GetEnv("DB_PASSWORD", ""),
		GetEnv("DB_NAME", ""),
		GetEnv("DB_PORT", "5432"),
	)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	logrus.Info("Connected to Postgres database success")

	DB.AutoMigrate(&models.User{}, &models.UserSession{})
}
