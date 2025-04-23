package client

import (
	"common/proto/user_enter"
	"context"
	"google.golang.org/grpc"
)

func NewUserEnterClient(cc grpc.ClientConnInterface) user_enter.UserEnterClient {
	return user_enter.NewUserEnterClient(cc)
}

// UserEnterClients 封装的商户服务客户端处理函数
func UserEnterClients(ctx context.Context, handler func(ctx context.Context, server user_enter.UserEnterClient) (interface{}, error)) (interface{}, error) {
	return GenericClient(ctx, "127.0.0.1:8084", NewUserEnterClient, handler)
}

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
func AddProduct(ctx context.Context, in *user_enter.AddProductRequest) (*user_enter.AddProductResponse, error) {
	clients, err := UserEnterClients(ctx, func(ctx context.Context, server user_enter.UserEnterClient) (interface{}, error) {
		register, err := server.AddProduct(ctx, in)
		if err != nil {
			return nil, err
		}
		return register, nil
	})
	if err != nil {
		return nil, err
	}
	return clients.(*user_enter.AddProductResponse), nil
}
