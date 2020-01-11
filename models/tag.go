package models

type Tag struct {
	Model
	Name  string  `gorm:"column:name;not null"`
	Tasks []*Task `gorm:"many2many:users_tasks"`
}
