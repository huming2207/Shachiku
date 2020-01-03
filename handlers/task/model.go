package task

import (
	"github.com/jinzhu/gorm"
	"shachiku/handlers/tag"
	"shachiku/handlers/user"
	"time"
)

type Task struct {
	gorm.Model
	Title     string    `gorm:"column:title;not null"`
	Location  string    `gorm:"column:location"`
	Owner     user.User `gorm:"foreignkey:OwnerID"`
	OwnerID   uint
	Attendees []*user.User `gorm:"many2many:users_tasks;"`
	StartAt   time.Time    `gorm:"column:start_at;not null"`
	EndAt     time.Time    `gorm:"column:end_at"`
	Comment   string       `gorm:"column:comment"`
	Tags      []*tag.Tag   `gorm:"many2many:users_tags"`
}
