package models

type Tag struct {
	Model
	Name  string  `gorm:"column:name;not null" json:"name"`
	Tasks []*Task `gorm:"many2many:tags_tasks" json:"tasks"`
}

func (ctx *Tag) LoadTasks() error {
	db := GetDb()
	return db.Preload("Tasks").First(ctx).Error
}
