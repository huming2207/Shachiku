package models

import (
	"time"
)

type Task struct {
	Model
	Title     string    `gorm:"column:title;not null" json:"title"`
	Location  string    `gorm:"column:location" json:"location"`
	OwnerID   uint      `json:"-"`
	Owner     User      `gorm:"foreignkey:OwnerID" json:"owner"`
	Attendees []*User   `gorm:"many2many:users_tasks;" json:"attendees"`
	StartAt   time.Time `gorm:"column:start_at;not null" json:"start_at"`
	EndAt     time.Time `gorm:"column:end_at" json:"end_at"`
	Comment   string    `gorm:"column:comment" json:"comment"`
	Tags      []*Tag    `gorm:"many2many:users_tags" json:"tags"`
}

func (ctx *Task) LoadOwner() error {
	db := GetDb()
	return db.Model(ctx).Related(&ctx.Owner, "Owner").Error
}
