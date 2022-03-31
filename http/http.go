package http

import (
	"delay_mq_v2/conf"
	"delay_mq_v2/library/net/http"
	"delay_mq_v2/service"
	"log"
)

func Start(c *conf.Config, s *service.Service)  {
	//log.Printf("http listen %s\n", c.HTTP.Address)
	//err := http.ListenAndServe(c.HTTP.Address, nil)
	//log.Fatalln(err)
	srv := http.InitHttp(c.HTTP)
	log.Printf("http start, listen %s\n", c.HTTP.Address)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}