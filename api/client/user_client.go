package client

import (
	"common/proto/user"
	"context"
	"google.golang.org/grpc"
)

func NewUserClient(cc grpc.ClientConnInterface) user.UserClient {
	return user.NewUserClient(cc)
}

// UserClients 封装的用户服务客户端处理函数
func UserClients(ctx context.Context, handler func(ctx context.Context, server user.UserClient) (interface{}, error)) (interface{}, error) {
	return GenericClient(ctx, "127.0.0.1:8081", NewUserClient, handler)
}

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
func ImproveUser(ctx context.Context, req *user.ImproveUserRequest) (*user.ImproveUserResponse, error) {
	client, err := UserClients(ctx, func(ctx context.Context, client user.UserClient) (interface{}, error) {
		register, err := client.ImproveUser(ctx, req)
		if err != nil {
			return nil, err
		}
		return register, nil
	})
	if err != nil {
		return nil, err
	}
	return client.(*user.ImproveUserResponse), nil
}
