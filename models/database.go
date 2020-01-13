package models

import (
	"fmt"
	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
	"log"
	"shachiku/common"
	"time"
)

type TimeRecords struct {
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt time.Time `pg:",soft_delete" json:"-"`
}

var db *pg.DB

func GetDb() *pg.DB {
	cfg := common.GetConfig().Section(common.DatabaseSection)

	if db != nil {
		return db
	} else {
		db = pg.Connect(&pg.Options{
			Addr: fmt.Sprintf("%s:%s",
				cfg.Key(common.DatabaseHost).String(),
				cfg.Key(common.DatabasePort).String()),
			User:            cfg.Key(common.DatabaseUser).String(),
			Password:        cfg.Key(common.DatabasePassword).String(),
			Database:        cfg.Key(common.DatabaseName).String(),
			ApplicationName: "",
			TLSConfig:       nil,
		})
	}

	// Check connection
	_, err := db.Exec("SELECT 1")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	err = db.CreateTable(&User{}, &orm.CreateTableOptions{IfNotExists: true})
	if err != nil {
		log.Fatalf("Failed to create User table: %v", err)
	}

	err = db.CreateTable(&Task{}, &orm.CreateTableOptions{IfNotExists: true})
	if err != nil {
		log.Fatalf("Failed to create Task table: %v", err)
	}

	err = db.CreateTable(&Tag{}, &orm.CreateTableOptions{IfNotExists: true})
	if err != nil {
		log.Fatalf("Failed to create Tag table: %v", err)
	}

	err = db.CreateTable(&Role{}, &orm.CreateTableOptions{IfNotExists: true})
	if err != nil {
		log.Fatalf("Failed to create Role table: %v", err)
	}

	return db
}
