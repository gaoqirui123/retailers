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

// CategoryList 查询文章管理分类列表
func CategoryList(ctx context.Context, in *article.CategoryListRequest) (*article.CategoryListResponse, error) {
	return ArticleClients(ctx, in, func(ctx context.Context, client article.ArticleClient, req *article.CategoryListRequest) (*article.CategoryListResponse, error) {
		return client.CategoryList(ctx, req)
	})
}

// ArticleSearch 查询文章管理分类列表
func ArticleSearch(ctx context.Context, in *article.ArticleSearchRequest) (*article.ArticleSearchResponse, error) {
	return ArticleClients(ctx, in, func(ctx context.Context, client article.ArticleClient, req *article.ArticleSearchRequest) (*article.ArticleSearchResponse, error) {
		return client.ArticleSearch(ctx, req)
	})
}

// EditArticle 编辑文章
func EditArticle(ctx context.Context, in *article.EditArticleRequest) (*article.EditArticleResponse, error) {
	return ArticleClients(ctx, in, func(ctx context.Context, client article.ArticleClient, req *article.EditArticleRequest) (*article.EditArticleResponse, error) {
		return client.EditArticle(ctx, req)
	})
}

// DeleteArticle 删除文章管理
func DeleteArticle(ctx context.Context, in *article.DeleteRequest) (*article.DeleteResponse, error) {
	return ArticleClients(ctx, in, func(ctx context.Context, client article.ArticleClient, req *article.DeleteRequest) (*article.DeleteResponse, error) {
		return client.DeleteArticle(ctx, req)
	})
}

// DeleteArticleCategory 删除文章分类
func DeleteArticleCategory(ctx context.Context, in *article.DeleteRequest) (*article.DeleteResponse, error) {
	return ArticleClients(ctx, in, func(ctx context.Context, client article.ArticleClient, req *article.DeleteRequest) (*article.DeleteResponse, error) {
		return client.DeleteArticleCategory(ctx, req)
	})
}

// PostAComment 发布评论
func PostAComment(ctx context.Context, in *article.PostACommentRequest) (*article.PostACommentResponse, error) {
	return ArticleClients(ctx, in, func(ctx context.Context, client article.ArticleClient, req *article.PostACommentRequest) (*article.PostACommentResponse, error) {
		return client.PostAComment(ctx, req)
	})
}

// ArticleThumbsUp 文章点赞
func ArticleThumbsUp(ctx context.Context, in *article.ArticleThumbsUpRequest) (*article.ArticleThumbsUpResponse, error) {
	return ArticleClients(ctx, in, func(ctx context.Context, client article.ArticleClient, req *article.ArticleThumbsUpRequest) (*article.ArticleThumbsUpResponse, error) {
		return client.ArticleThumbsUp(ctx, req)
	})
}

// ArticleThumbsUp //删除评论
func DeleteComment(ctx context.Context, in *article.DeleteCommentRequest) (*article.DeleteCommentResponse, error) {
	return ArticleClients(ctx, in, func(ctx context.Context, client article.ArticleClient, req *article.DeleteCommentRequest) (*article.DeleteCommentResponse, error) {
		return client.DeleteComment(ctx, req)
	})
}

// ArticleThumbsUp //删除评论
func TopLikeArticleRanking(ctx context.Context, in *article.TopLikeArticleRankingRequest) (*article.TopLikeArticleRankingResponse, error) {
	return ArticleClients(ctx, in, func(ctx context.Context, client article.ArticleClient, req *article.TopLikeArticleRankingRequest) (*article.TopLikeArticleRankingResponse, error) {
		return client.TopLikeArticleRanking(ctx, req)
	})
}
