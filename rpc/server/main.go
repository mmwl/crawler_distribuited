package main

import (
	"net/rpc"
	"awesomeProject/rpc"
	"net"
	"fmt"
	"net/rpc/jsonrpc"
)

func main() {
	rpc.Register(rpcdemo.DemoService{})

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		fmt.Println(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		go jsonrpc.ServeConn(conn)
	}

}
