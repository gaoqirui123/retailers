package internal

import (
	"cart_srv/internal/logic"
	"common/proto/cart"
	"google.golang.org/grpc"
)

func RegisterCartServer(server *grpc.Server) {
	cart.RegisterCartServer(server, logic.CartServer{})
}
