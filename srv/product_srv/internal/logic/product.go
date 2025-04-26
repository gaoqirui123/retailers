package logic

import (
	"common/proto/product"
	"context"
	"product_srv/internal/handler"
)

type ProductServer struct {
	product.UnimplementedProductServer
}

// 修改商品表是否砍价状态
func (p ProductServer) ProductUpdate(ctx context.Context, req *product.ProductUpdateRequest) (*product.ProductUpdateResponse, error) {
	update, err := handler.ProductUpdate(req)
	if err != nil {
		return nil, err
	}
	return update, nil
}

// 创建砍价商品信息
func (p ProductServer) BargainCreate(ctx context.Context, req *product.BargainCreateRequest) (*product.BargainCreateResponse, error) {
	create, err := handler.BargainCreate(req)
	if err != nil {
		return nil, err
	}
	return create, nil
}

// 修改砍价商品表是否删除
func (p ProductServer) BargainUpdate(ctx context.Context, req *product.BargainUpdateRequest) (*product.BargainUpdateResponse, error) {
	update, err := handler.BargainUpdate(req)
	if err != nil {
		return nil, err
	}
	return update, nil
}

// 砍价商品表详情
func (p ProductServer) BargainShow(ctx context.Context, req *product.BargainShowRequest) (*product.BargainShowResponse, error) {
	show, err := handler.BargainShow(req)
	if err != nil {
		return nil, err
	}
	return show, nil
}

// 砍价商品表列表
func (p ProductServer) BargainList(ctx context.Context, req *product.BargainListRequest) (*product.BargainListResponse, error) {
	list, err := handler.BargainList(req)
	if err != nil {
		return nil, err
	}
	return list, nil
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
		return nil, err
	}
	return userLevelPowerList, nil
}

// JoinGroupBuying  用户发起拼团
func (p ProductServer) JoinGroupBuying(ctx context.Context, in *product.JoinGroupBuyingRequest) (*product.JoinGroupBuyingResponse, error) {
	userLevelPowerList, err := handler.JoinGroupBuying(in)
	if err != nil {
		return nil, err
	}
	return userLevelPowerList, nil
}
