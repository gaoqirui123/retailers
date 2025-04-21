package logic

import (
	"article_srv/internal/handler"
	"common/proto/article"
	"context"
)

type Article struct {
	article.UnimplementedArticleServer
}

// 文章发布
func (a Article) ArticleRelease(ctx context.Context, in *article.ArticleAddReq) (*article.ArticleAddResp, error) {
	add, err := handler.ArticleAdd(in)
	if err != nil {
		return nil, err
	}
	return add, nil
}

// 文章分类添加
func (a Article) CategoryAdd(ctx context.Context, in *article.CategoryAddReq) (*article.CategoryAddResp, error) {
	add, err := handler.CategoryAdd(in)
	if err != nil {
		return nil, err
	}
	return add, nil
}

// 查询文章管理列表
func (a Article) ArticleList(ctx context.Context, in *article.ArticleListReq) (*article.ArticleListResp, error) {
	add, err := handler.ArticleList(in)
	if err != nil {
		return nil, err
	}
	return add, nil
}
