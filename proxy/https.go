package proxy

import (
	"github.com/chuccp/httpProxy/net"
	net2 "net"
)

type Https struct {
	local  *net.NetStream
	header *Header
	remote *net.NetStream
}

func NewHttps(stream *net.NetStream,header *Header) *Https {
	return &Https{local: stream, header: header}
}

func (h *Https) Conn() (err error) {
	address, err := net2.ResolveTCPAddr("tcp", h.header.header["Host"])
	if err != nil {
		return err
	}
	client := net.NewXConn(address.IP.String(), address.Port)
	h.remote, err = client.Create()
	if err != nil {
		return err
	}
	return nil
}
func (h *Https) Start() {
	h.remote.ReadFunc(func(data []byte) bool {
		_, err2 := h.local.WriteAndFlush(data)
		if err2 == nil {
			return true
		}
		return false
	}, func() {
		h.local.ManualClose()
	})
	h.local.ReadFunc(func(data []byte) bool {
		_, err2 := h.remote.WriteAndFlush(data)
		if err2 == nil {
			return true
		}
		return false
	}, func() {
		h.remote.ManualClose()
	})
}
