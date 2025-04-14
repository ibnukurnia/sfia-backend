package db

import (
	"fmt"
	"os"
	"strconv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDb() (*gorm.DB, error) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		dbHost,
		dbUser,
		dbPass,
		dbName,
		dbPort,
	)

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		return nil, err
	}

	isDebug, _ := strconv.ParseBool(os.Getenv("APP_DEBUG"))

	if isDebug {
		db = db.Debug()
	}

	return db, nil
}
