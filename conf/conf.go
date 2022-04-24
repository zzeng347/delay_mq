package conf

import (
	"bytes"
	"delay_mq_v2/library/cache/redis"
	"delay_mq_v2/library/net/http"
	"flag"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
)

var (
	confPath string
	Conf *Config
)

type Config struct {
	REDIS		*redis.Config
	HTTP		*http.Config
	HTTPCLIENT	*http.ClientConfig
}


func init() {
	flag.StringVar(&confPath, "conf", "", "default config path")
}

func Init() error {
	Conf = new(Config)
	redisConf := new(redis.Config)
	httpConf := new(http.Config)
	httpClientConf := new(http.ClientConfig)

	viper.SetConfigType("toml")
	viper.AutomaticEnv()
	if confPath != "" {
		content, err := ioutil.ReadFile(confPath)
		if err != nil {
			return err
		}
		if err = viper.ReadConfig(bytes.NewBuffer(content)); err != nil {
			return err
		}
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigName("conf")

		if err := viper.ReadInConfig(); err != nil {
			return err
		}
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}

	// redis
	redisConf.Host = viper.GetString("redis.host")
	redisConf.Port = viper.GetString("redis.port")
	redisConf.Password = viper.GetString("redis.password")
	redisConf.DB = viper.GetInt("redis.db")

	// http
	httpConf.Address = viper.GetString("http.address")

	// http client
	httpClientConf.MaxIdleConns = viper.GetInt("http_client.max_idle_conns")
	httpClientConf.IdleConnTimeout = viper.GetDuration("http_client.idle_conn_timeout")
	httpClientConf.Timeout = viper.GetDuration("http_client.timeout")

	// Conf assign
	Conf.REDIS = redisConf
	Conf.HTTP = httpConf
	Conf.HTTPCLIENT = httpClientConf

	return nil
}