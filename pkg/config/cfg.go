package config

import (
	"bytes"

	"gopkg.in/yaml.v2"
)

// ProxyMap 转发表
type ProxyMap struct {
	Host   string `yaml:"host"`
	Target string `yaml:"target"`
	Scheme string `yaml:"scheme"`
}

// Home 主站端口
type Home struct {
	Port int `yaml:"port"`
}

// Configure 配置
type Configure struct {
	Home      *Home       `yaml:"home"`
	ProxyMaps []*ProxyMap `yaml:"proxy"`
}

// Parse 加载配置信息
func Parse(data []byte) (*Configure, error) {
	cfg := &Configure{}
	if err := yaml.NewDecoder(bytes.NewReader(data)).
		Decode(&cfg); err != nil {
		// log.Fatal(err)
		return nil, err
	}
	return cfg, nil
}
