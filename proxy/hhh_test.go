package proxy

import (
	"strings"
	"testing"
)

func TestName(t *testing.T) {

	kv := strings.SplitN("kk",":",2)
	for _, s := range kv {
		println(strings.TrimSpace(s))
	}
	
}
