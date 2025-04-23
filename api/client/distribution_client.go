package client

import (
	"common/proto/distribution"
	"context"
)

func GenerateInvitationCode(ctx context.Context, in *distribution.GenerateInvitationCodeRequest) (*distribution.GenerateInvitationCodeResponse, error) {
	clients, err := DistributionClients(ctx, func(ctx context.Context, server distribution.DistributionClient) (interface{}, error) {
		code, err := server.GenerateInvitationCode(ctx, in)
		if err != nil {
			return nil, err
		}
		return code, err
	})
	if err != nil {
		return nil, err
	}
	return clients.(*distribution.GenerateInvitationCodeResponse), err
}
