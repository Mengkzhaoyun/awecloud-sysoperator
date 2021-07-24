package common

import "net"

type args struct {
	Port           int
	BindAddress    net.IP
	DefaultCertDir string
	CertFile       string
	KeyFile        string
	TokenTTL       int
	APIToken       string
}

// Args , cmd 初始化参数
var Args = &args{}
