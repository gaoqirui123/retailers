package logic

import (
	"common/proto/user_enter"
	"context"
	"user_enter_srv/internal/handler"
)

type UserEnterServer struct {
	user_enter.UnimplementedUserEnterServer
}

func (es UserEnterServer) Register(ctx context.Context, in *user_enter.UserEnterRegisterRequest) (*user_enter.UserEnterRegisterResponse, error) {
	register, err := handler.Register(in)
	if err != nil {
		return nil, err
	}
	return register, nil
}
