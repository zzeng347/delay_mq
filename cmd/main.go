package main

import (
	"delay_mq_v2/conf"
	"fmt"
)

func main()  {
	err := conf.Init()
	fmt.Println(err)
	fmt.Println("1")
}
