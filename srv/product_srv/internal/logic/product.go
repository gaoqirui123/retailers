package logic

import (
	"common/proto/product"
	"context"
	"errors"
	"product_srv/internal/handler"
)

type ProductServer struct {
	product.UnimplementedProductServer
}

// CombinationList TODO:拼团商品列表展示
func (p ProductServer) CombinationList(ctx context.Context, in *product.CombinationListRequest) (*product.CombinationListResponse, error) {
	list, err := handler.CombinationList(in)
	if err != nil {
		return nil, err
	}
	return list, nil
}

// GroupBuying 用户发起拼团
func (p ProductServer) GroupBuying(ctx context.Context, in *product.GroupBuyingRequest) (*product.GroupBuyingResponse, error) {
	userLevelPowerList, err := handler.GroupBuying(in)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return userLevelPowerList, nil
}
