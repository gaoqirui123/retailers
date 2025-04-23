package main

import (
	"common/consul"
	"common/global"
	"common/grpc"
	_ "common/initialize"
	"distribution_srv/internal"
	"fmt"
	grpc2 "google.golang.org/grpc"
)

func main() {
	Conf := global.Config.DistributionSrv
	consuls, err := consul.NewConsul()
	if err != nil {
		return
	}
	err = consuls.RegisterConsul("distribution_srv", Conf.Host, Conf.Port, []string{"distribution"})
	if err != nil {
		return
	}
	fromConsul, err := consuls.GetServiceFromConsul("distribution_srv")
	if err != nil {
		return
	}
	fmt.Println(fromConsul)
	grpc.RegisterGrpc(Conf.Host, Conf.Port, func(server *grpc2.Server) {
		internal.RegisterDistributionServer(server)
	})
}
