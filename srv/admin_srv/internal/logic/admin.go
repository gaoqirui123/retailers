package logic

import (
	"admin_srv/internal/handler"
	administrators "common/proto/admin"
	"context"
)

type AdministratorsServer struct {
	administrators.UnimplementedAdministratorsServer
}

// AdminLogin TODO:管理员登录
func (as AdministratorsServer) AdminLogin(ctx context.Context, in *administrators.AdminLoginReq) (*administrators.AdminLoginResp, error) {
	login, err := handler.AdminLogin(in)
	if err != nil {
		return nil, err
	}
	return login, nil
}

// ProcessEnter  TODO:管理员审核商家申请
func (as AdministratorsServer) ProcessEnter(ctx context.Context, in *administrators.ProcessEnterReq) (*administrators.ProcessEnterResp, error) {
	login, err := handler.ProcessEnter(in)
	if err != nil {
		return nil, err
	}
	return login, nil
}
