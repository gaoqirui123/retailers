package internal

import (
	"common/proto/product"
	"google.golang.org/grpc"
	"product_srv/internal/logic"
)

func RegisterProductServer(server *grpc.Server) {
	product.RegisterProductServer(server, logic.ProductServer{})
}
