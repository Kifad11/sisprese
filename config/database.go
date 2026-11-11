package config

import (
	"fmt"
	"os"

	"sisprese/models"

	_ "modernc.org/sqlite"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// OpenDatabase opens the SQLite database and keeps its schema in sync with the models.
func OpenDatabase() (*gorm.DB, error) {
	databasePath := os.Getenv("DATABASE_PATH")
	if databasePath == "" {
		databasePath = "sisprese.db"
	}

	db, err := gorm.Open(sqlite.New(sqlite.Config{
		DriverName: "sqlite",
		DSN:        databasePath,
	}), &gorm.Config{TranslateError: true})
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}
	if err := db.AutoMigrate(&models.Employee{}, &models.Attendance{}); err != nil {
		return nil, fmt.Errorf("migrate database: %w", err)
	}
	return db, nil
}
