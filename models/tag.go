package models

type Tag struct {
	ID uint `pg:"id,pk"  json:"id"`
	TimeRecords
	Name  string  `pg:"name,notnull" json:"name"`
	Tasks []*Task `pg:"many2many:tag_tasks" json:"tasks"`
}

type TagTask struct {
	TagID  uint `pg:",pk"`
	TaskID uint `pg:",pk"`
	TimeRecords
}

func (ctx *Tag) LoadTasks() error {
	db := GetDb()
	return db.Model(ctx).Column("task.*").Relation("Tasks").Select()
}
