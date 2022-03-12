package main

import (
	"github.com/chuccp/httpProxy/net"
	"github.com/chuccp/httpProxy/proxy"
)

func main() {
	tcp := net.NewTCPServer(8080)
	err := tcp.Bind()
	if err == nil {
		for {
			stream, err1 := tcp.Accept()
			if err1 == nil {
				conn := proxy.NewConn(stream)
				go conn.Handle()
			} else {
				panic(err1)
			}
		}
	} else {
		panic(err)
	}
}
