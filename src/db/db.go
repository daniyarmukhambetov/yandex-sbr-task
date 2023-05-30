package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"yandex-team.ru/bstask/config"
	"yandex-team.ru/bstask/models"
)

func InitDB(cfg *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.PgURL), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&models.Courier{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&models.Region{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&models.Order{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&models.OrderStatus{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(&models.OrderAssign{}); err != nil {
		return err
	}
	return nil
}
