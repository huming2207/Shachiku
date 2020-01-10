package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"shachiku/common"
)

var db *gorm.DB

func GetDb() *gorm.DB {
	var err error
	cfg := common.GetConfig().Section(common.DatabaseSection)
	dialect := cfg.Key(common.DatabaseDialect).String()

	if db != nil {
		return db
	} else if dialect == "sqlite3" {
		log.Println("Connecting to SQLite3 database")
		db, err = gorm.Open("sqlite3", cfg.Key(common.DatabasePath).String())
	} else if dialect == "postgres" {
		log.Println("Connecting to PostgreSQL database")
		db, err = gorm.Open("postgres",
			fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
				cfg.Key(common.DatabaseHost).String(),
				cfg.Key(common.DatabasePort).String(),
				cfg.Key(common.DatabaseUser).String(),
				cfg.Key(common.DatabaseName).String(),
				cfg.Key(common.DatabasePassword).String()))
	} else if dialect == "mysql" {
		log.Println("Connecting to MySQL database")
		db, err = gorm.Open("mysql",
			fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
				cfg.Key(common.DatabaseUser).String(),
				cfg.Key(common.DatabasePassword).String(),
				cfg.Key(common.DatabaseHost).String(),
				cfg.Key(common.DatabasePort).String(),
				cfg.Key(common.DatabaseName).String()))
	}

	if err != nil {
		log.Fatalln("Failed to open database: %w", err)
	}

	db.CreateTable(&User{})
	db.CreateTable(&Tag{})
	db.CreateTable(&Task{})

	return db
}
