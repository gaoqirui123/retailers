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

func (u UserServer) UserDetail(ctx context.Context, in *user.UserDetailRequest) (*user.UserDetailResponse, error) {
	userDetail, err := handler.UserDetail(in)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return userDetail, nil
}

func (u UserServer) ImproveUser(ctx context.Context, in *user.ImproveUserRequest) (*user.ImproveUserResponse, error) {
	improveUser, err := handler.ImproveUser(in)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return improveUser, nil
}

func (u UserServer) UpdatedPassword(ctx context.Context, in *user.UpdatedPasswordRequest) (*user.UpdatedPasswordResponse, error) {
	updatedPassword, err := handler.UpdatedPassword(in)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return updatedPassword, nil
}

func (u UserServer) UserLevelList(ctx context.Context, in *user.UserLevelListRequest) (*user.UserLevelListResponse, error) {
	userLevelList, err := handler.UserLevelList(in)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return userLevelList, nil
}

func (u UserServer) UserLevelPowerList(ctx context.Context, in *user.UserLevelPowerListRequest) (*user.UserLevelPowerListResponse, error) {
	userLevelPowerList, err := handler.UserLevelPowerList(in)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return userLevelPowerList, nil
}

func (u UserServer) AddUsePower(ctx context.Context, in *user.AddUsePowerRequest) (*user.AddUsePowerResponse, error) {
	addUsePower, err := handler.AddUsePower(in)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return addUsePower, nil
}

func (u UserServer) UsePowerList(ctx context.Context, in *user.UsePowerListRequest) (*user.UsePowerListResponse, error) {
	usePowerList, err := handler.UsePowerList(in)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return usePowerList, nil
}

func (u UserServer) UserSignIn(ctx context.Context, in *user.UserSignInRequest) (*user.UserSignInResponse, error) {
	sign, err := handler.UserSignIn(in)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return sign, nil
}

func (u UserServer) UserMakeupSignIn(ctx context.Context, in *user.UserMakeupSignInRequest) (*user.UserMakeupSignInResponse, error) {
	makeupSignIn, err := handler.UserMakeupSignIn(in)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return makeupSignIn, nil
}
