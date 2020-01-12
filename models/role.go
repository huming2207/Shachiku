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
	UserID uint      `gorm:"column:user_id;not null" json:"-"`
	User   *User     `json:"user,omitempty"`
	TaskID uint      `gorm:"column:task_id;not null" json:"-"`
	Task   *Task     `json:"task,omitempty"`
	Level  RoleLevel `gorm:"column:level" json:"role_level"`
}
