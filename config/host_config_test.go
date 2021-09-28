package config

import (
	"reflect"
	"testing"
)

func TestLoadHostCfg(t *testing.T) {
	want := HostCfg{}
	got := LoadHostCfg()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %V, want %V", got, want)
	}
}
