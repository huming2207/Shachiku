package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Task struct {
	gorm.Model
	Title		string		`gorm:"column:title"`
	Location    string		`gorm:"column:location"`
	Owner		User		`gorm:"foreignkey:OwnerID"`
	OwnerID		uint
	Attendees   []*User 	`gorm:"many2many:user_task;"`
	StartAt     time.Time	`gorm:"column:start_at"`
	EndAt		time.Time	`gorm:"column:end_at"`
	Comment     string		`gorm:"column:comment"`
}
