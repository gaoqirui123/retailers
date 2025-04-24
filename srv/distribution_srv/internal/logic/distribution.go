package logic

import (
	"common/proto/distribution"
	"context"
	"distribution_srv/internal/handler"
)

type DistributionServer struct {
	distribution.UnimplementedDistributionServer
}

func (d DistributionServer) GenerateInvitationCode(ctx context.Context, in *distribution.GenerateInvitationCodeRequest) (*distribution.GenerateInvitationCodeResponse, error) {
	code, err := handler.GenerateInvitationCode(in)
	if err != nil {
		return nil, err
	}
	return code, err
}
