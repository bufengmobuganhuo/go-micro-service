package handler

import (
	"context"
	"github.com/bufengmobuganhuo/go-micro-service/user/domain/model"
	"github.com/bufengmobuganhuo/go-micro-service/user/domain/service"
	user "github.com/bufengmobuganhuo/go-micro-service/user/proto/user"
)

type User struct {
	UserDataService service.IUserDataService
}

func (u User) Register(ctx context.Context, request *user.UserRegisterRequest, response *user.UserRegisterResponse) error {
	userRegistry := &model.User{
		UserName:     request.UserName,
		FirstName:    request.FirstName,
		HashPassword: request.Pwd,
	}
	_, err := u.UserDataService.AddUser(userRegistry)
	if err != nil {
		return err
	}
	response.Message = "添加成功"
	return nil
}

func (u User) Login(ctx context.Context, request *user.UserLoginRequest, response *user.UserLoginResponse) error {
	isOk, err := u.UserDataService.CheckPwd(request.UserName, request.Pwd)
	if err != nil {
		return err
	}
	response.IsSuccess = isOk
	return nil
}

func (u User) GetUserInfo(ctx context.Context, request *user.UserInfoRequest, response *user.UserInfoResponse) error {
	user, err := u.UserDataService.FindUserByName(request.UserName)
	if err != nil {
		return err
	}
	response = UserForResponse(user)
	return nil
}

func UserForResponse(userModel *model.User) *user.UserInfoResponse {
	response := &user.UserInfoResponse{}
	response.UserId = userModel.ID
	response.FirstName = userModel.FirstName
	response.UserName = userModel.UserName
	return response
}
