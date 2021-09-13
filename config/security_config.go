// Copyright 2021 the System Transparency Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

const SecurityConfigVersion int = 1

type BootMode int

// Bootmodes values defines where to load a OS package from.
const (
	LocalBoot BootMode = iota
	NetworkBoot
)

func (b BootMode) String() string {
	return []string{"local", "network"}[b]
}

// SecurityConfig contains ecurity critical configuration data for a System Transparency host.
type SecurityConfig struct {
	Version                 int
	ValidSignatureThreshold int
	BootMode                BootMode
	UsePkgCache             bool
}
