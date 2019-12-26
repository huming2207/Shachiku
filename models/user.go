package models

import (
	"errors"
	"github.com/alexedwards/argon2id"
	"github.com/jinzhu/gorm"
	"shachiku/common"
)

type User struct {
	gorm.Model
	Username string  `gorm:"column:username"`
	Email    string  `gorm:"column:email;unique_index"`
	Bio      string  `gorm:"column:bio;size:1024"`
	Image    *string `gorm:"column:image"`
	Password string  `gorm:"column:password;not null"`
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
	return db.Save(&ctx)
}

func FindOneUser(query interface{}) (User, error) {
	var user User
	db := common.GetDb()
	err := db.Where(query).First(&user).Error
	return user, err
}

func DeleteUser(query interface{}) error {
	return common.GetDb().Where(query).Delete(User{}).Error
}