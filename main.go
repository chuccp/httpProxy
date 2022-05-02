package main

import (
	"github.com/chuccp/httpProxy/config"
	"github.com/chuccp/httpProxy/core"
)

func main() {
	cfg, err := config.LoadFile("config.json")
	if err==nil{
		register:=core.NewRegister(cfg)
		register.Start()
	}
}
