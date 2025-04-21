package client

import (
	"common/proto/article"
	"common/proto/order"
	"common/proto/product"
	"common/proto/user"
	"common/proto/user_enter"
	"context"
	"fmt"
	"google.golang.org/grpc"
)

// GenericClient 通用的 gRPC 客户端处理函数
func GenericClient[T any](ctx context.Context, addr string, newClient func(cc grpc.ClientConnInterface) T, handler func(ctx context.Context, client T) (interface{}, error)) (interface{}, error) {
	// 连接 gRPC 服务
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("连接异常： %s\n", err)
		return nil, err
	}
	defer conn.Close()

	// 创建客户端
	client := newClient(conn)

	// 执行处理函数
	res, err := handler(ctx, client)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func NewUserClient(cc grpc.ClientConnInterface) user.UserClient {
	return user.NewUserClient(cc)
}

func NewProductClient(cc grpc.ClientConnInterface) product.ProductClient {
	return product.NewProductClient(cc)
}

func NewOrderClient(cc grpc.ClientConnInterface) order.OrderClient {
	return order.NewOrderClient(cc)
}

func NewUserEnterClient(cc grpc.ClientConnInterface) user_enter.UserEnterClient {
	return user_enter.NewUserEnterClient(cc)
}
func NewArticleClient(cc grpc.ClientConnInterface) article.ArticleClient {
	return article.NewArticleClient(cc)
}

// UserClients 封装的用户服务客户端处理函数
func UserClients(ctx context.Context, handler func(ctx context.Context, server user.UserClient) (interface{}, error)) (interface{}, error) {
	return GenericClient(ctx, "127.0.0.1:8081", NewUserClient, handler)
}

// ProductClients 封装的产品服务客户端处理函数
func ProductClients(ctx context.Context, handler func(ctx context.Context, server product.ProductClient) (interface{}, error)) (interface{}, error) {
	return GenericClient(ctx, "127.0.0.1:8082", NewProductClient, handler)
}

// OrderClients  封装的订单服务客户端处理函数
func OrderClients(ctx context.Context, handler func(ctx context.Context, server order.OrderClient) (interface{}, error)) (interface{}, error) {
	return GenericClient(ctx, "127.0.0.1:8083", NewOrderClient, handler)
}

// UserEnterClients 封装的购物车服务客户端处理函数
func UserEnterClients(ctx context.Context, handler func(ctx context.Context, server user_enter.UserEnterClient) (interface{}, error)) (interface{}, error) {
	return GenericClient(ctx, "127.0.0.1:8084", NewUserEnterClient, handler)
}

// NewArticleClients 封装的文章服务客户端处理函数
func NewArticleClients(ctx context.Context, handler func(ctx context.Context, server article.ArticleClient) (interface{}, error)) (interface{}, error) {
	return GenericClient(ctx, "127.0.0.1:8085", NewArticleClient, handler)
}
