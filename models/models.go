package models

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"os"
	"shachiku/common"
)



var db *gorm.DB

func GetDb() *gorm.DB {
	var err error
	dialect := os.Getenv(common.DatabaseDialect)

	if db != nil{
		return db
	} else if dialect == "sqlite3" {
		db, err = gorm.Open("sqlite3", common.DatabasePath)
	} else if dialect == "postgres" {
		db, err = gorm.Open("postgres",
			fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s",
				os.Getenv(common.DatabaseHost),
				os.Getenv(common.DatabasePort),
				os.Getenv(common.DatabaseUser),
				os.Getenv(common.DatabaseName),
				os.Getenv(common.DatabasePassword)))
	} else if dialect == "mysql" {
		db, err = gorm.Open("mysql",
			fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
				os.Getenv(common.DatabaseUser),
				os.Getenv(common.DatabasePassword),
				os.Getenv(common.DatabaseHost),
				os.Getenv(common.DatabasePort),
				os.Getenv(common.DatabaseName)))
	}

	if err != nil {
		log.Fatalln("Failed to open database: %w", err)
	}

	return db
}
