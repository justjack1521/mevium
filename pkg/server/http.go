package server

import (
	"fmt"
	"net/http"
)

func RunHTTPServer(port string, register func(server *http.Server)) {
	addr := fmt.Sprintf(":%s", port)
	RunHTTPServerOnAddr(addr, register)
}

func RunHTTPServerOnAddr(addr string, registerServer func(server *http.Server)) {

	svr := &http.Server{Addr: addr}

	registerServer(svr)

	if err := svr.ListenAndServe(); err != nil {
		panic(err)
	}

}
