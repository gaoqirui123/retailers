package internal

import (
	"common/proto/user"
	"google.golang.org/grpc"
	"user_srv/internal/logic"
)

func RegisterUserServer(server *grpc.Server) {
	user.RegisterUserServer(server, logic.UserServer{})
}
