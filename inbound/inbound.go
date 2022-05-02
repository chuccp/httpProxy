package inbound

import (
	"github.com/chuccp/httpProxy/core"
)

type Inbound struct {
	inboundConfigs []*core.InboundConfig
}

func (i *Inbound) Init(ctx *core.Context)  {
	i.inboundConfigs = ctx.InboundConfigs()
}
func (i *Inbound) Start() error {

	for _, config := range i.inboundConfigs {

		switch config.Network {

		case core.TCP:

		case core.WS:


		}
	}

	return nil
}


