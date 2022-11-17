package repository

import (
	"github.com/bufengmobuganhuo/go-micro-service/user/domain/model"
	"github.com/jinzhu/gorm"
)

type IUserRepository interface {
	InitTable() error
	FindAll() ([]model.User, error)
	FindUserByName(string) (*model.User, error)
	FindUserByID(int64) (*model.User, error)
	CreateUser(*model.User) (userId int64, err error)
	DeleteUserByID(int64) error
	UpdateUser(*model.User) error
}

type UseRepository struct {
	mysqlDb *gorm.DB
}

func (u *UseRepository) FindAll() (userAll []model.User, err error) {
	return userAll, u.mysqlDb.Find(&userAll).Error
}

func (u *UseRepository) FindUserByName(name string) (*model.User, error) {
	user := &model.User{}
	return user, u.mysqlDb.Where("user_name = ?", name).Find(user).Error
}

func (u *UseRepository) FindUserByID(id int64) (*model.User, error) {
	user := &model.User{}
	return user, u.mysqlDb.Where("id = ?", id).Find(user).Error
}

func (u *UseRepository) CreateUser(user *model.User) (userId int64, err error) {
	return user.ID, u.mysqlDb.Create(user).Error
}

func (u *UseRepository) DeleteUserByID(id int64) error {
	return u.mysqlDb.Where("id = ?", id).Delete(&model.User{}).Error
}

func (u *UseRepository) UpdateUser(user *model.User) error {
	return u.mysqlDb.Model(user).Update(&user).Error
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UseRepository{mysqlDb: db}
}

func (u *UseRepository) InitTable() error {
	return u.mysqlDb.Create(&model.User{}).Error
}
