package client

import (
	"common/proto/article"
	"context"
)

func ArticleRelease(ctx context.Context, in *article.ArticleAddReq) (*article.ArticleAddResp, error) {
	clients, err := NewArticleClients(ctx, func(ctx context.Context, server article.ArticleClient) (interface{}, error) {
		release, err := server.ArticleRelease(ctx, in)
		if err != nil {
			return nil, err
		}
		return release, nil
	})
	if err != nil {
		return nil, err
	}
	return clients.(*article.ArticleAddResp), err
}

// 文章分类添加
func CategoryAdd(ctx context.Context, in *article.CategoryAddReq) (*article.CategoryAddResp, error) {
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
	return clients.(*article.CategoryAddResp), err
}

// 查询文章管理列表
func ArticleList(ctx context.Context, in *article.ArticleListReq) (*article.ArticleListResp, error) {
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
	return clients.(*article.ArticleListResp), err
}
