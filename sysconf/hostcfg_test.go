package sysconf

import (
	"reflect"
	"testing"
)

func TestParseHostCfg(t *testing.T) {
	var c = genericCfg{
		"version": "eins",
	}
	var want = &HostCfg{
		Version: 1,
	}
	got, err := ParseHostCfg(c)
	if err != nil {
		t.Fatalf("Something went wrong: %v", err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("Parsing failed\n got %+v\nwant %+v", got, want)
	}
}
