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

// 查询文章管理分类列表
func (a ArticleServer) CategoryList(ctx context.Context, in *article.CategoryListRequest) (*article.CategoryListResponse, error) {
	list, err := handler.CategoryList(in)
	if err != nil {
		return nil, err
	}
	return list, nil
}

// 文章标题搜索
func (a ArticleServer) ArticleSearch(ctx context.Context, in *article.ArticleSearchRequest) (*article.ArticleSearchResponse, error) {
	list, err := handler.ArticleSearch(in)
	if err != nil {
		return nil, err
	}
	return list, nil
}

// 编辑文章
func (a ArticleServer) EditArticle(ctx context.Context, in *article.EditArticleRequest) (*article.EditArticleResponse, error) {
	list, err := handler.EditArticle(in)
	if err != nil {
		return nil, err
	}
	return list, nil
}

// 删除文章管理
func (a ArticleServer) DeleteArticle(ctx context.Context, in *article.DeleteRequest) (*article.DeleteResponse, error) {
	list, err := handler.DeleteArticle(in)
	if err != nil {
		return nil, err
	}
	return list, nil
}

// 删除文章分类
func (a ArticleServer) DeleteArticleCategory(ctx context.Context, in *article.DeleteRequest) (*article.DeleteResponse, error) {
	list, err := handler.DeleteArticleCategory(in)
	if err != nil {
		return nil, err
	}
	return list, nil
}

// 发布评论
func (a ArticleServer) PostAComment(ctx context.Context, in *article.PostACommentRequest) (*article.PostACommentResponse, error) {
	list, err := handler.PostAComment(in)
	if err != nil {
		return nil, err
	}
	return list, nil
}

// 文章点赞
func (a ArticleServer) ArticleThumbsUp(ctx context.Context, in *article.ArticleThumbsUpRequest) (*article.ArticleThumbsUpResponse, error) {
	list, err := handler.ArticleThumbsUp(in)
	if err != nil {
		return nil, err
	}
	return list, nil
}

// 删除评论
func (a ArticleServer) DeleteComment(ctx context.Context, in *article.DeleteCommentRequest) (*article.DeleteCommentResponse, error) {
	list, err := handler.DeleteComment(in)
	if err != nil {
		return nil, err
	}
	return list, nil
}

// 删除评论
func (a ArticleServer) TopLikeArticleRanking(ctx context.Context, in *article.TopLikeArticleRankingRequest) (*article.TopLikeArticleRankingResponse, error) {
	list, err := handler.TopLikeArticleRanking(in)
	if err != nil {
		return nil, err
	}
	return list, nil
}
