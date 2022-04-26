package http

import (
	"delay_mq_v2/conf"
	"delay_mq_v2/library/net/http"
	"delay_mq_v2/service"
	"github.com/gin-gonic/gin"
	"log"
	nhttp "net/http"
)

const (
	SuccessCode = 1
	ErrorCode = 0
)

var (
	hSrv  *service.Service
	statusOk = nhttp.StatusOK
)

// JSON common json struct.
type JSON struct {
	Code    int         `json:"code"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

type MapJSON map[string]interface{}

func Start(c *conf.Config, s *service.Service)  {
	hSrv = s
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

	r.GET("/", Hello)
	r.POST("/push", Push)
	r.POST("/pop", Pop)
	r.POST("/finish", Finish)
	r.POST("/delete", Delete)
	r.POST("/biz", Biz)
	r.POST("/json_raw", JsonRaw)
	r.POST("/map", Map)

	return r
}

// Success return json
func Success(c *gin.Context, data interface{})  {
	c.JSON(statusOk, JSON{
		Code: SuccessCode,
		Data: data,
	})
}

// Fail return json
func Fail(c *gin.Context, message string)  {
	c.JSON(statusOk, JSON{
		Code: ErrorCode,
		Message: message,
	})
}