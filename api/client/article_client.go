package client

import (
	"common/proto/article"
	"context"
	"google.golang.org/grpc"
)

func NewArticleClient(cc grpc.ClientConnInterface) article.ArticleClient {
	return article.NewArticleClient(cc)
}

// ArticleClients 封装的文章服务客户端处理函数
func ArticleClients(ctx context.Context, handler func(ctx context.Context, server article.ArticleClient) (interface{}, error)) (interface{}, error) {
	return GenericClient(ctx, "127.0.0.1:8085", NewArticleClient, handler)
}

func ArticleAdd(ctx context.Context, in *article.ArticleAddRequest) (*article.ArticleAddResponse, error) {
	clients, err := ArticleClients(ctx, func(ctx context.Context, server article.ArticleClient) (interface{}, error) {
		release, err := server.ArticleAdd(ctx, in)
		if err != nil {
			return nil, err
		}
		return release, nil
	})
	if err != nil {
		return nil, err
	}
	return clients.(*article.ArticleAddResponse), err
}

// 文章分类添加
func CategoryAdd(ctx context.Context, in *article.CategoryAddRequest) (*article.CategoryAddResponse, error) {
	clients, err := ArticleClients(ctx, func(ctx context.Context, server article.ArticleClient) (interface{}, error) {
		release, err := server.CategoryAdd(ctx, in)
		if err != nil {
			return nil, err
		}
		return release, nil
	})
	if err != nil {
		return nil, err
	}
	return clients.(*article.CategoryAddResponse), err
}

// 查询文章管理列表
func ArticleList(ctx context.Context, in *article.ArticleListRequest) (*article.ArticleListResponse, error) {
	clients, err := ArticleClients(ctx, func(ctx context.Context, server article.ArticleClient) (interface{}, error) {
		release, err := server.ArticleList(ctx, in)
		if err != nil {
			return nil, err
		}
		return release, nil
	})
	if err != nil {
		return nil, err
	}
	return clients.(*article.ArticleListResponse), err
}
