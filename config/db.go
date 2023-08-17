package config

import (
	"fmt"
	"os"

	"github.com/amikus123/cash-flowy/models"
	_ "github.com/joho/godotenv/autoload"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {

	dsn := fmt.Sprintf(
		"host=%s dbname=%s user=%s password=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&models.User{}, &models.ExpenseCategory{}, &models.User{})

	DB = db
}
