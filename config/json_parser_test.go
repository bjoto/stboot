package config

import (
	"bytes"
	"reflect"
	"testing"
)

func TestJSONParser(t *testing.T) {
	json := `{ "version" :1 }`
	j := JSONParser{bytes.NewBufferString(json)}
	want := &HostCfg{Version: 1}

	got, err := j.Parse()

	assertNoError(t, err)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
