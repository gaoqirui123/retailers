package logic

import (
	"common/proto/product"
	"context"
	"product_srv/internal/handler"
)

type ProductServer struct {
	product.UnimplementedProductServer
}

// TODO:拼团商品列表展示
func (p ProductServer) CombinationList(ctx context.Context, in *product.CombinationListRequest) (*product.CombinationListResponse, error) {
	list, err := handler.CombinationList(in)
	if err != nil {
		return nil, err
	}
	return list, nil
}
