package models

import (
	"time"
)

type Task struct {
	ID uint `pg:"id,pk" json:"id"`
	TimeRecords
	Title    string    `pg:"title,notnull" json:"title"`
	Location string    `pg:"location" json:"location"`
	People   []*Role   `json:"people"`
	StartAt  time.Time `pg:"start_at,notnull" json:"start_at"`
	EndAt    time.Time `pg:"end_at" json:"end_at"`
	Comment  string    `pg:"comment" json:"comment"`
	Tags     []Tag     `pg:"many2many:tags_tasks" json:"tags"`
}

func (ctx *Task) LoadPeople() error {
	db := GetDb()
	err := db.Model(&ctx.People).Relation("User").First()
	return err
}
