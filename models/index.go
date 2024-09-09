package models

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetSqlConnection() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		os.Getenv("PG_HOST"),   // Urutan variabel host
		os.Getenv("PG_USER"),   // Urutan variabel user
		os.Getenv("PG_PASS"),   // Urutan variabel password
		os.Getenv("PG_DB"),     // Urutan variabel dbname
		os.Getenv("PG_PORT"),   // Urutan variabel port
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
