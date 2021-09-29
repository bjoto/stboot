// Copyright 2021 the System Transparency Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import (
	"fmt"
	"net"
	"net/url"

	"github.com/vishvananda/netlink"
)

const HostCfgVersion int = 1

type InvalidError string

func (e InvalidError) Error() string {
	return string(e)
}

var (
	ErrVersionMissmatch  = InvalidError("version missmatch, want version " + fmt.Sprint(HostCfgVersion))
	ErrMissingIPAddrMode = InvalidError("IP address mode must be set")
	ErrUnknownIPAddrMode = InvalidError("unknown IP address mode")
	ErrMissingProvURLs   = InvalidError("provisioning server URL list must not be empty")
	ErrInvalidProvURLs   = InvalidError("missing or unsupported scheme in provisioning URLs")
	ErrMissingIPAddr     = InvalidError("IP address must not be empty when static IP mode is set")
	ErrMissingGateway    = InvalidError("default gateway must not be empty when static IP mode is set")
	ErrMissingID         = InvalidError("ID must not be empty when a URL contains '$ID'")
	ErrInvalidID         = InvalidError("invalid ID string, max 64 characters [a-z,A-Z,0-9,-,_]")
	ErrMissingAUTH       = InvalidError("Auth must not be empty when a URL contains '$AUTH'")
	ErrInvalidAUTH       = InvalidError("invalid auth string, max 64 characters [a-z,A-Z,0-9,-,_]")
)

type IPAddrMode int

const (
	UnsetIPAddrMode IPAddrMode = iota
	Static
	Dynamic
)

func (n IPAddrMode) String() string {
	switch n {
	case UnsetIPAddrMode:
		return "unset"
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
	NetworkMode      IPAddrMode
	HostIP           *netlink.Addr
	DefaultGateway   *net.IP
	DNSServer        *net.IP
	NetworkInterface *net.HardwareAddr
	ProvisioningURLs []*url.URL
	ID               string
	Auth             string
}

type HostCfgParser interface {
	Parse() (*HostCfg, error)
}

// LoadHostCfg returns a HostCfg using the provided pa
func LoadHostCfg(p HostCfgParser) (*HostCfg, error) {
	return nil, InvalidError("invalid config")
}
