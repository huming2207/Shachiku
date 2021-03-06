package models

import (
	"context"
	"fmt"
	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
	"log"
	"shachiku/common"
)

type dbLogger struct{}

func (d dbLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

func (d dbLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	log.Println(q.FormattedQuery())
	return nil
}

var pgDb *pg.DB

func GetDb() *pg.DB {
	cfg := common.GetConfig().Section(common.DatabaseSection)

	if pgDb != nil {
		return pgDb
	} else {
		pgDb = pg.Connect(&pg.Options{
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
	_, err := pgDb.Exec("SELECT 1")
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}

	err = pgDb.CreateTable(&User{}, &orm.CreateTableOptions{IfNotExists: true})
	if err != nil {
		log.Fatalf("Failed to create User table: %v", err)
	}

	err = pgDb.CreateTable(&Task{}, &orm.CreateTableOptions{IfNotExists: true})
	if err != nil {
		log.Fatalf("Failed to create Task table: %v", err)
	}

	err = pgDb.CreateTable(&Tag{}, &orm.CreateTableOptions{IfNotExists: true})
	if err != nil {
		log.Fatalf("Failed to create Tag table: %v", err)
	}

	err = pgDb.CreateTable(&TagTask{}, &orm.CreateTableOptions{IfNotExists: true})
	if err != nil {
		log.Fatalf("Failed to create TagTask table: %v", err)
	}

	err = pgDb.CreateTable(&Role{}, &orm.CreateTableOptions{IfNotExists: true})
	if err != nil {
		log.Fatalf("Failed to create Role table: %v", err)
	}

	pgDb.AddQueryHook(dbLogger{})
	return pgDb
}
