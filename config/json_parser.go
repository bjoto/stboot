package config

import "io"

type JSONParser struct {
	r io.Reader
}

func (j *JSONParser) Parse() (*HostCfg, error) {
	return nil, nil
}
