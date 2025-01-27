// Copyright 2021 the System Transparency Authors. All rights reserved
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package host

import (
	"fmt"
	"time"

	"github.com/system-transparency/stboot/stlog"
	"github.com/u-root/u-root/pkg/mount"
	"github.com/u-root/u-root/pkg/mount/block"
)

const (
	DataPartitionFSType     = "ext4"
	DataPartitionLabel      = "STDATA"
	DataPartitionMountPoint = "data"
	BootPartitionFSType     = "vfat"
	BootPartitionLabel      = "STBOOT"
	BootPartitionMountPoint = "boot"
)

// Files at STBOOT partition
const (
	HostConfigFile = "/host_configuration.json"
)

// Files at STDATA partition
const (
	TimeFixFile        = "stboot/etc/system_time_fix"
	CurrentOSPkgFile   = "stboot/etc/current_ospkg_pathname"
	LocalOSPkgDir      = "stboot/os_pkgs/local/"
	LocalBootOrderFile = "stboot/os_pkgs/local/boot_order"
	NetworkOSpkgCache  = "stboot/os_pkgs/cache"
)

func MountBootPartition() error {
	return mountPartitionRetry(BootPartitionLabel, BootPartitionFSType, BootPartitionMountPoint, 8, 1)
}

func MountDataPartition() error {
	return mountPartitionRetry(DataPartitionLabel, DataPartitionFSType, DataPartitionMountPoint, 8, 1)
}

func mountPartitionRetry(label, fsType, mountPoint string, retries, retryWait uint) error {
	if retries == 0 {
		retries = 1
	}
	var err error = nil
	for i := uint(0); i < retries; i++ {
		err := MountPartition(label, fsType, mountPoint)
		if err == nil {
			break
		}
		time.Sleep(time.Second * time.Duration(retryWait))
		stlog.Debug("Failed to mount %s to %s, retry %v", label, mountPoint, i + 1)
	}
	return err
}

func MountPartition(label, fsType, mountPoint string) error {
	devs, err := block.GetBlockDevices()
	if err != nil {
		return fmt.Errorf("host storage: %v", err)
	}

	devs = devs.FilterPartLabel(label)
	if len(devs) == 0 {
		return fmt.Errorf("host storage: no partition with label %s", label)
	}
	if len(devs) > 1 {
		return fmt.Errorf("host storage: multiple partitions with label %s", label)
	}

	d := devs[0].DevicePath()
	mp, err := mount.Mount(d, mountPoint, fsType, "", 0)
	if err != nil {
		return fmt.Errorf("host storage: %v", err)
	}

	stlog.Debug("Mounted device %s at %s", mp.Device, mp.Path)
	return nil
}
