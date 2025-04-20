package main

import (
	"common/consul"
	"common/global"
	"common/grpc"
	_ "common/initialize"
	"fmt"
	grpc2 "google.golang.org/grpc"
	"product_srv/internal"
)

func main() {
	Conf := global.Config.ProductSrv
	consuls, err := consul.NewConsul()
	if err != nil {
		return
	}
	err = consuls.RegisterConsul("product_srv", Conf.Host, Conf.Port, []string{"product"})
	if err != nil {
		return
	}
	fromConsul, err := consuls.GetServiceFromConsul("product_srv")
	if err != nil {
		return
	}
	fmt.Println(fromConsul)
	grpc.RegisterGrpc(Conf.Host, Conf.Port, func(server *grpc2.Server) {
		internal.RegisterProductServer(server)
	})
}
