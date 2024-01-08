package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {

	var err error

	db, err = InitializeDataConnection()

	if err != nil {
		return fmt.Errorf("Error initializing sqlite: %v", err)
	}

	return nil
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}

func GetPostgress() *gorm.DB {
	return db
}
