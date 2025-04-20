package internal

import (
	"common/proto/order"
	"google.golang.org/grpc"
	"order_srv/internal/logic"
)

func RegisterOrderServer(server *grpc.Server) {
	order.RegisterOrderServer(server, logic.OrderServer{})
}
