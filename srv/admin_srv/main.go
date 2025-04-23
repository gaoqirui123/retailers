package main

import (
	"admin_srv/internal"
	"common/consul"
	"common/global"
	"common/grpc"
	_ "common/initialize"
	"fmt"
	grpc2 "google.golang.org/grpc"
)

func main() {
	Conf := global.Config.AdministratorsSrv
	consuls, err := consul.NewConsul()
	if err != nil {
		return
	}
	err = consuls.RegisterConsul("admin_srv", Conf.Host, Conf.Port, []string{"admin"})
	if err != nil {
		return
	}
	fromConsul, err := consuls.GetServiceFromConsul("admin_srv")
	if err != nil {
		return
	}
	fmt.Println(fromConsul)
	grpc.RegisterGrpc(Conf.Host, Conf.Port, func(server *grpc2.Server) {
		internal.RegisterAdministratorsServer(server)
	})
}
