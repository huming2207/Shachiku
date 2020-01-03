package tag

import (
	"github.com/jinzhu/gorm"
	"shachiku/handlers/task"
)

type Tag struct {
	gorm.Model
	Name  string       `gorm:"column:name;not null"`
	Tasks []*task.Task `gorm:"many2many:users_tasks"`
}
