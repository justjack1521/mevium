package server

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"net"
)

func RunGRPCServer(port string, credentials credentials.TransportCredentials, register func(server *grpc.Server)) {
	addr := fmt.Sprintf("0.0.0.0:%s", port)
	RunGRPCServerOnAddr(addr, credentials, register)
}

func RunGRPCServerOnAddr(addr string, credentials credentials.TransportCredentials, registerServer func(server *grpc.Server)) {

	svr := grpc.NewServer(grpc.Creds(credentials))

	registerServer(svr)

	reflection.Register(svr)

	listen, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	fmt.Println(fmt.Sprintf("Now listening on: %s", listen.Addr()))

	err = svr.Serve(listen)
	if err != nil {
		panic(err)
	}

}
