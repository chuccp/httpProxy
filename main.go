package main

import (
	"github.com/chuccp/utils/io"
	"github.com/chuccp/utils/log"
	"strings"
	"time"
)

func main() {
	tcp:=io.NewTCPServer(8080)
	err:=tcp.Bind()
	if err==nil{
		go func() {
			stream,err:=tcp.Accept()
			if err==nil{
				for{
					data,err:=stream.ReadLine()
					if err==nil{
						query:=string(data)
						index:=strings.IndexByte(string(data),' ')
						protocol:=query[:index]
						log.Info("protocol",protocol)
					}
				}
			}else{
				log.Info(err)
			}
		}()
	}

	time.Sleep(time.Hour)

}
