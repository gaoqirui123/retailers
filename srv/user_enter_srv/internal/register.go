package internal

import (
	"common/proto/user_enter"
	"google.golang.org/grpc"
	"user_enter_srv/internal/logic"
)

func RegisterUserEnterServer(server *grpc.Server) {
	user_enter.RegisterUserEnterServer(server, &logic.UserEnterServer{})
}
