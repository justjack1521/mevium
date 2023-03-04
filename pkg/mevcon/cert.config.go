package mevcon

import (
	"crypto/tls"
	"google.golang.org/grpc/credentials"
)

type CertConfig struct {
	ServerCert string `required:"true"`
	ServerKey  string `required:"true"`
}

func (c CertConfig) NewTransportCredentials() credentials.TransportCredentials {
	cert, err := tls.LoadX509KeyPair(c.ServerCert, c.ServerKey)
	if err != nil {
		panic(err)
	}
	return credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.NoClientCert,
	})
}
