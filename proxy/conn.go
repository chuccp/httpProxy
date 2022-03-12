package proxy

import (
	"github.com/chuccp/httpProxy/net"
)


type Conn struct {
	stream *net.NetStream
}

func NewConn(stream *net.NetStream) *Conn {
	return &Conn{stream: stream}
}
func (c *Conn) Handle()  {
	stream:=NewStream(c.stream)
	stream.Proxy()
}