package sysconf

type genericCfg map[string]interface{}

type ConfigLoader interface {
	Load() (genericCfg, error)
}

type cfgField struct {
	name      string
	jsonTag   string
	validator func(interface{}) (interface{}, error)
}

type cfgFields = []cfgField
