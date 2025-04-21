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
