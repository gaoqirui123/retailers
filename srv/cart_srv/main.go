package main

import (
	"cart_srv/internal"
	"common/consul"
	"common/global"
	"common/grpc"
	_ "common/initialize"
	"fmt"
	grpc2 "google.golang.org/grpc"
)

func main() {
	Conf := global.Config.CartSrv
	consuls, err := consul.NewConsul()
	if err != nil {
		return
	}
	err = consuls.RegisterConsul("cart_srv", Conf.Host, Conf.Port, []string{"cart"})
	if err != nil {
		return
	}
	fromConsul, err := consuls.GetServiceFromConsul("cart_srv")
	if err != nil {
		return
	}
	fmt.Println(fromConsul)
	grpc.RegisterGrpc(Conf.Host, Conf.Port, func(server *grpc2.Server) {
		internal.RegisterCartServer(server)
	})
}
