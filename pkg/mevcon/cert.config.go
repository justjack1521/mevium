package mevcon

type CertConfig struct {
	ServerCert string `required:"true"`
	ServerKey  string `required:"true"`
}
