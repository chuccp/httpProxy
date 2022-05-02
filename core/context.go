package core

import "github.com/chuccp/httpProxy/config"

type Context struct {
	config  *config.Config
}

func (c *Context) InboundConfigs() []*InboundConfig {
	inbounds:=make([]*InboundConfig,0)
	array:=c.config.Result().Get("inbounds").Array()
	for _, result := range array {
		inbounds = append(inbounds, NewInboundConfig(&result))
	}
	return inbounds
}



func newContext(config *config.Config) *Context {
	return &Context{ config: config}
}