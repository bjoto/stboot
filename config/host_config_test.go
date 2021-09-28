package config

import (
	"testing"
)

type StubHostCfgParser struct {
	hc *HostCfg
}

func (s StubHostCfgParser) Parse() (*HostCfg, error) {
	return s.hc, nil
}

func TestLoadHostCfg(t *testing.T) {
	p := StubHostCfgParser{&HostCfg{}}

	_, err := LoadHostCfg(p)
	if err == nil {
		t.Errorf("expected error")
	}
}
