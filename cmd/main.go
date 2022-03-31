package main

import (
	"delay_mq_v2/conf"
	"delay_mq_v2/http"
	"delay_mq_v2/service"
	"fmt"
)

var (
	srv *service.Service
)

func main()  {
	err := conf.Init()
	if err != nil {
		fmt.Printf("conf init error#%v", err)
		return
	}
	srv = service.New(conf.Conf)
	http.Start(conf.Conf, srv)
}
