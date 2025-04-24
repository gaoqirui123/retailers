package client

import (
	"common/proto/cart"
	"context"
	"google.golang.org/grpc"
)

func NewCartClient(cc grpc.ClientConnInterface) cart.CartClient {
	return cart.NewCartClient(cc)
}

// CartClients 封装购物车操作的逻辑
func CartClients[TRequest, TResponse any](ctx context.Context, request TRequest, operation func(ctx context.Context, client cart.CartClient, req TRequest) (TResponse, error)) (TResponse, error) {
	return ExecuteGRPCOperation(ctx, "127.0.0.1:8087", NewCartClient, request, operation)
}

// AddCart 向购物车添加商品
func AddCart(ctx context.Context, in *cart.AddCartRequest) (*cart.AddCartResponse, error) {
	return CartClients(ctx, in, func(ctx context.Context, client cart.CartClient, req *cart.AddCartRequest) (*cart.AddCartResponse, error) {
		return client.AddCart(ctx, req)
	})
}

// ClearCart 清空购物车
func ClearCart(ctx context.Context, in *cart.ClearCartRequest) (*cart.ClearCartResponse, error) {
	return CartClients(ctx, in, func(ctx context.Context, client cart.CartClient, req *cart.ClearCartRequest) (*cart.ClearCartResponse, error) {
		return client.ClearCart(ctx, req)
	})
}

// DeleteCart 删除购物车中的商品
func DeleteCart(ctx context.Context, in *cart.DeleteCartRequest) (*cart.DeleteCartResponse, error) {
	return CartClients(ctx, in, func(ctx context.Context, client cart.CartClient, req *cart.DeleteCartRequest) (*cart.DeleteCartResponse, error) {
		return client.DeleteCart(ctx, req)
	})
}

// GetCartList 获取购物车列表
func GetCartList(ctx context.Context, in *cart.GetCartListRequest) (*cart.GetCartListResponse, error) {
	return CartClients(ctx, in, func(ctx context.Context, client cart.CartClient, req *cart.GetCartListRequest) (*cart.GetCartListResponse, error) {
		return client.GetCartList(ctx, req)
	})
}
