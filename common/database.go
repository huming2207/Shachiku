package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

var db *gorm.DB

func GetDb() *gorm.DB {
	var err error
	cfg := GetConfig().Section(DatabaseSection)
	dialect := cfg.Key(DatabaseDialect).String()

	if db != nil {
		return db
	} else if dialect == "sqlite3" {
		log.Println("Connecting to SQLite3 database")
		db, err = gorm.Open("sqlite3", cfg.Key(DatabasePath).String())
	} else if dialect == "postgres" {
		log.Println("Connecting to PostgreSQL database")
		db, err = gorm.Open("postgres",
			fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
				cfg.Key(DatabaseHost).String(),
				cfg.Key(DatabasePort).String(),
				cfg.Key(DatabaseUser).String(),
				cfg.Key(DatabaseName).String(),
				cfg.Key(DatabasePassword).String()))
	} else if dialect == "mysql" {
		log.Println("Connecting to MySQL database")
		db, err = gorm.Open("mysql",
			fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
				cfg.Key(DatabaseUser).String(),
				cfg.Key(DatabasePassword).String(),
				cfg.Key(DatabaseHost).String(),
				cfg.Key(DatabasePort).String(),
				cfg.Key(DatabaseName).String()))
	}

	if err != nil {
		log.Fatalln("Failed to open database: %w", err)
	}

	return db
}
