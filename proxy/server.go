package proxy

import (
	"github.com/chuccp/utils/io"
	"github.com/chuccp/utils/log"
	"strings"
)

type Proxy struct {
	server *io.TCPServer
}

func NewProxy(port int)*Proxy  {
	server:=io.NewTCPServer(port)
	return &Proxy{server: server}
}
func (proxy*Proxy)Start()  {
	err:=proxy.server.Bind()
	if err==nil{
		go func() {
			stream,err:=proxy.server.Accept()
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
}