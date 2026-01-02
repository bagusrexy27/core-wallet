package config

import (
	"fmt"

	"github.com/core-wallet/app/models"
	"github.com/core-wallet/app/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		AppConfig.DBUser, AppConfig.DBPassword, AppConfig.DBHost, AppConfig.DBPort,
		AppConfig.DBName,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		utils.ErrorLog("INITIALIZE - APP failed to connect database", err, true)
		return err
	}

	DB = db
	utils.InfoLog("connected to database Postgres")

	autoMigrate()
	return nil
}

func autoMigrate() {
	err := DB.AutoMigrate(
		&models.Wallet{},
		&models.Transaction{},
		// &models.Attachment{},
		// &models.Anggaran{},
		// &models.AnggaranDumb{},
	)
	if err != nil {
		utils.ErrorLog("INITIALIZE - APP failed to migrate database", err, true)
	}

	utils.InfoLog("INITIALIZE - APP database migrated successfully")
}
