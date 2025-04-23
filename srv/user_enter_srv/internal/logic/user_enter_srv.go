package logic

import (
	"common/proto/user_enter"
	"context"
	"user_enter_srv/internal/handler"
)

type UserEnterServer struct {
	user_enter.UnimplementedUserEnterServer
}

// Register TODO:商户申请注册
func (es UserEnterServer) Register(ctx context.Context, in *user_enter.UserEnterRegisterRequest) (*user_enter.UserEnterRegisterResponse, error) {
	register, err := handler.Register(in)
	if err != nil {
		return nil, err
	}
	return register, nil
}

// AddProduct TODO:商户发布商品
func (es UserEnterServer) AddProduct(ctx context.Context, in *user_enter.AddProductRequest) (*user_enter.AddProductResponse, error) {
	register, err := handler.AddProduct(in)
	if err != nil {
		return nil, err
	}
	return register, nil
}

// AddCombinationProduct TODO:添加拼团商品
func (es UserEnterServer) AddCombinationProduct(ctx context.Context, in *user_enter.AddCombinationProductRequest) (*user_enter.AddCombinationProductResponse, error) {
	product, err := handler.AddCombinationProduct(in)
	if err != nil {
		return nil, err
	}
	return product, nil
}

// ProcessInvoice TODO:商家审核用户的发票invoice申请
func (es UserEnterServer) ProcessInvoice(ctx context.Context, in *user_enter.ProcessInvoiceRequest) (*user_enter.ProcessInvoiceResponse, error) {
	product, err := handler.ProcessInvoice(in)
	if err != nil {
		return nil, err
	}
	return product, nil
}
