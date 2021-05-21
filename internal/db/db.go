package db

import (
	"fmt"
	"os"

	"github.com/plyama/auth/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewPgGorm() (*gorm.DB, error) {
	return gorm.Open(postgres.Open(pgDSN()), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
}

func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&models.User{}, &models.Task{})
}

func pgDSN() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"), os.Getenv("DB_PORT"), os.Getenv("DB_SSLMODE"),
		os.Getenv("DB_TIMEZONE"))
}
