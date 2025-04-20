package logic

import (
	"common/proto/user"
	"context"
	"errors"
	"user_srv/internal/handler"
)

type UserServer struct {
	user.UnimplementedUserServer
}

func (u UserServer) UserLogin(ctx context.Context, in *user.UserLoginRequest) (*user.UserLoginResponse, error) {
	login, err := handler.UserLogin(in)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return login, nil
}

func (u UserServer) UserRegister(ctx context.Context, in *user.UserRegisterRequest) (*user.UserRegisterResponse, error) {
	register, err := handler.UserRegister(in)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return register, nil
}
