package internal

import (
	"admin_srv/internal/logic"
	administrators "common/proto/admin"
	"google.golang.org/grpc"
)

func RegisterAdministratorsServer(server *grpc.Server) {
	administrators.RegisterAdministratorsServer(server, logic.AdministratorsServer{})
}
