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
func UpdatedPassword(ctx context.Context, req *user.UpdatedPasswordRequest) (*user.UpdatedPasswordResponse, error) {
	client, err := UserClients(ctx, func(ctx context.Context, client user.UserClient) (interface{}, error) {
		updatedPassword, err := client.UpdatedPassword(ctx, req)
		if err != nil {
			return nil, err
		}
		return updatedPassword, nil
	})
	if err != nil {
		return nil, err
	}
	return client.(*user.UpdatedPasswordResponse), nil
}
func UserLevelList(ctx context.Context, req *user.UserLevelListRequest) (*user.UserLevelListResponse, error) {
	client, err := UserClients(ctx, func(ctx context.Context, client user.UserClient) (interface{}, error) {
		userLevelList, err := client.UserLevelList(ctx, req)
		if err != nil {
			return nil, err
		}
		return userLevelList, nil
	})
	if err != nil {
		return nil, err
	}
	return client.(*user.UserLevelListResponse), nil
}
func UserLevelPowerList(ctx context.Context, req *user.UserLevelPowerListRequest) (*user.UserLevelPowerListResponse, error) {
	client, err := UserClients(ctx, func(ctx context.Context, client user.UserClient) (interface{}, error) {
		userLevelPowerList, err := client.UserLevelPowerList(ctx, req)
		if err != nil {
			return nil, err
		}
		return userLevelPowerList, nil
	})
	if err != nil {
		return nil, err
	}
	return client.(*user.UserLevelPowerListResponse), nil
}
func GroupBuying(ctx context.Context, req *user.GroupBuyingRequest) (*user.GroupBuyingResponse, error) {
	client, err := UserClients(ctx, func(ctx context.Context, client user.UserClient) (interface{}, error) {
		userLevelPowerList, err := client.GroupBuying(ctx, req)
		if err != nil {
			return nil, err
		}
		return userLevelPowerList, nil
	})
	if err != nil {
		return nil, err
	}
	return client.(*user.GroupBuyingResponse), nil
}
