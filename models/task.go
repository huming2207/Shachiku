package models

import (
	"time"
)

type Task struct {
	Model
	Title    string    `gorm:"column:title;not null" json:"title"`
	Location string    `gorm:"column:location" json:"location"`
	People   []*Role   `json:"people"`
	StartAt  time.Time `gorm:"column:start_at;not null" json:"start_at"`
	EndAt    time.Time `gorm:"column:end_at" json:"end_at"`
	Comment  string    `gorm:"column:comment" json:"comment"`
	Tags     []*Tag    `gorm:"many2many:tags_tasks" json:"tags"`
}

func (ctx *Task) LoadPeople() error {
	db := GetDb()
	return db.Preload("People.User").First(ctx).Error
}
