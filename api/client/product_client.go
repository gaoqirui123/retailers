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
func ProductClients[TRequest, TResponse any](ctx context.Context, request TRequest, operation func(ctx context.Context, client product.ProductClient, in TRequest) (TResponse, error)) (TResponse, error) {
	return ExecuteGRPCOperation(ctx, "127.0.0.1:8082", NewProductClient, request, operation)
}

// CombinationList 拼团商品列表展示
func CombinationList(ctx context.Context, in *product.CombinationListRequest) (*product.CombinationListResponse, error) {
	return ProductClients(ctx, in, func(ctx context.Context, client product.ProductClient, in *product.CombinationListRequest) (*product.CombinationListResponse, error) {
		return client.CombinationList(ctx, in)
	})
}

// GroupBuying 发起拼团
func GroupBuying(ctx context.Context, in *product.GroupBuyingRequest) (*product.GroupBuyingResponse, error) {
	return ProductClients(ctx, in, func(ctx context.Context, client product.ProductClient, in *product.GroupBuyingRequest) (*product.GroupBuyingResponse, error) {
		return client.GroupBuying(ctx, in)
	})
}

// JoinGroupBuying 参与拼团
func JoinGroupBuying(ctx context.Context, in *product.JoinGroupBuyingRequest) (*product.JoinGroupBuyingResponse, error) {
	return ProductClients(ctx, in, func(ctx context.Context, client product.ProductClient, in *product.JoinGroupBuyingRequest) (*product.JoinGroupBuyingResponse, error) {
		return client.JoinGroupBuying(ctx, in)
	})
}

// 修改商品表是否砍价状态
func ProductUpdate(ctx context.Context, in *product.ProductUpdateRequest) (*product.ProductUpdateResponse, error) {
	return ProductClients(ctx, in, func(ctx context.Context, client product.ProductClient, in *product.ProductUpdateRequest) (*product.ProductUpdateResponse, error) {
		return client.ProductUpdate(ctx, in)
	})
}

// 创建砍价商品表
func BargainCreate(ctx context.Context, in *product.BargainCreateRequest) (*product.BargainCreateResponse, error) {
	return ProductClients(ctx, in, func(ctx context.Context, client product.ProductClient, req *product.BargainCreateRequest) (*product.BargainCreateResponse, error) {
		return client.BargainCreate(ctx, in)
	})
}

// 修改砍价商品表是否删除
func BargainUpdate(ctx context.Context, in *product.BargainUpdateRequest) (*product.BargainUpdateResponse, error) {
	return ProductClients(ctx, in, func(ctx context.Context, client product.ProductClient, req *product.BargainUpdateRequest) (*product.BargainUpdateResponse, error) {
		return client.BargainUpdate(ctx, in)
	})
}

// 砍价商品表详情
func BargainShow(ctx context.Context, in *product.BargainShowRequest) (*product.BargainShowResponse, error) {
	return ProductClients(ctx, in, func(ctx context.Context, client product.ProductClient, req *product.BargainShowRequest) (*product.BargainShowResponse, error) {
		return client.BargainShow(ctx, in)
	})
}

// 砍价商品表列表
func BargainList(ctx context.Context, in *product.BargainListRequest) (*product.BargainListResponse, error) {
	return ProductClients(ctx, in, func(ctx context.Context, client product.ProductClient, req *product.BargainListRequest) (*product.BargainListResponse, error) {
		return client.BargainList(ctx, in)
	})
}
