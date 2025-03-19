package config

import (
	"log"
	"time"

	"github.com/Whitfrost21/EVault/evault/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Initdatabse() {
	dsn := "host=localhost user=api1 password=evault dbname=evault_database port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	var err error

	newlogger := logger.New(
		log.New(log.Writer(), "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Warn,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)
	models.Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newlogger})
	if err != nil {
		log.Fatalf("error while opening database : %v", err)
	}
	log.Println("database connected successfully!", models.Db)
}

func Migrateschema() {
	if err := models.Db.AutoMigrate(&models.Pickuprequest{}); err != nil {
		log.Fatalf("failed while migrating schema of pickup queue:%v", err)
	}
	if err := models.Db.AutoMigrate(&models.Collectedrequests{}); err != nil {
		log.Fatalf("failed while migrating schema of collected queue:%v", err)
	}
	if err := models.Db.AutoMigrate(&models.History{}); err != nil {
		log.Fatalf("failed while migrating schema of history:%v", err)
	}
	log.Println("migrated schema successfully")
}
