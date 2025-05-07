package logic

import (
	"common/proto/product"
	"context"
	"product_srv/internal/handler"
)

type ProductServer struct {
	product.UnimplementedProductServer
}

// TODO:用户参与砍价信息列表
func (p ProductServer) BargainUserList(ctx context.Context, req *product.BargainUserListRequest) (*product.BargainUserListResponse, error) {
	list, err := handler.BargainUserList(req)
	if err != nil {
		return nil, err
	}
	return list, nil
}

// TODO:砍价帮助记录列表
func (p ProductServer) BargainUserHelpList(ctx context.Context, req *product.BargainUserHelpListRequest) (*product.BargainUserHelpListResponse, error) {
	list, err := handler.BargainUserHelpList(req)
	if err != nil {
		return nil, err
	}
	return list, nil
}

// TODO:创建用户参与砍价接口
func (p ProductServer) BargainUserCreate(ctx context.Context, req *product.BargainUserCreateRequest) (*product.BargainUserCreateResponse, error) {
	create, err := handler.BargainUserCreate(req)
	if err != nil {
		return nil, err
	}
	return create, nil
}

// TODO:用户参与砍价信息详情
func (p ProductServer) BargainUserShow(ctx context.Context, req *product.BargainUserShowRequest) (*product.BargainUserShowResponse, error) {
	show, err := handler.BargainUserShow(req)
	if err != nil {
		return nil, err
	}
	return show, nil
}

// TODO:砍价帮助记录详情
func (p ProductServer) BargainUserHelpShow(ctx context.Context, req *product.BargainUserHelpShowRequest) (*product.BargainUserHelpShowResponse, error) {
	show, err := handler.BargainUserHelpShow(req)
	if err != nil {
		return nil, err
	}
	return show, nil
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

// GroupBuying TODO:用户发起拼团
func (p ProductServer) GroupBuying(ctx context.Context, in *product.GroupBuyingRequest) (*product.GroupBuyingResponse, error) {
	userLevelPowerList, err := handler.GroupBuying(in)
	if err != nil {
		return nil, err
	}
	return userLevelPowerList, nil
}

// JoinGroupBuying  TODO:用户参与拼团
func (p ProductServer) JoinGroupBuying(ctx context.Context, in *product.JoinGroupBuyingRequest) (*product.JoinGroupBuyingResponse, error) {
	userLevelPowerList, err := handler.JoinGroupBuying(in)
	if err != nil {
		return nil, err
	}
	return userLevelPowerList, nil
}
func (p ProductServer) GetCombinationInfo(ctx context.Context, in *product.GetCombinationInfoRequest) (*product.GetCombinationInfoResponse, error) {
	userLevelPowerList, err := handler.GetCombinationInfo(in)
	if err != nil {
		return nil, err
	}
	return userLevelPowerList, nil
}
