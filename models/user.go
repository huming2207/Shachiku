package models

import (
	"errors"
	"github.com/alexedwards/argon2id"
	"github.com/jinzhu/gorm"
)

type User struct {
	Model
	Username string  `gorm:"column:username;unique_index;not null" json:"username"`
	Email    string  `gorm:"column:email;unique_index;not null" json:"email"`
	Bio      string  `gorm:"column:bio;size:1024" json:"bio"`
	Image    *string `gorm:"column:image" json:"image"`
	Password string  `gorm:"column:password;not null" json:"-"` // No JSON operations allowed for password
	Tasks    []*Task `gorm:"many2many:users_tasks;" json:"tasks"`
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

func FindOneUser(query interface{}) (User, error) {
	var user User
	db := GetDb()
	err := db.Where(query).First(&user).Error
	return user, err
}

func DeleteUser(query interface{}) error {
	return GetDb().Where(query).Delete(User{}).Error
}
