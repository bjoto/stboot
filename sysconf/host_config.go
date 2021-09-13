// Copyright 2021 the System Transparency Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sysconf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/url"

	"github.com/system-transparency/stboot/stlog"
	"github.com/vishvananda/netlink"
)

const (
	HostConfigVersion  int    = 1
	HostConfigFileName string = "host_configuration.json"
)

type IPConfiguration int

const (
	StaticIP IPConfiguration = iota + 1
	DynamicIP
)

func (n IPConfiguration) String() string {
	return []string{"static", "dynamic"}[n]
}

const (
	VersionJSONKey          = "version"
	NetworkModeJSONKey      = "network_mode"
	HostIPJSONTag           = "host_ip"
	DefaultGatewaJSONKey    = "gateway"
	DNSServerJSONKey        = "dns"
	NetworkInterfaceJSONKey = "network_interface"
	ProvisioningURLsJSONKey = "provisioning_urls"
	IDJSONKey               = "identity"
	AuthJSONKey             = "authentication"
)

// HostCfg contains configuration data for a System Transparency host.
type HostCfg struct {
	Version          int
	NetworkMode      IPConfiguration
	HostIP           *netlink.Addr
	DefaultGateway   *net.IP
	DNSServer        *net.IP
	NetworkInterface *net.HardwareAddr
	ProvisioningURLs []*url.URL
	ID               string
	Auth             string
}

var hostCfgFields = cfgFields{
	cfgField{
		name:      "Version",
		jsonTag:   VersionJSONKey,
		validator: nil,
	}, {
		name:      "NetworkMode",
		jsonTag:   NetworkInterfaceJSONKey,
		validator: validateTest,
	}, {
		name:      "HostIP",
		jsonTag:   HostIPJSONTag,
		validator: validateTest,
	}, {
		name:      "DefaultGateway",
		jsonTag:   DNSServerJSONKey,
		validator: validateTest,
	}, {
		name:      "DNSServer",
		jsonTag:   DNSServerJSONKey,
		validator: validateTest,
	}, {
		name:      "NetworkInterface",
		jsonTag:   NetworkInterfaceJSONKey,
		validator: validateTest,
	}, {
		name:      "ProvisioningURLs",
		jsonTag:   ProvisioningURLsJSONKey,
		validator: validateTest,
	}, {
		name:      "ID",
		jsonTag:   IDJSONKey,
		validator: validateTest,
	}, {
		name:      "Auth",
		jsonTag:   AuthJSONKey,
		validator: validateTest,
	},
}

func ParseHostCfg(cfg genericCfg) (*HostCfg, error) {
	var hc = new(HostCfg)
	validated := make(map[string]interface{})

	for _, field := range hostCfgFields {
		rawValue, ok := cfg[field.jsonTag]
		if !ok {
			return nil, fmt.Errorf("missing field: %s", VersionJSONKey)
		}
		validValue, err := field.validator(rawValue)
		if err != nil {
			return nil, fmt.Errorf("validator: %v", err)
		}
		validated[field.name] = validValue
	}

	hc.Version = validated["Version"].(int)
	hc.NetworkMode = validated["NetworkMode"].(IPConfiguration)
	hc.HostIP = validated["HostIP"].(*netlink.Addr)
	hc.DefaultGateway = validated["DefaultGateway"].(*net.IP)
	hc.DNSServer = validated["DNSServer"].(*net.IP)
	hc.NetworkInterface = validated["NetworkInterface"].(*net.HardwareAddr)
	hc.ProvisioningURLs = validated["ProvisioningURLs"].([]*url.URL)
	hc.ID = validated["ID"].(string)
	hc.Auth = validated["Auth"].(string)

	return hc, nil
}

func validateTest(interface{}) (interface{}, error) {
	fmt.Print("Validation ongoing")
	return nil, nil
}
