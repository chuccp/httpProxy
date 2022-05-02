package config

import (
	"github.com/chuccp/utils/file"
	"github.com/tidwall/gjson"
	"io"
)

type Config struct {
	json []byte
	result *gjson.Result
}
func (c *Config) Result()*gjson.Result  {
	return c.result
}
func (c *Config) init()*Config  {
	var result = gjson.ParseBytes(c.json)
	c.result = &result
	return c
}
func New(json []byte)(*Config,error){
	c:=&Config{json: json}
	return c.init(), nil
}
func LoadFile(fileName string) (*Config, error) {
	f, err := file.NewFile(fileName)
	if err != nil {
		return nil, err
	}
	fi, err := f.ToRawFile()
	if err != nil {
		return nil, err
	}
	data, err := io.ReadAll(fi)
	if err != nil {
		return nil, err
	}
	return New(data)
}
