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

// 分销等级设置
func (d DistributionServer) DistributionLevelSetting(ctx context.Context, in *distribution.DistributionLevelSettingRequest) (*distribution.DistributionLevelSettingResponse, error) {
	code, err := handler.DistributionLevelSetting(in)
	if err != nil {
		return nil, err
	}
	return code, err
}

// 佣金排行榜
func (d DistributionServer) TheCharts(ctx context.Context, in *distribution.TheChartsRequest) (*distribution.TheChartsResponse, error) {
	code, err := handler.TheCharts(in)
	if err != nil {
		return nil, err
	}
	return code, err
}

// TODO 用户下级展示
func (d DistributionServer) LookDoneUp(ctx context.Context, in *distribution.LookDoneOrUpReq) (*distribution.LookDoneOrUpResp, error) {
	code, err := handler.LookDoneUp(in)
	if err != nil {
		return nil, err
	}
	return code, err
}

// TODO 用户上级展示
func (d DistributionServer) LookUp(ctx context.Context, in *distribution.LookDoneOrUpReq) (*distribution.LookDoneOrUpResp, error) {
	code, err := handler.LookUp(in)
	if err != nil {
		return nil, err
	}
	return code, err
}
