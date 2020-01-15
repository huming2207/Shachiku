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
	Tags     []*Tag    `pg:"many2many:tag_tasks" json:"tags"`
}

func (ctx *Task) LoadPeople() error {
	return GetDb().Model(&ctx.People).Relation("User").First()
}

func (ctx *Task) LoadTags() error {
	return GetDb().Model(ctx).
		Column("task.*").
		Where("task.id = ?", ctx.ID).
		Relation("Tags").First()
}

func (ctx *Task) Create() error {
	_, err := db.Model(ctx).Returning("id").Insert()
	if err != nil {
		return err
	}

	if ctx.People != nil {
		ctx.People[0].TaskID = ctx.ID // Assign the actual task ID
		err = db.Insert(ctx.People[0])
		if err != nil {
			return err
		}
	}

	return nil
}

func (ctx *Task) Read() error {
	db := GetDb()
	err := db.Select(ctx)
	if err != nil {
		return err
	}

	return ctx.LoadPeople()
}

func (ctx *Task) Update() error {
	db := GetDb()
	return db.Update(ctx)
}

func (ctx *Task) Delete() error {
	db := GetDb()
	return db.Delete(ctx)
}
