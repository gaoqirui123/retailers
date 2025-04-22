package logic

import (
	"article_srv/internal/handler"
	"common/proto/article"
	"context"
)

type ArticleServer struct {
	article.UnimplementedArticleServer
}

// 文章发布
func (a ArticleServer) ArticleAdd(ctx context.Context, in *article.ArticleAddRequest) (*article.ArticleAddResponse, error) {
	add, err := handler.ArticleAdd(in)
	if err != nil {
		return nil, err
	}
	return add, nil
}

// 文章分类添加
func (a ArticleServer) CategoryAdd(ctx context.Context, in *article.CategoryAddRequest) (*article.CategoryAddResponse, error) {
	add, err := handler.CategoryAdd(in)
	if err != nil {
		return nil, err
	}
	return add, nil
}

// 查询文章管理列表
func (a ArticleServer) ArticleList(ctx context.Context, in *article.ArticleListRequest) (*article.ArticleListResponse, error) {
	list, err := handler.ArticleList(in)
	if err != nil {
		return nil, err
	}
	return list, nil
}
