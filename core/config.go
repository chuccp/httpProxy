package core

import "github.com/tidwall/gjson"

type Network uint8
type Security uint8

const (
	TCP Network = iota
	WS
)
type InboundConfig struct {
	Network Network
	InboundConfig  *gjson.Result
}

func NewInboundConfig(result *gjson.Result)*InboundConfig {
	str := result.Get("streamSettings.network").Str
	if str=="ws" {
		return &InboundConfig{Network: WS}
	}
	return &InboundConfig{Network: TCP}
}

