package client

import (
	"common/proto/order"
	"context"
)

func AddOrder(ctx context.Context, in *order.AddOrderRequest) (*order.AddOrderResponse, error) {
	client, err := OrderClients(ctx, func(ctx context.Context, client order.OrderClient) (interface{}, error) {
		login, err := client.AddOrder(ctx, in)
		if err != nil {
			return nil, err
		}
		return login, nil
	})
	if err != nil {
		return nil, err
	}
	return client.(*order.AddOrderResponse), nil
}
func PayCallback(ctx context.Context, in *order.PayCallbackRequest) (*order.PayCallbackResponse, error) {
	client, err := OrderClients(ctx, func(ctx context.Context, client order.OrderClient) (interface{}, error) {
		callback, err := client.PayCallback(ctx, in)
		if err != nil {
			return nil, err
		}
		return callback, nil
	})
	if err != nil {
		return nil, err
	}
	return client.(*order.PayCallbackResponse), nil
}
func OrderList(ctx context.Context, in *order.OrderListRequest) (*order.OrderListResponse, error) {
	client, err := OrderClients(ctx, func(ctx context.Context, client order.OrderClient) (interface{}, error) {
		list, err := client.OrderList(ctx, in)
		if err != nil {
			return nil, err
		}
		return list, nil
	})
	if err != nil {
		return nil, err
	}
	return client.(*order.OrderListResponse), nil
}
