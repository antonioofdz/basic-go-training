package database

import (
	"fmt"

	"github.com/juju/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"basic-go-training/internal/config"
	"basic-go-training/internal/domain/entities/patients"
)

var Repo *Database

type Database struct {
	gormRepo *gorm.DB
}

func Connect() error {
	connection := fmt.Sprintf(
		"user=%s password=%s host=%s port=%s dbname=%s sslmode=disable",
		config.Settings.Database.User,
		config.Settings.Database.Password,
		config.Settings.Database.Host,
		config.Settings.Database.Port,
		config.Settings.Database.Name,
	)
	gormRepo, err := gorm.Open(postgres.Open(connection), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return errors.Trace(err)
	}

	Repo = &Database{
		gormRepo: gormRepo,
	}

	if err := gormRepo.AutoMigrate(&patients.Patient{}); err != nil {
		return errors.Trace(err)
	}

	return nil
}

func (db *Database) AllPatients() *gorm.DB {
	return db.gormRepo.Table("patients")
}

func (db *Database) Patients(id int) *gorm.DB {
	return db.gormRepo.Table("patients").Where("id = ?", id)
}
