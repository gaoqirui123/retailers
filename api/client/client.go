package client

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
)

// GenericClient 通用的 gRPC 客户端处理函数
func GenericClient[T any](ctx context.Context, addr string, newClient func(cc grpc.ClientConnInterface) T, handler func(ctx context.Context, client T) (interface{}, error)) (interface{}, error) {
	// 连接 gRPC 服务
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("连接异常： %s\n", err)
		return nil, err
	}
	defer conn.Close()

	// 创建客户端
	client := newClient(conn)

	// 执行处理函数
	res, err := handler(ctx, client)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// ExecuteGRPCOperation 封装 gRPC 操作的通用逻辑
func ExecuteGRPCOperation[TClient any, TRequest any, TResponse any](ctx context.Context, addr string, newClient func(cc grpc.ClientConnInterface) TClient, request TRequest, operation func(ctx context.Context, client TClient, req TRequest) (TResponse, error)) (TResponse, error) {
	var zero TResponse
	res, err := GenericClient(ctx, addr, newClient, func(ctx context.Context, client TClient) (interface{}, error) {
		resp, err := operation(ctx, client, request)
		if err != nil {
			return nil, err
		}
		return resp, nil
	})
	if err != nil {
		return zero, err
	}
	return res.(TResponse), nil
}
