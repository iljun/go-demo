package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"iljun.me/demo/pkg/config"
	"iljun.me/demo/pkg/dao"
	"iljun.me/demo/pkg/model"
)

type UserMySqlDao struct {
	mySqlClient *gorm.DB
}

func (dao UserMySqlDao) NewMySqlDao(config config.Config) dao.UserDao {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		config.MySqlUserName,
		config.MySqlPassword,
		config.MySqlURL,
		config.MySqlPort,
		config.MySqlDBName,
		)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn,
	}))

	if err != nil {
		panic("mysql database connection fail")
	}
	dao.mySqlClient = db
	return dao
}

func (dao UserMySqlDao) GetUser(id uint64) model.User {
	user := &model.User{}
	tx := dao.mySqlClient.First(user, id)
	if tx.Error != nil {
		return model.User{}
	}
	return *user
}

func (dao UserMySqlDao) SaveUser(user model.User) (model.User, error) {
	tx := dao.mySqlClient.Save(user)
	if tx.Error != nil {
		return model.User{}, tx.Error
	}
	return user, nil
}

func (dao UserMySqlDao) UpdateUser(id uint64, user model.User) (model.User, error) {
	model := model.User{}
	tx := dao.mySqlClient.First(model, id)
	if tx.Error != nil {
		return model, tx.Error
	}
	model.Age = user.Age
	model.Birthday = user.Birthday
	model.Name = user.Name
	model.Email = user.Email
	model.MemberNumber = user.MemberNumber

	updateTx := dao.mySqlClient.Save(model)
	if updateTx.Error != nil {
		return model, tx.Error
	}

	return model, nil
}

func (dao UserMySqlDao) DeleteUser(id uint64) error {
	user := &model.User{}

	tx := dao.mySqlClient.First(user, id)
	if tx.Error != nil {
		return tx.Error
	}

	deleteTx := dao.mySqlClient.Delete(user)
	if deleteTx.Error != nil {
		return deleteTx.Error
	}
	return nil
}
