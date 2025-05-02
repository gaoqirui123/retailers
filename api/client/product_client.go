package client

import (
	"common/proto/product"
	"context"
	"google.golang.org/grpc"
)

// NewProductClient 创建管理员服务客户端
func NewProductClient(cc grpc.ClientConnInterface) product.ProductClient {
	return product.NewProductClient(cc)
}

// ProductClients 封装的商品服务客户端处理函数
func ProductClients[TRequest, TResponse any](ctx context.Context, request TRequest, operation func(ctx context.Context, client product.ProductClient, req TRequest) (TResponse, error)) (TResponse, error) {
	return ExecuteGRPCOperation(ctx, "127.0.0.1:8082", NewProductClient, request, operation)
}

// CombinationList 拼团商品列表展示
func CombinationList(ctx context.Context, in *product.CombinationListRequest) (*product.CombinationListResponse, error) {
	return ProductClients(ctx, in, func(ctx context.Context, client product.ProductClient, req *product.CombinationListRequest) (*product.CombinationListResponse, error) {
		return client.CombinationList(ctx, req)
	})
}

// GroupBuying 发起拼团
func GroupBuying(ctx context.Context, in *product.GroupBuyingRequest) (*product.GroupBuyingResponse, error) {
	return ProductClients(ctx, in, func(ctx context.Context, client product.ProductClient, req *product.GroupBuyingRequest) (*product.GroupBuyingResponse, error) {
		return client.GroupBuying(ctx, req)
	})
}

// JoinGroupBuying 参与拼团
func JoinGroupBuying(ctx context.Context, in *product.JoinGroupBuyingRequest) (*product.JoinGroupBuyingResponse, error) {
	return ProductClients(ctx, in, func(ctx context.Context, client product.ProductClient, req *product.JoinGroupBuyingRequest) (*product.JoinGroupBuyingResponse, error) {
		return client.JoinGroupBuying(ctx, req)
	})
}

// AddSeckillProduct 添加秒杀商品
func AddSeckillProduct(ctx context.Context, in *product.AddSeckillProductRequest) (*product.AddSeckillProductResponse, error) {
	return ProductClients(ctx, in, func(ctx context.Context, client product.ProductClient, req *product.AddSeckillProductRequest) (*product.AddSeckillProductResponse, error) {
		return client.AddSeckillProduct(ctx, req)
	})
}

// ReverseStock 秒杀后反还剩余的商品
func ReverseStock(ctx context.Context, in *product.ReverseStockRequest) (*product.ReverseStockResponse, error) {
	return ProductClients(ctx, in, func(ctx context.Context, client product.ProductClient, req *product.ReverseStockRequest) (*product.ReverseStockResponse, error) {
		return client.ReverseStock(ctx, req)
	})
}
