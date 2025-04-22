package client

import (
	"common/proto/product"
	"context"
)

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
