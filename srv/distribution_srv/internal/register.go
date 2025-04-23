package internal

import (
	"common/proto/distribution"
	"distribution_srv/internal/logic"
	"google.golang.org/grpc"
)

func RegisterDistributionServer(server *grpc.Server) {
	distribution.RegisterDistributionServer(server, logic.DistributionServer{})
}
