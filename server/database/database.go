package database

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/CavasCallahan/firstGo/server/database/migrations"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {
	dts := url.URL{
		User:   url.UserPassword("postgres", "password"),
		Scheme: "postgres",
		Host:   fmt.Sprintf("%s:%d", "localhost", 5432),
		Path:   "spacedb",
	}

	database, err := gorm.Open(postgres.Open(dts.String()), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db = database

	migrations.RunMigrations(db)

	config, _ := db.DB()

	config.SetConnMaxIdleTime(10)
	config.SetMaxIdleConns(100)
	config.SetConnMaxLifetime(time.Hour)
}

func GetDataBase() *gorm.DB {
	return db
}
