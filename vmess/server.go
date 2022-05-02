package vmess

import "github.com/chuccp/httpProxy/core"

type Server struct {

}

func (s *Server) Start() error {
	return nil
}
func (s *Server) Init(context *core.Context) {



}
func (s *Server) Name() string {
	return "vmess"
}
