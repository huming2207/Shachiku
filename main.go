package main

import (
	"log"
	"shachiku/common"
)

func main() {
	db := common.GetDb()
	log.Println(db)
}
