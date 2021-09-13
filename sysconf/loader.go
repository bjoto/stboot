package sysconf

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/system-transparency/stboot/host"
)

// PartitionLoader loads a configuration from a JSON file from a partition
// with a certain label
type PartitionLoader struct {
	label string
	path  string
}

func (l *PartitionLoader) Load() (genericCfg, error) {
	mp, err := os.MkdirTemp("", "")
	if err != nil {
		return nil, fmt.Errorf("PartitionLoader: %v", err)
	}
	err = host.MountPartition(l.label, l.path, mp, 60)
	if err != nil {
		return nil, fmt.Errorf("PartitionLoader: %v", err)
	}
	p := filepath.Join(mp, HostConfigFileName)
	buf, err := ioutil.ReadFile(p)
	if err != nil {
		return nil, fmt.Errorf("PartitionLoader: %v", err)
	}
	var g genericCfg
	err = json.Unmarshal(buf, &g)
	if err != nil {
		return nil, fmt.Errorf("PartitionLoader: %v", err)
	}
	return g, nil
}
