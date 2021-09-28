package config

import (
	"reflect"
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

	want := &HostCfg{}
	got, _ := LoadHostCfg(p)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
