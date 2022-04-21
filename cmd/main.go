package main

import (
	"context"
	"delay_mq_v2/conf"
	"delay_mq_v2/server/http"
	"delay_mq_v2/service"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	srv *service.Service
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	err := conf.Init()
	if err != nil {
		fmt.Printf("conf init error#%v", err)
		return
	}
	srv = service.New(conf.Conf)
	go srv.Run(ctx)
	go http.Start(conf.Conf, srv)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	for {
		s := <-c
		log.Printf("dmq service get a signal %s", s.String())
		switch s {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			cancel()
			log.Printf("dmq service exit")
			time.Sleep(time.Second)
			return
		case syscall.SIGHUP:
		default:
			return
		}
	}
}
