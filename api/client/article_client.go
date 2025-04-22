package client

import (
	"common/proto/article"
	"context"
)

func ArticleRelease(ctx context.Context, in *article.ArticleAddRequest) (*article.ArticleAddResponse, error) {
	clients, err := NewArticleClients(ctx, func(ctx context.Context, server article.ArticleClient) (interface{}, error) {
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
	clients, err := NewArticleClients(ctx, func(ctx context.Context, server article.ArticleClient) (interface{}, error) {
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
	clients, err := NewArticleClients(ctx, func(ctx context.Context, server article.ArticleClient) (interface{}, error) {
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
