package client

import (
	"common/proto/user"
	"context"
)

func UserLogin(ctx context.Context, in *user.UserLoginRequest) (*user.UserLoginResponse, error) {
	client, err := UserClients(ctx, func(ctx context.Context, client user.UserClient) (interface{}, error) {
		login, err := client.UserLogin(ctx, in)
		if err != nil {
			return nil, err
		}
		return login, nil
	})
	if err != nil {
		return nil, err
	}
	return client.(*user.UserLoginResponse), nil
}
func UserRegister(ctx context.Context, req *user.UserRegisterRequest) (*user.UserRegisterResponse, error) {
	client, err := UserClients(ctx, func(ctx context.Context, client user.UserClient) (interface{}, error) {
		register, err := client.UserRegister(ctx, req)
		if err != nil {
			return nil, err
		}
		return register, nil
	})
	if err != nil {
		return nil, err
	}
	return client.(*user.UserRegisterResponse), nil
}
func UserDetail(ctx context.Context, req *user.UserDetailRequest) (*user.UserDetailResponse, error) {
	client, err := UserClients(ctx, func(ctx context.Context, client user.UserClient) (interface{}, error) {
		register, err := client.UserDetail(ctx, req)
		if err != nil {
			return nil, err
		}
		return register, nil
	})
	if err != nil {
		return nil, err
	}
	return client.(*user.UserDetailResponse), nil
}
