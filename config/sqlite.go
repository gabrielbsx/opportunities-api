package config

import (
	"os"

	"github.com/gabrielbsx/opportunities-api/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")

	dbPath := "./db/opportunities.db"

	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		logger.Info("Database file not found, creating a new one")

		err = os.MkdirAll("./db", os.ModePerm)

		if err != nil {
			logger.Error("Failed to create database directory")
			return nil, err
		}

		file, err := os.Create(dbPath)

		if err != nil {
			logger.Error("Failed to create database file")
			return nil, err
		}

		file.Close()

		logger.Info("Database file created")
	}

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		logger.Error("Failed to connect to database")
		return nil, err
	}

	logger.Info("Connected to database")

	err = db.AutoMigrate(&schemas.Opening{})

	if err != nil {
		logger.Error("Failed to migrate database")
		return nil, err
	}

	logger.Info("Migrated database")

	return db, nil
}
