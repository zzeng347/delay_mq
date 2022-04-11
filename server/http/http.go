package http

import (
	"delay_mq_v2/conf"
	"delay_mq_v2/library/net/http"
	"delay_mq_v2/server/http/api"
	"delay_mq_v2/service"
	"github.com/gin-gonic/gin"
	"log"
)

func Start(c *conf.Config, s *service.Service)  {
	router := InitRouter()
	srv := http.InitHttp(c.HTTP, router)
	log.Printf("http start, listen %s\n", c.HTTP.Address)
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}

func InitRouter() *gin.Engine {
	r := gin.New()
	// debug or release
	gin.SetMode("debug")

	r.GET("/", api.Hello)
	r.GET("/push", api.Push)
	r.GET("/pop", api.Pop)
	r.GET("/finish", api.Finish)
	r.GET("/delete", api.Delete)

	return r
}