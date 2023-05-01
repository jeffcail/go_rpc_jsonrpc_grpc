package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

type Param struct {
	Width, Height int
}

func main() {

	conn, err := jsonrpc.Dial("tcp", ":8096")
	if err != nil {
		log.Panicln(err)
	}

	ret := 0
	err = conn.Call("Rect.Area", Param{50, 100}, &ret)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println("area: ", ret)
}
