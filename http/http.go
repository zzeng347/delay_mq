package http

import (
	"delay_mq_v2/conf"
	"delay_mq_v2/service"
	"log"
	"net/http"
)

type Config struct {
	Address string
}

func Init(c *conf.Config, s *service.Service)  {
	log.Printf("http listen %s\n", c.HTTP.Address)
	err := http.ListenAndServe(c.HTTP.Address, nil)
	log.Fatalln(err)
}