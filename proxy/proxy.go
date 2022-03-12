package proxy

import (
	"bytes"
	"errors"
	"github.com/chuccp/httpProxy/net"
	"github.com/chuccp/utils/log"
	"strings"
)

const (
	GET     = "GET"
	POST    = "POST"
	CONNECT = "CONNECT"
)

type Header struct {
	method  string
	url     string
	version string
	header  map[string]string
}

func newHeader() *Header {
	return &Header{header: make(map[string]string)}
}

func (h *Header) Bytes(buff *bytes.Buffer) {
	buff.WriteString(h.method + " " + h.url + " " + h.version + "\r\n")
	for key, value := range h.header {
		buff.WriteString(key + ": " + value + "\r\n")
	}
	buff.WriteString("\r\n")
}

type Stream struct {
	header *Header
	stream *net.NetStream
}

func NewStream(stream *net.NetStream) *Stream {
	return &Stream{stream: stream, header: newHeader()}
}
func (s *Stream) ParseQueryUrl() error {
	data, err := s.stream.ReadLineLimit(8192)
	if err != nil {
		return err
	}
	query := string(data)
	queryUrl := strings.Split(query, " ")
	if len(queryUrl) == 3 {
		s.header.method = queryUrl[0]
		s.header.url = queryUrl[1]
		s.header.version = queryUrl[2]
		return nil
	} else {
		return errors.New("format error")
	}
}
func (s *Stream) ParseHeader() error {
	for {
		data, err2 := s.stream.ReadLineLimit(3145728)
		if err2 != nil {
			return err2
		}
		if len(data) == 0 {
			break
		} else {
			log.Info(string(data))
			kv := strings.SplitN(string(data), ":", 2)
			if len(kv) == 2 {
				s.header.header[strings.TrimSpace(kv[0])] = strings.TrimSpace(kv[1])
			}
		}
	}

	return nil
}
func (s *Stream) Proxy() error {
	err := s.ParseQueryUrl()
	if err != nil {
		return err
	}
	err1 := s.ParseHeader()
	if err1 != nil {
		return err1
	}
	if s.header.method == CONNECT {
		https := NewHttps(s.stream, s.header)
		err3 := https.Conn()
		if err3 == nil {
			_, err6 := s.stream.WriteAndFlush([]byte("HTTP/1.1 200 Connection Established\r\n\r\n"))
			if err6 != nil {
				return err6
			} else {
				https.Start()
			}
		} else {
			return err3
		}
	} else {
		http := NewHttp(s.stream, s.header)
		err3 := http.Conn()
		if err3 == nil {
			http.Start()
		} else {
			log.Info("!!!!", err3)
		}
	}

	return nil
}
