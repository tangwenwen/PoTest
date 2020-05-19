package main

import (
	"PoTest/config"
	"PoTest/plugins/logs"
	"flag"
	"github.com/micro/go-web"
)

func main() {

	port := flag.String("http port", "7777", "http listen port")
	flag.Parse()

	service := web.NewService(
		web.Address("0.0.0.0:" + *port),
	)

	service.Init()
	service.Handle("/", config.GetRouterContainer())

	logs.Info("http has listen port on 7777")
	if err := service.Run(); err != nil {
		logs.Error(err)
	}

}
