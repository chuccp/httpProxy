package core

import "github.com/chuccp/httpProxy/config"

type Register struct {
	context *Context
}

func (s *Register) Start()  {





}

func NewRegister(config *config.Config)*Register  {
	return &Register{context:newContext(config)}
}
