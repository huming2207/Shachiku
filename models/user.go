package models

import (
	"errors"
	"github.com/alexedwards/argon2id"
	"github.com/jinzhu/gorm"
)

type User struct {
	Model
	Username     string  `pg:"username,unique,notnull" json:"username"`
	Email        string  `pg:"email,unique,notnull" json:"email"`
	Bio          string  `pg:"bio" json:"bio"`
	Image        *string `pg:"image" json:"image"`
	Password     string  `pg:"password,notnull" json:"-"` // No JSON operations allowed for password
	RelatedTasks []*Role `json:"related_tasks,omitempty"`
}

func (ctx *User) SetPassword(pass string) (err error) {
	if len(pass) == 0 {
		return errors.New("password should not be empty")
	}

	ctx.Password, err = argon2id.CreateHash(pass, argon2id.DefaultParams)
	return err
}

func (ctx *User) CheckPassword(pass string) (match bool, err error) {
	return argon2id.ComparePasswordAndHash(pass, ctx.Password)
}

func (ctx *User) Update() *gorm.DB {
	return GetDb().Save(&ctx)
}

func (ctx *User) LoadRelatedTasks() error {
	return GetDb().Debug().
		Preload("RelatedTasks", "user_id = ?", ctx.ID).
		Preload("RelatedTasks.Task").First(ctx).Error
}

func FindOneUser(query interface{}) (User, error) {
	var user User
	db := GetDb()
	err := db.Where(query).First(&user).Error
	return user, err
}

func DeleteUser(query interface{}) error {
	return GetDb().Where(query).Delete(User{}).Error
}
