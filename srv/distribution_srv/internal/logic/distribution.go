package logic

import (
	"common/proto/distribution"
	"context"
	"distribution_srv/internal/handler"
)

type DistributionServer struct {
	distribution.UnimplementedDistributionServer
}

// 生成邀请码
func (d DistributionServer) GenerateInvitationCode(ctx context.Context, in *distribution.GenerateInvitationCodeRequest) (*distribution.GenerateInvitationCodeResponse, error) {
	code, err := handler.GenerateInvitationCode(in)
	if err != nil {
		return nil, err
	}
	return code, err
}

// 用户填写邀请码
func (d DistributionServer) UserFillsInInvitationCode(ctx context.Context, in *distribution.UserFillsInInvitationCodeRequest) (*distribution.UserFillsInInvitationCodeResponse, error) {
	code, err := handler.UserFillsInInvitationCode(in)
	if err != nil {
		return nil, err
	}
	return code, err
}
