package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Config struct {
	Address string
}

func InitHttp(c *Config, r *gin.Engine) *http.Server {
	s := &http.Server{
		Addr: c.Address,
		Handler: r,
	}
	return s
}
