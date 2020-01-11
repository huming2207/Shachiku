package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Task struct {
	gorm.Model
	Title     string    `gorm:"column:title;not null" json:"title"`
	Location  string    `gorm:"column:location" json:"location"`
	Owner     User      `gorm:"foreignkey:OwnerID" json:"owner"`
	OwnerID   uint      `json:"-"`
	Attendees []*User   `gorm:"many2many:users_tasks;" json:"attendees"`
	StartAt   time.Time `gorm:"column:start_at;not null" json:"start_at"`
	EndAt     time.Time `gorm:"column:end_at" json:"end_at"`
	Comment   string    `gorm:"column:comment" json:"comment"`
	Tags      []*Tag    `gorm:"many2many:users_tags" json:"tags"`
}
