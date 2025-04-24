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
func UserClients[TRequest, TResponse any](ctx context.Context, request TRequest, operation func(ctx context.Context, client user.UserClient, req TRequest) (TResponse, error)) (TResponse, error) {
	return ExecuteGRPCOperation(ctx, "127.0.0.1:8081", NewUserClient, request, operation)
}

// UserLogin 用户登录
func UserLogin(ctx context.Context, in *user.UserLoginRequest) (*user.UserLoginResponse, error) {
	return UserClients(ctx, in, func(ctx context.Context, client user.UserClient, req *user.UserLoginRequest) (*user.UserLoginResponse, error) {
		return client.UserLogin(ctx, req)
	})
}

// UserRegister 用户注册
func UserRegister(ctx context.Context, req *user.UserRegisterRequest) (*user.UserRegisterResponse, error) {
	return UserClients(ctx, req, func(ctx context.Context, client user.UserClient, req *user.UserRegisterRequest) (*user.UserRegisterResponse, error) {
		return client.UserRegister(ctx, req)
	})
}

// UserDetail 获取用户详情
func UserDetail(ctx context.Context, req *user.UserDetailRequest) (*user.UserDetailResponse, error) {
	return UserClients(ctx, req, func(ctx context.Context, client user.UserClient, req *user.UserDetailRequest) (*user.UserDetailResponse, error) {
		return client.UserDetail(ctx, req)
	})
}

// ImproveUser 完善用户信息
func ImproveUser(ctx context.Context, req *user.ImproveUserRequest) (*user.ImproveUserResponse, error) {
	return UserClients(ctx, req, func(ctx context.Context, client user.UserClient, req *user.ImproveUserRequest) (*user.ImproveUserResponse, error) {
		return client.ImproveUser(ctx, req)
	})
}

// UpdatedPassword 更新用户密码
func UpdatedPassword(ctx context.Context, req *user.UpdatedPasswordRequest) (*user.UpdatedPasswordResponse, error) {
	return UserClients(ctx, req, func(ctx context.Context, client user.UserClient, req *user.UpdatedPasswordRequest) (*user.UpdatedPasswordResponse, error) {
		return client.UpdatedPassword(ctx, req)
	})
}

// UserLevelList 获取用户等级列表
func UserLevelList(ctx context.Context, req *user.UserLevelListRequest) (*user.UserLevelListResponse, error) {
	return UserClients(ctx, req, func(ctx context.Context, client user.UserClient, req *user.UserLevelListRequest) (*user.UserLevelListResponse, error) {
		return client.UserLevelList(ctx, req)
	})
}

// UserLevelPowerList 获取用户等级权限列表
func UserLevelPowerList(ctx context.Context, req *user.UserLevelPowerListRequest) (*user.UserLevelPowerListResponse, error) {
	return UserClients(ctx, req, func(ctx context.Context, client user.UserClient, req *user.UserLevelPowerListRequest) (*user.UserLevelPowerListResponse, error) {
		return client.UserLevelPowerList(ctx, req)
	})
}
