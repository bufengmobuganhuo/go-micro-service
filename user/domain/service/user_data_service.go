package service

import (
	"errors"
	"github.com/bufengmobuganhuo/go-micro-service/user/domain/model"
	"github.com/bufengmobuganhuo/go-micro-service/user/domain/repository"
	"golang.org/x/crypto/bcrypt"
)

type IUserDataService interface {
	AddUser(user *model.User) (userId int64, err error)
	DeleteUser(userId int64) error
	UpdateUser(user *model.User, isChangePwd bool) error
	FindUserByName(string) (*model.User, error)
	CheckPwd(userName string, pwd string) (isOk bool, err error)
}

type UserDataService struct {
	UserRepository repository.IUserRepository
}

func NewUserDataService(userRepository repository.IUserRepository) IUserDataService {
	return &UserDataService{UserRepository: userRepository}
}

func GenerateUserPassword(userPassword string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
}

func ValidateUserPassword(userPassword string, hashed string) (isOk bool, err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(hashed), []byte(userPassword)); err != nil {
		return false, errors.New("密码比对错误")
	}
	return true, nil
}

func (u UserDataService) AddUser(user *model.User) (userId int64, err error) {
	pwdByte, err := GenerateUserPassword(user.HashPassword)
	if err != nil {
		return user.ID, err
	}
	user.HashPassword = string(pwdByte)
	return u.UserRepository.CreateUser(user)
}

func (u UserDataService) DeleteUser(userId int64) error {
	return u.UserRepository.DeleteUserByID(userId)
}

func (u UserDataService) UpdateUser(user *model.User, isChangePwd bool) error {
	if isChangePwd {
		pwdByte, err := GenerateUserPassword(user.HashPassword)
		if err != nil {
			return err
		}
		user.HashPassword = string(pwdByte)
	}
	return u.UserRepository.UpdateUser(user)
}

func (u UserDataService) FindUserByName(username string) (*model.User, error) {
	return u.FindUserByName(username)
}

func (u UserDataService) CheckPwd(userName string, pwd string) (isOk bool, err error) {
	user, err := u.UserRepository.FindUserByName(userName)
	if err != nil {
		return false, err
	}
	return ValidateUserPassword(pwd, user.HashPassword)
}
