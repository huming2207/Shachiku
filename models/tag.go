package models

type Tag struct {
	ID uint `pg:"id,pk"  json:"id"`
	TimeRecords
	Name  string  `pg:"name,notnull" json:"name"`
	Tasks []*Task `pg:"many2many:tags_tasks" json:"tasks"`
}

func (ctx *Tag) LoadTasks() error {
	db := GetDb()
	return db.Model(ctx).Column("tags.*").Relation("Tasks").Select()
}
