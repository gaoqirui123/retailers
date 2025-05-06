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

// AddUsePower 用户使用权益
func AddUsePower(ctx context.Context, req *user.AddUsePowerRequest) (*user.AddUsePowerResponse, error) {
	return UserClients(ctx, req, func(ctx context.Context, client user.UserClient, req *user.AddUsePowerRequest) (*user.AddUsePowerResponse, error) {
		return client.AddUsePower(ctx, req)
	})
}

// UsePowerList 用户使用权益表展示
func UsePowerList(ctx context.Context, req *user.UsePowerListRequest) (*user.UsePowerListResponse, error) {
	return UserClients(ctx, req, func(ctx context.Context, client user.UserClient, req *user.UsePowerListRequest) (*user.UsePowerListResponse, error) {
		return client.UsePowerList(ctx, req)
	})
}

// AddText 用户使用权益表展示
func AddText(ctx context.Context, req *user.AddTextRequest) (*user.AddTextResponse, error) {
	return UserClients(ctx, req, func(ctx context.Context, client user.UserClient, req *user.AddTextRequest) (*user.AddTextResponse, error) {
		return client.AddText(ctx, req)
	})
}

// AddUserAddress 用户添加地址
func AddUserAddress(ctx context.Context, req *user.AddUserAddressRequest) (*user.AddUserAddressResponse, error) {
	return UserClients(ctx, req, func(ctx context.Context, client user.UserClient, req *user.AddUserAddressRequest) (*user.AddUserAddressResponse, error) {
		return client.AddUserAddress(ctx, req)
	})
}

// UserSignIn 用户签到
func UserSignIn(ctx context.Context, req *user.UserSignInRequest) (*user.UserSignInResponse, error) {
	return UserClients(ctx, req, func(ctx context.Context, client user.UserClient, req *user.UserSignInRequest) (*user.UserSignInResponse, error) {
		return client.UserSignIn(ctx, req)
	})
}

// UserMakeupSignIn 用户补签
func UserMakeupSignIn(ctx context.Context, req *user.UserMakeupSignInRequest) (*user.UserMakeupSignInResponse, error) {
	return UserClients(ctx, req, func(ctx context.Context, client user.UserClient, req *user.UserMakeupSignInRequest) (*user.UserMakeupSignInResponse, error) {
		return client.UserMakeupSignIn(ctx, req)
	})
}

// UserApplication 用户申请发票
func UserApplication(ctx context.Context, req *user.UserApplicationRequest) (*user.UserApplicationResponse, error) {
	return UserClients(ctx, req, func(ctx context.Context, client user.UserClient, req *user.UserApplicationRequest) (*user.UserApplicationResponse, error) {
		return client.UserApplication(ctx, req)
	})
}

// UpdatedAddress 用户修改地址
func UpdatedAddress(ctx context.Context, req *user.UpdatedAddressRequest) (*user.UpdatedAddressResponse, error) {
	return UserClients(ctx, req, func(ctx context.Context, client user.UserClient, req *user.UpdatedAddressRequest) (*user.UpdatedAddressResponse, error) {
		return client.UpdatedAddress(ctx, req)
	})
}

// UserReceiveCoupon 用户申请发票
func UserReceiveCoupon(ctx context.Context, req *user.UserReceiveCouponRequest) (*user.UserReceiveCouponResponse, error) {
	return UserClients(ctx, req, func(ctx context.Context, client user.UserClient, req *user.UserReceiveCouponRequest) (*user.UserReceiveCouponResponse, error) {
		return client.UserReceiveCoupon(ctx, req)
	})
}

// UserReceiveCoupon 用户提现
func UserWithdraw(ctx context.Context, req *user.UserWithdrawRequest) (*user.UserWithdrawResponse, error) {
	return UserClients(ctx, req, func(ctx context.Context, client user.UserClient, req *user.UserWithdrawRequest) (*user.UserWithdrawResponse, error) {
		return client.UserWithdraw(ctx, req)
	})
}

// UserAddressList 用户地址列表
func UserAddressList(ctx context.Context, req *user.UserAddressListRequest) (*user.UserAddressListResponse, error) {
	return UserClients(ctx, req, func(ctx context.Context, client user.UserClient, req *user.UserAddressListRequest) (*user.UserAddressListResponse, error) {
		return client.UserAddressList(ctx, req)
	})
}
