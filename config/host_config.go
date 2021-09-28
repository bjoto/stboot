// Copyright 2021 the System Transparency Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import (
	"net"
	"net/url"

	"github.com/vishvananda/netlink"
)

const HostConfigVersion int = 1

type IPAddressMode int

const (
	Static IPAddressMode = iota + 1
	Dynamic
)

func (n IPAddressMode) String() string {
	switch n {
	case Static:
		return "static"
	case Dynamic:
		return "dynamic"
	default:
		return "unknown"
	}
}

// HostCfg contains configuration data for a System Transparency host.
type HostCfg struct {
	Version          int
	NetworkMode      IPAddressMode
	HostIP           *netlink.Addr
	DefaultGateway   *net.IP
	DNSServer        *net.IP
	NetworkInterface *net.HardwareAddr
	ProvisioningURLs []*url.URL
	ID               string
	Auth             string
}

func LoadHostCfg() *HostCfg {
	return nil
}
