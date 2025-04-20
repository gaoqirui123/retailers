package main

import (
	"common/consul"
	"common/global"
	"common/grpc"
	_ "common/initialize"
	"fmt"
	grpc2 "google.golang.org/grpc"
	"order_srv/internal"
)

func main() {
	Conf := global.Config.OrderSrv
	consuls, err := consul.NewConsul()
	if err != nil {
		return
	}
	err = consuls.RegisterConsul("order_srv", Conf.Host, Conf.Port, []string{"order"})
	if err != nil {
		return
	}
	fromConsul, err := consuls.GetServiceFromConsul("order_srv")
	if err != nil {
		return
	}
	fmt.Println(fromConsul)
	grpc.RegisterGrpc(Conf.Host, Conf.Port, func(server *grpc2.Server) {
		internal.RegisterOrderServer(server)
	})
}
