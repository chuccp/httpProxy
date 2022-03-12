package proxy

import (
	"bytes"
	"github.com/chuccp/utils/io"
	net2 "net"
	"strings"
)

type Http struct {
	local  *io.NetStream
	header *Header
	remote *io.NetStream
}

func NewHttp(stream *io.NetStream,header *Header) *Http {
	return &Http{local: stream, header: header}
}

func (h *Http) Conn() (err error) {
	host := h.header.header["Host"]
	if strings.Index(host,":")<0{
		host = host+":80"
	}
	address, err := net2.ResolveTCPAddr("tcp", host)
	if err != nil {
		return err
	}
	client := io.NewXConn(address.IP.String(), address.Port)
	h.remote, err = client.Create()
	if err != nil {
		return err
	}
	return nil
}

func (h *Http)replayOK()error{
	_, err6 := h.local.WriteAndFlush([]byte("HTTP/1.1 200 Connection Established\r\n\r\n"))
	return err6
}

func (h *Http)writeHeader()error{
	var buff = new(bytes.Buffer)
	h.header.Bytes(buff)
	_, err6 := h.remote.WriteAndFlush(buff.Bytes())
	return err6
}


func (h *Http) Start() {
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
