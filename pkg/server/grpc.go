package server

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
)

func RunGRPCServer(port string, register func(server *grpc.Server)) {
	addr := fmt.Sprintf("0.0.0.0:%s", port)
	RunGRPCServerOnAddr(addr, register)
}

func RunGRPCServerOnAddr(addr string, registerServer func(server *grpc.Server)) {

	svr := grpc.NewServer()
	registerServer(svr)

	listen, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	err = svr.Serve(listen)
	if err != nil {
		panic(err)
	}

}
