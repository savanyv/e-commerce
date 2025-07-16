package database

import (
	"errors"
	"fmt"

	"github.com/savanyv/e-commerce/config"
	"github.com/savanyv/e-commerce/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectPostgres(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.PG.HostDB,
		cfg.PG.UserDB,
		cfg.PG.PassDB,
		cfg.PG.NameDB,
		cfg.PG.PortDB,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, errors.New("error connecting to postgres")
	}

	if err := db.AutoMigrate(
		&models.Brand{},
		&models.Product{},
	); err != nil {
		return nil, errors.New("error migrating database")
	}

	DB = db
	return db, nil
}
