package logic

import (
	"cart_srv/internal/handler"
	"common/proto/cart"
	"context"
)

type CartServer struct {
	cart.UnimplementedCartServer
}

func (c CartServer) AddCart(ctx context.Context, in *cart.AddCartRequest) (*cart.AddCartResponse, error) {
	addCart, err := handler.AddCart(in)
	if err != nil {
		return nil, err
	}
	return addCart, nil
}
func (c CartServer) ClearCart(ctx context.Context, in *cart.ClearCartRequest) (*cart.ClearCartResponse, error) {
	clearCart, err := handler.ClearCart(in)
	if err != nil {
		return nil, err
	}
	return clearCart, nil
}
func (c CartServer) DeleteCart(ctx context.Context, in *cart.DeleteCartRequest) (*cart.DeleteCartResponse, error) {
	deleteCart, err := handler.DeleteCart(in)
	if err != nil {
		return nil, err
	}
	return deleteCart, nil
}
func (c CartServer) GetCartList(ctx context.Context, in *cart.GetCartListRequest) (*cart.GetCartListResponse, error) {
	list, err := handler.GetCartList(in)
	if err != nil {
		return nil, err
	}
	return list, nil
}
func (c CartServer) UpdateCart(ctx context.Context, in *cart.UpdateCartRequest) (*cart.UpdateCartResponse, error) {
	list, err := handler.UpdateCart(in)
	if err != nil {
		return nil, err
	}
	return list, nil
}
