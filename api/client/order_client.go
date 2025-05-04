package client

import (
	"common/proto/order"
	"context"
	"google.golang.org/grpc"
)

func NewOrderClient(cc grpc.ClientConnInterface) order.OrderClient {
	return order.NewOrderClient(cc)
}

// OrderClients 封装订单操作的逻辑
func OrderClients[TRequest, TResponse any](ctx context.Context, request TRequest, operation func(ctx context.Context, client order.OrderClient, req TRequest) (TResponse, error)) (TResponse, error) {
	return ExecuteGRPCOperation(ctx, "127.0.0.1:8083", NewOrderClient, request, operation)
}

// AddOrder 用户创建订单
func AddOrder(ctx context.Context, in *order.AddOrderRequest) (*order.AddOrderResponse, error) {
	return OrderClients(ctx, in, func(ctx context.Context, client order.OrderClient, req *order.AddOrderRequest) (*order.AddOrderResponse, error) {
		return client.AddOrder(ctx, req)
	})
}

// PayCallback 支付回调
func PayCallback(ctx context.Context, in *order.PayCallbackRequest) (*order.PayCallbackResponse, error) {
	return OrderClients(ctx, in, func(ctx context.Context, client order.OrderClient, req *order.PayCallbackRequest) (*order.PayCallbackResponse, error) {
		return client.PayCallback(ctx, req)
	})
}

// OrderList 获取订单列表
func OrderList(ctx context.Context, in *order.OrderListRequest) (*order.OrderListResponse, error) {
	return OrderClients(ctx, in, func(ctx context.Context, client order.OrderClient, req *order.OrderListRequest) (*order.OrderListResponse, error) {
		return client.OrderList(ctx, req)
	})
}

// TODO:二维码核销
func QrCodeVerification(ctx context.Context, in *order.QrCodeVerificationRequest) (*order.QrCodeVerificationResponse, error) {
	return OrderClients(ctx, in, func(ctx context.Context, client order.OrderClient, req *order.QrCodeVerificationRequest) (*order.QrCodeVerificationResponse, error) {
		return client.QrCodeVerification(ctx, req)
	})
}
