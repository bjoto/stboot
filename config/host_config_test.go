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
	assertError(t, err, ErrHostCfgInvalid)
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("expected an error")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
