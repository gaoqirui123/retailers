package client

import (
	"common/proto/product"
	"context"
	"google.golang.org/grpc"
)

func NewProductClient(cc grpc.ClientConnInterface) product.ProductClient {
	return product.NewProductClient(cc)
}

// ProductClients 封装的商品服务客户端处理函数
func ProductClients(ctx context.Context, handler func(ctx context.Context, server product.ProductClient) (interface{}, error)) (interface{}, error) {
	return GenericClient(ctx, "127.0.0.1:8082", NewProductClient, handler)
}

func CombinationList(ctx context.Context, in *product.CombinationListRequest) (*product.CombinationListResponse, error) {
	clients, err := ProductClients(ctx, func(ctx context.Context, server product.ProductClient) (interface{}, error) {
		list, err := server.CombinationList(ctx, in)
		if err != nil {
			return nil, err
		}
		return list, nil
	})
	if err != nil {
		return nil, err
	}
	return clients.(*product.CombinationListResponse), nil
}
