package client

import (
	"common/proto/article"
	"context"
)

func ArticleAdd(ctx context.Context, in *article.ArticleAddRequest) (*article.ArticleAddResponse, error) {
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

// 查询文章管理分类列表
func CategoryList(ctx context.Context, in *article.CategoryListRequest) (*article.CategoryListResponse, error) {
	clients, err := NewArticleClients(ctx, func(ctx context.Context, server article.ArticleClient) (interface{}, error) {
		release, err := server.CategoryList(ctx, in)
		if err != nil {
			return nil, err
		}
		return release, nil
	})
	if err != nil {
		return nil, err
	}
	return clients.(*article.CategoryListResponse), err
}

// 查询文章管理分类列表
func ArticleSearch(ctx context.Context, in *article.ArticleSearchRequest) (*article.ArticleSearchResponse, error) {
	clients, err := NewArticleClients(ctx, func(ctx context.Context, server article.ArticleClient) (interface{}, error) {
		release, err := server.ArticleSearch(ctx, in)
		if err != nil {
			return nil, err
		}
		return release, nil
	})
	if err != nil {
		return nil, err
	}
	return clients.(*article.ArticleSearchResponse), err
}

// 编辑文章
func EditArticle(ctx context.Context, in *article.EditArticleRequest) (*article.EditArticleResponse, error) {
	clients, err := NewArticleClients(ctx, func(ctx context.Context, server article.ArticleClient) (interface{}, error) {
		release, err := server.EditArticle(ctx, in)
		if err != nil {
			return nil, err
		}
		return release, nil
	})
	if err != nil {
		return nil, err
	}
	return clients.(*article.EditArticleResponse), err
}

// 删除文章管理
func DeleteArticle(ctx context.Context, in *article.DeleteRequest) (*article.DeleteResponse, error) {
	clients, err := NewArticleClients(ctx, func(ctx context.Context, server article.ArticleClient) (interface{}, error) {
		release, err := server.DeleteArticle(ctx, in)
		if err != nil {
			return nil, err
		}
		return release, nil
	})
	if err != nil {
		return nil, err
	}
	return clients.(*article.DeleteResponse), err
}

// 删除文章分类
func DeleteArticleCategory(ctx context.Context, in *article.DeleteRequest) (*article.DeleteResponse, error) {
	clients, err := NewArticleClients(ctx, func(ctx context.Context, server article.ArticleClient) (interface{}, error) {
		release, err := server.DeleteArticleCategory(ctx, in)
		if err != nil {
			return nil, err
		}
		return release, nil
	})
	if err != nil {
		return nil, err
	}
	return clients.(*article.DeleteResponse), err
}
