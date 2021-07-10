package dao

import (
	"iljun.me/demo/pkg/config"
	"iljun.me/demo/pkg/dao/mysql"
	"iljun.me/demo/pkg/model"
)

type UserDao interface {
	GetUser(id uint64) model.User
	SaveUser(user model.User) (model.User, error)
	UpdateUser(id uint64, user model.User) (model.User, error)
	DeleteUser(id uint64) error
}

func NewUserDao(config config.Config) UserDao {
	return mysql.UserMySqlDao{}.NewMySqlDao(config)
}