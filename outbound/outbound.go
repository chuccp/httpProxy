package outbound

import "github.com/chuccp/utils/io"

type Outbound struct {
	remote *io.NetStream
	local  *io.NetStream
}

func (outbound *Outbound) Start() {
	outbound.remote.ReadFunc(func(data []byte) bool {
		_, err2 := outbound.local.WriteAndFlush(data)
		if err2 == nil {
			return true
		}
		return false
	}, func() {
		outbound.local.ManualClose()
	})
	outbound.local.ReadFunc(func(data []byte) bool {
		_, err2 := outbound.remote.WriteAndFlush(data)
		if err2 == nil {
			return true
		}
		return false
	}, func() {
		outbound.remote.ManualClose()
	})
}