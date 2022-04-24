package http

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type Client struct {
	client		*http.Client
	mutex		sync.RWMutex
}

type ClientConfig struct {
	MaxIdleConns		int
	IdleConnTimeout		time.Duration
	Timeout				time.Duration
}

func NewHttpClient(c *ClientConfig) (client *Client) {
	client = &Client{}
	transport := &http.Transport{
		MaxIdleConns: c.MaxIdleConns,
		IdleConnTimeout: c.IdleConnTimeout,
	}
	client.client = &http.Client{
		Transport: transport,
		Timeout: time.Second * c.Timeout,
	}
	return
}

func (c *Client) Get(url string) (body []byte, err error) {
	resp, err := c.client.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	return
}

func (c *Client) Post(url string, data interface{}, contentType string) (body []byte, err error) {
	jsonStr, err := json.Marshal(data)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return
	}
	req.Header.Add("Content-Type", contentType)
	defer req.Body.Close()

	resp, err := c.client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	return
}