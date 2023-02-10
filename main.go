package main

import (
	"github.com/chuccp/httpProxy/proxy"
	"github.com/chuccp/utils/io"
	"os"
	"strconv"
)

func main() {
	export := os.Getenv("PORT")
	if export ==""{
		export = "20400"
	}
	port,_:=strconv.Atoi(export)
	tcp := io.NewTCPServer(port)
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
