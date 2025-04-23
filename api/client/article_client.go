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
func ArticleClients[TRequest, TResponse any](ctx context.Context, request TRequest, operation func(ctx context.Context, client article.ArticleClient, req TRequest) (TResponse, error)) (TResponse, error) {
	return ExecuteGRPCOperation(ctx, "127.0.0.1:8085", NewArticleClient, request, operation)
}

// ArticleAdd 添加文章
func ArticleAdd(ctx context.Context, in *article.ArticleAddRequest) (*article.ArticleAddResponse, error) {
	return ArticleClients(ctx, in, func(ctx context.Context, client article.ArticleClient, req *article.ArticleAddRequest) (*article.ArticleAddResponse, error) {
		return client.ArticleAdd(ctx, req)
	})
}

// CategoryAdd 文章分类添加
func CategoryAdd(ctx context.Context, in *article.CategoryAddRequest) (*article.CategoryAddResponse, error) {
	return ArticleClients(ctx, in, func(ctx context.Context, client article.ArticleClient, req *article.CategoryAddRequest) (*article.CategoryAddResponse, error) {
		return client.CategoryAdd(ctx, req)
	})
}

// ArticleList 查询文章管理列表
func ArticleList(ctx context.Context, in *article.ArticleListRequest) (*article.ArticleListResponse, error) {
	return ArticleClients(ctx, in, func(ctx context.Context, client article.ArticleClient, req *article.ArticleListRequest) (*article.ArticleListResponse, error) {
		return client.ArticleList(ctx, req)
	})
}
