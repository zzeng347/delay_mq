package server

import (
	"context"
	"delay_mq_v2/conf"
	pb "delay_mq_v2/server/grpc/proto"
	"delay_mq_v2/service"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"strings"
)

// RpcServer 业务实现方法的容器
type RpcServer struct{}

func RunRpcServer(c *conf.Config, s *service.Service)  {
	lis, err := net.Listen(c.RPC.Network, c.RPC.Address)
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		log.Fatalln(err)
	}
	log.Printf("rpc listen: %s\n", c.RPC.Address)
	srv := grpc.NewServer() // 创建gRPC服务器
	pb.RegisterDelayServer(srv, &RpcServer{}) // 在gRPC服务端注册服务

	reflection.Register(srv) //在给定的gRPC服务器上注册服务器反射服务
	// Serve方法在lis上接受传入连接，为每个连接创建一个ServerTransport和RpcServer的goroutine。
	// 该goroutine读取gRPC请求，然后调用已注册的处理程序来响应它们。
	err = srv.Serve(lis)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		log.Fatalln(err)
	}
}

func (s *RpcServer) RpcServerTest(ctx context.Context, in *pb.Req) (*pb.Replay, error) {
	jobId := strings.TrimSpace(in.Id)
	body := &pb.Replay{}
	body.Code = 1
	body.Message = jobId
	return body, nil
}