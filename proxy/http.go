package proxy

import (
	"bytes"
	"errors"
	"github.com/chuccp/httpProxy/net"
	net2 "net"
	"strings"
)

type Http struct {
	header *Header
	local  *net.NetStream
	remote *net.NetStream
}

func NewHttp(stream *net.NetStream,header *Header)*Http  {

	return &Http{local:stream,header: header}
}
func (http *Http) Conn() (err error){

	host := http.header.header["Host"]
	if strings.Index(host,":")<0{
		host = host+":80"
	}
	address, err := net2.ResolveTCPAddr("tcp", host)
	if err != nil {
		return err
	}
	client := net.NewXConn(address.IP.String(), address.Port)
	http.remote, err = client.Create()
	if err != nil {
		return err
	}
	return nil
}

func (http *Http) ParseQueryUrl() error {
	data, err := http.local.ReadLineLimit(8192)
	if err != nil {
		return err
	}
	query := string(data)
	queryUrl := strings.Split(query, " ")
	if len(queryUrl) == 3 {
		http.header.method = queryUrl[0]
		http.header.url = queryUrl[1]
		http.header.version = queryUrl[2]
		return nil
	} else {
		return errors.New("format error")
	}
}
func (http *Http) ParseHeader() error {
	for {
		data, err2 := http.local.ReadLineLimit(3145728)
		if err2 != nil {
			return err2
		}
		if len(data) == 0 {
			break
		} else {
			kv := strings.SplitN(string(data), ":", 2)
			if len(kv) == 2 {
				http.header.header[strings.TrimSpace(kv[0])] = strings.TrimSpace(kv[1])
			}
		}
	}

	return nil
}

func (http *Http) Start()  {
	var buff = new(bytes.Buffer)
	http.header.Bytes(buff)
	http.remote.WriteAndFlush(buff.Bytes())
	http.local.ReadFunc(func(data []byte) bool {
		_, err2 := http.remote.WriteAndFlush(data)
		if err2 == nil {
			return true
		}
		return false
	}, func() {
		http.remote.ManualClose()
	})
	http.remote.ReadFunc(func(data []byte) bool {
		_, err2 := http.local.WriteAndFlush(data)
		if err2 == nil {
			return true
		}
		return false
	}, func() {
		http.local.ManualClose()
	})
}
