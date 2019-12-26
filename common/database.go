package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"sync"
)

type Database struct {
	*gorm.DB
}

var once sync.Once
var instance *Database

func GetDb() *Database {
	once.Do(func() {
		db, err := gorm.Open("sqlite3", "./../gorm.db")
		if err != nil {
			fmt.Println("db err: ", err)
		}
		instance = &Database{ db }
	})

	return instance
}