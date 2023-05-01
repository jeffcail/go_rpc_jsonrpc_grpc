package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type Param struct {
	Width, Height int
}

type Rect struct{}

func (r *Rect) Area(p Param, ret *int) error {
	*ret = p.Width * p.Height
	return nil
}

func main() {
	rpc.Register(new(Rect))
	lis, err := net.Listen("tcp", ":8096")
	if err != nil {
		log.Panicln(err)
	}

	for {
		conn, err := lis.Accept()
		if err != nil {
			continue
		}
		go func(conn net.Conn) {
			fmt.Println("new client")
			jsonrpc.ServeConn(conn)
		}(conn)
	}
}
