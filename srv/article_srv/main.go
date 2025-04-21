package main

import (
	"article_srv/internal"
	"common/consul"
	"common/global"
	"common/grpc"
	_ "common/initialize"
	"fmt"
	grpc2 "google.golang.org/grpc"
)

func main() {
	Conf := global.Config.ArticleSrv
	consuls, err := consul.NewConsul()
	if err != nil {
		return
	}
	err = consuls.RegisterConsul("article_srv", Conf.Host, Conf.Port, []string{"article"})
	if err != nil {
		return
	}
	fromConsul, err := consuls.GetServiceFromConsul("article_srv")
	if err != nil {
		return
	}
	fmt.Println(fromConsul)
	grpc.RegisterGrpc(Conf.Host, Conf.Port, func(server *grpc2.Server) {
		internal.RegisterArticleServer(server)
	})
}
