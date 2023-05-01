package service

import (
	"context"
	"rpc-demo/protorpc/server/pb"
)

type ArithSerive struct {
	pb.UnimplementedArithSeriveServer
}

func (this *ArithSerive) Multiply(ctx context.Context, req *pb.ArithRequst) (*pb.ArithResponse, error) {

	return &pb.ArithResponse{Pro: req.A * req.B}, nil
}
