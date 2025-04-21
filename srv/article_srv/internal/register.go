package internal

import (
	"article_srv/internal/logic"
	"common/proto/article"
	"google.golang.org/grpc"
)

func RegisterArticleServer(server *grpc.Server) {
	article.RegisterArticleServer(server, logic.Article{})
}
