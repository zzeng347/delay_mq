package main

import (
	"context"
	"delay_mq_v2/server/grpc/proto/delay"
	"fmt"

	"google.golang.org/grpc"
)

func main() {
	// 连接服务器
	conn, err := grpc.Dial("192.168.10.31:8972", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("faild to connect: %v", err)
	}
	defer conn.Close()

	c := delay.NewDelayClient(conn)
	// 调用服务端

	// test
	r, err := c.RpcServerTest(context.Background(), &delay.Req{Id: "job_id:1111"})
	if err != nil {
		fmt.Printf("test failed: %v", err)
	} else {
		fmt.Printf("test success: %s!\n", r.Message)
		fmt.Printf("job id: %s, body: %s \n", r.Code, r.Message)
	}

}