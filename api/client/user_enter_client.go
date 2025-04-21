package client

import (
	"common/proto/user_enter"
	"context"
)

func Register(ctx context.Context, in *user_enter.UserEnterRegisterRequest) (*user_enter.UserEnterRegisterResponse, error) {
	clients, err := UserEnterClients(ctx, func(ctx context.Context, server user_enter.UserEnterClient) (interface{}, error) {
		register, err := server.Register(ctx, in)
		if err != nil {
			return nil, err
		}
		return register, nil
	})
	if err != nil {
		return nil, err
	}
	return clients.(*user_enter.UserEnterRegisterResponse), nil
}
