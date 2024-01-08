package config

import (
	"github.com/IsaqueAmorim/DevHunter/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDataConnection() (*gorm.DB, error) {

	logger := GetLogger("Postgres")

	dsn := "host=db user=root password=root dbname=DEV_HUNTER port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{IgnoreRelationshipsWhenMigrating: false})
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	return db, nil
}
