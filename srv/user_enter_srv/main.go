package main

import (
	"common/consul"
	"common/global"
	"common/grpc"
	_ "common/initialize"
	"fmt"
	grpc2 "google.golang.org/grpc"
	"user_enter_srv/internal"
)

func main() {
	Conf := global.Config.UserEnterSrv
	consuls, err := consul.NewConsul()
	if err != nil {
		return
	}
	err = consuls.RegisterConsul("user_enter_srv", Conf.Host, Conf.Port, []string{"user_enter"})
	if err != nil {
		return
	}
	fromConsul, err := consuls.GetServiceFromConsul("user_enter_srv")
	if err != nil {
		return
	}
	fmt.Println(fromConsul)
	grpc.RegisterGrpc(Conf.Host, Conf.Port, func(server *grpc2.Server) {
		internal.RegisterUserEnterServer(server)
	})
}
