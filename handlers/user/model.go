package user

import (
	"errors"
	"github.com/alexedwards/argon2id"
	"github.com/jinzhu/gorm"
	"shachiku/common"
	"shachiku/handlers/task"
)

type User struct {
	gorm.Model
	Username string       `gorm:"column:username;unique_index;not null"`
	Email    string       `gorm:"column:email;unique_index;not null"`
	Bio      string       `gorm:"column:bio;size:1024"`
	Image    *string      `gorm:"column:image"`
	Password string       `gorm:"column:password;not null"`
	Tasks    []*task.Task `gorm:"many2many:users_tasks;"`
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
	return common.GetDb().Save(&ctx)
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
