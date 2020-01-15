package models

import "time"

type TimeRecords struct {
	CreatedAt time.Time `pg:"default:now()" json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt time.Time `pg:",soft_delete" json:"-"`
}

type CrudModel interface {
	Create() error
	Read() error
	Update() error
	Delete() error
}
