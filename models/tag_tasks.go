package models

import "errors"

type TagTask struct {
	TagID  uint `pg:",pk"`
	TaskID uint `pg:",pk"`
	TimeRecords
}

func (ctx *TagTask) Create() error {
	return GetDb().Insert(ctx)
}

func (ctx *TagTask) Read() error {
	return errors.New("many-to-many table cannot be read directly")
}

func (ctx *TagTask) Update() error {
	return errors.New("many-to-many table cannot be updated directly")
}

func (ctx *TagTask) Delete() error {
	return errors.New("many-to-many table cannot be deleted directly")
}
