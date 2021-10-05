package config

import (
	"encoding/json"
	"fmt"
	"io"
)

const (
	VersionJSONKey          = "version"
	NetworkModeJSONKey      = "network_mode"
	HostIPJSONKey           = "host_ip"
	DefaultGatewayJSONKey   = "gateway"
	DNSServerJSONKey        = "dns"
	NetworkInterfaceJSONKey = "network_interface"
	ProvisioningURLsJSONKey = "provisioning_urls"
	IdJSONKey               = "identity"
	AuthJSONKey             = "authentication"
)

type JSONParser struct {
	r io.Reader
}

func (p *JSONParser) Parse() (*HostCfg, error) {
	jsonBlob, err := io.ReadAll(p.r)
	if err != nil {
		return nil, err
	}

	var m map[string]interface{}
	if err = json.Unmarshal(jsonBlob, &m); err != nil {
		return nil, err
	}

	if val, ok := m["version"]; ok {
		ver, ok := val.(float64)
		if !ok {
			return nil, fmt.Errorf("for key 'version': want number, got %T", val)
		}
		return &HostCfg{Version: int(ver)}, nil
	}

	return &HostCfg{}, nil
}
