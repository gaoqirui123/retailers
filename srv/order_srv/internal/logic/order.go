package logic

import (
	"common/proto/order"
	"context"
	"errors"
	"order_srv/internal/handler"
)

type OrderServer struct {
	order.UnimplementedOrderServer
}

func (o OrderServer) AddOrder(ctx context.Context, in *order.AddOrderRequest) (*order.AddOrderResponse, error) {
	add, err := handler.AddOrder(in)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return add, err
}

func (o OrderServer) PayCallback(ctx context.Context, in *order.PayCallbackRequest) (*order.PayCallbackResponse, error) {
	callback, err := handler.PayCallback(in)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return callback, err
}

func (o OrderServer) OrderList(ctx context.Context, in *order.OrderListRequest) (*order.OrderListResponse, error) {
	list, err := handler.OrderList(in)
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return list, err
}
