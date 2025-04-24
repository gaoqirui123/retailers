package client

import (
	administrators "common/proto/admin"
	"context"
	"google.golang.org/grpc"
)

// NewAdministratorsClient 创建管理员服务客户端
func NewAdministratorsClient(cc grpc.ClientConnInterface) administrators.AdministratorsClient {
	return administrators.NewAdministratorsClient(cc)
}

// AdministratorsClients 封装的管理员服务客户端处理函数
func AdministratorsClients[TRequest, TResponse any](ctx context.Context, request TRequest, operation func(ctx context.Context, client administrators.AdministratorsClient, req TRequest) (TResponse, error)) (TResponse, error) {
	return ExecuteGRPCOperation(ctx, "127.0.0.1:8086", NewAdministratorsClient, request, operation)
}

// AdminLogin 管理员登录
func AdminLogin(ctx context.Context, in *administrators.AdminLoginReq) (*administrators.AdminLoginResp, error) {
	return AdministratorsClients(ctx, in, func(ctx context.Context, client administrators.AdministratorsClient, req *administrators.AdminLoginReq) (*administrators.AdminLoginResp, error) {
		return client.AdminLogin(ctx, req)
	})
}

// ProcessEnter 处理商户入驻
func ProcessEnter(ctx context.Context, in *administrators.ProcessEnterReq) (*administrators.ProcessEnterResp, error) {
	return AdministratorsClients(ctx, in, func(ctx context.Context, client administrators.AdministratorsClient, req *administrators.ProcessEnterReq) (*administrators.ProcessEnterResp, error) {
		return client.ProcessEnter(ctx, req)
	})
}
