package http

import (
	"net/http"
)

type Config struct {
	Address string
}

func InitHttp(c *Config) *http.Server {
	s := &http.Server{
		Addr: c.Address,
	}
	return s
}
