package models

type RoleLevel uint

const (
	Admin    RoleLevel = 0
	Owner    RoleLevel = 1
	Attendee RoleLevel = 2
)

// Usage: refer to https://github.com/jinzhu/gorm/issues/719
type Role struct {
	Model
	UserID uint      `pg:"user_id,pk" json:"-"`
	User   *User     `json:"user,omitempty"`
	TaskID uint      `pg:"task_id,pk" json:"-"`
	Task   *Task     `json:"task,omitempty"`
	Level  RoleLevel `pg:"level,type:smallint,notnull" json:"role_level"`
}
