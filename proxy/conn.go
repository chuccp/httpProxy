package proxy

import (
	"github.com/chuccp/utils/io"
)


type Conn struct {
	stream *io.NetStream
}

func NewConn(stream *io.NetStream) *Conn {
	return &Conn{stream: stream}
}
func (c *Conn) Handle()  {
	stream:=NewStream(c.stream)
	stream.Proxy()
}