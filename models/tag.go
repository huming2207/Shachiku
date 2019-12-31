package models

import "github.com/jinzhu/gorm"

type Tag struct {
	gorm.Model
	Name 	string
	Tasks 	[]*Task		`gorm:"many2many:users_tasks"`
}
