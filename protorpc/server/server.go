package main

import (
	"net"
	"rpc-demo/protorpc/server/pb"
	"rpc-demo/protorpc/server/service"

	"google.golang.org/grpc"
)

type Arith struct{}

func (a *Arith) Multiply(req *pb.ArithRequst, res *pb.ArithResponse) error {
	res.Pro = req.GetA() * req.GetB()
	return nil
}

func main() {
	rpcServer := grpc.NewServer()
	pb.RegisterArithSeriveServer(rpcServer, new(service.ArithSerive))
	listen, _ := net.Listen("tcp", ":8087")
	rpcServer.Serve(listen)
}
