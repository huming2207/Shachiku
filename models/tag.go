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

func (ctx *Tag) Create() error {
	db := GetDb()
	_, err := db.Model(ctx).Returning("id").Insert()
	return err
}

func (ctx *Tag) Read() error {
	db := GetDb()
	err := db.Select(ctx)
	if err != nil {
		return err
	}

	return ctx.LoadTasks()
}

func (ctx *Tag) Update() error {
	db := GetDb()
	return db.Update(ctx)
}

func (ctx *Tag) Delete() error {
	db := GetDb()
	return db.Delete(ctx)
}
