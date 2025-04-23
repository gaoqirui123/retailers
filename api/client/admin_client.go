package client

import (
	administrators "common/proto/admin"
	"context"
)

func AdminLogin(ctx context.Context, in *administrators.AdminLoginReq) (*administrators.AdminLoginResp, error) {
	client, err := NewAdministratorsClients(ctx, func(ctx context.Context, client administrators.AdministratorsClient) (interface{}, error) {
		login, err := client.AdminLogin(ctx, in)
		if err != nil {
			return nil, err
		}
		return login, nil
	})
	if err != nil {
		return nil, err
	}
	return client.(*administrators.AdminLoginResp), nil
}
func ProcessEnter(ctx context.Context, in *administrators.ProcessEnterReq) (*administrators.ProcessEnterResp, error) {
	client, err := NewAdministratorsClients(ctx, func(ctx context.Context, client administrators.AdministratorsClient) (interface{}, error) {
		login, err := client.ProcessEnter(ctx, in)
		if err != nil {
			return nil, err
		}
		return login, nil
	})
	if err != nil {
		return nil, err
	}
	return client.(*administrators.ProcessEnterResp), nil
}
