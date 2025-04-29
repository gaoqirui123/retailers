package client

import (
	"common/proto/distribution"
	"context"
	"google.golang.org/grpc"
)

// NewDistributionClient 创建一个新的 Distribution 服务客户端
func NewDistributionClient(cc grpc.ClientConnInterface) distribution.DistributionClient {
	return distribution.NewDistributionClient(cc)
}

// DistributionClients 封装 Distribution 服务客户端操作逻辑
func DistributionClients[TRequest, TResponse any](ctx context.Context, request TRequest, operation func(ctx context.Context, client distribution.DistributionClient, req TRequest) (TResponse, error)) (TResponse, error) {
	return ExecuteGRPCOperation(ctx, "127.0.0.1:8088", NewDistributionClient, request, operation)
}

// GenerateInvitationCode 生成邀请码
func GenerateInvitationCode(ctx context.Context, in *distribution.GenerateInvitationCodeRequest) (*distribution.GenerateInvitationCodeResponse, error) {
	return DistributionClients(ctx, in, func(ctx context.Context, client distribution.DistributionClient, req *distribution.GenerateInvitationCodeRequest) (*distribution.GenerateInvitationCodeResponse, error) {
		return client.GenerateInvitationCode(ctx, req)
	})
}

// GenerateInvitationCode 填写邀请码
func UserFillsInInvitationCode(ctx context.Context, in *distribution.UserFillsInInvitationCodeRequest) (*distribution.UserFillsInInvitationCodeResponse, error) {
	return DistributionClients(ctx, in, func(ctx context.Context, client distribution.DistributionClient, req *distribution.UserFillsInInvitationCodeRequest) (*distribution.UserFillsInInvitationCodeResponse, error) {
		return client.UserFillsInInvitationCode(ctx, req)
	})
}

// DistributionLevelSetting 分销等级设置
func DistributionLevelSetting(ctx context.Context, in *distribution.DistributionLevelSettingRequest) (*distribution.DistributionLevelSettingResponse, error) {
	return DistributionClients(ctx, in, func(ctx context.Context, client distribution.DistributionClient, req *distribution.DistributionLevelSettingRequest) (*distribution.DistributionLevelSettingResponse, error) {
		return client.DistributionLevelSetting(ctx, req)
	})
}

// TheCharts 佣金排行榜
func TheCharts(ctx context.Context, in *distribution.TheChartsRequest) (*distribution.TheChartsResponse, error) {
	return DistributionClients(ctx, in, func(ctx context.Context, client distribution.DistributionClient, req *distribution.TheChartsRequest) (*distribution.TheChartsResponse, error) {
		return client.TheCharts(ctx, req)
	})
}
