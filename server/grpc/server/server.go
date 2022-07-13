package server

import (
	"context"
	"delay_mq_v2/conf"
	"delay_mq_v2/server/grpc/proto/delay"
	"delay_mq_v2/service"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime" // 注意v2版本
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
	"strings"
)

// RpcServer 业务实现方法的容器
type RpcServer struct{
	delay.UnimplementedDelayServer
}

func RunRpcServer(c *conf.Config, s *service.Service)  {
	lis, err := net.Listen(c.RPC.Network, c.RPC.Address)
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		log.Fatalln(err)
	}
	log.Printf("rpc listen: %s\n", c.RPC.Address)
	srv := grpc.NewServer()                      // 创建gRPC服务器
	delay.RegisterDelayServer(srv, &RpcServer{}) // 在gRPC服务端注册服务

	go func() {
		// Serve方法在lis上接受传入连接，为每个连接创建一个ServerTransport和RpcServer的goroutine。
		// 该goroutine读取gRPC请求，然后调用已注册的处理程序来响应它们。
		log.Fatalln(srv.Serve(lis))
	}()

	// 创建一个连接到我们刚刚启动的 gRPC 服务器的客户端连接
	// gRPC-Gateway 就是通过它来代理请求（将HTTP请求转为RPC请求）
	conn, err := grpc.DialContext(
		context.Background(),
		"0.0.0.0:8972",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	// 注册Greeter
	err = delay.RegisterDelayHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}
	// 8090端口提供gRPC-Gateway服务
	// curl -X POST -k http://localhost:8090/v1/example/echo -d '{"id": "job_id:122"}'
	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())



	//reflection.Register(srv) //在给定的gRPC服务器上注册服务器反射服务
	//// Serve方法在lis上接受传入连接，为每个连接创建一个ServerTransport和RpcServer的goroutine。
	//// 该goroutine读取gRPC请求，然后调用已注册的处理程序来响应它们。
	//err = srv.Serve(lis)
	//if err != nil {
	//	fmt.Printf("failed to serve: %v", err)
	//	log.Fatalln(err)
	//}
}

func (s *RpcServer) RpcServerTest(ctx context.Context, in *delay.Req) (*delay.Replay, error) {
	jobId := strings.TrimSpace(in.Id)
	body := &delay.Replay{}
	body.Code = 1
	body.Message = jobId
	return body, nil
}