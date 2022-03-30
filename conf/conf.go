package conf

import (
	"bytes"
	"delay_mq_v2/http"
	"delay_mq_v2/library/cache/redis"
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
}


func init() {
	flag.StringVar(&confPath, "conf", "", "default config path")
}

func Init() error {
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



	fmt.Println("config load")

	return nil
}