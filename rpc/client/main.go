package main

import (
	"net"
	"net/rpc/jsonrpc"
	"awesomeProject/rpc"
	"fmt"
)

func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	client := jsonrpc.NewClient(conn)

	var result float64

	client.Call("DemoService.Div",
		rpcdemo.Args{10, 3}, &result)

	fmt.Println(result)

	err = client.Call("DemoService.Div",
		rpcdemo.Args{10, 0}, &result)

	if err != nil {
		fmt.Println(err)
	}else {
		fmt.Println(result)
	}

	var a1 int

	client.Call("DemoService.Jia",rpcdemo.Args{10,20},&a1)

	fmt.Println(a1)
}
