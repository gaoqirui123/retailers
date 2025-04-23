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
func UserEnterClients[TRequest, TResponse any](ctx context.Context, request TRequest, operation func(ctx context.Context, client user_enter.UserEnterClient, req TRequest) (TResponse, error)) (TResponse, error) {
	return ExecuteGRPCOperation(ctx, "127.0.0.1:8084", NewUserEnterClient, request, operation)
}

// Register 商户注册
func Register(ctx context.Context, in *user_enter.UserEnterRegisterRequest) (*user_enter.UserEnterRegisterResponse, error) {
	return UserEnterClients(ctx, in, func(ctx context.Context, client user_enter.UserEnterClient, req *user_enter.UserEnterRegisterRequest) (*user_enter.UserEnterRegisterResponse, error) {
		return client.Register(ctx, req)
	})
}

// AddProduct 添加商品
func AddProduct(ctx context.Context, in *user_enter.AddProductRequest) (*user_enter.AddProductResponse, error) {
	return UserEnterClients(ctx, in, func(ctx context.Context, client user_enter.UserEnterClient, req *user_enter.AddProductRequest) (*user_enter.AddProductResponse, error) {
		return client.AddProduct(ctx, req)
	})
}

// AddCombinationProduct 添加组合商品
func AddCombinationProduct(ctx context.Context, in *user_enter.AddCombinationProductRequest) (*user_enter.AddCombinationProductResponse, error) {
	return UserEnterClients(ctx, in, func(ctx context.Context, client user_enter.UserEnterClient, req *user_enter.AddCombinationProductRequest) (*user_enter.AddCombinationProductResponse, error) {
		return client.AddCombinationProduct(ctx, req)
	})
}
