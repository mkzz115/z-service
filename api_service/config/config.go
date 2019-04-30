package config

import "github.com/BurntSushi/toml"

type Config struct {
	Server HttpConfig `toml:"server"`
	Logger LogConfig  `toml:"logger"`
}

type HttpConfig struct {
	Addr string `toml:"addr"`
}

type LogConfig struct {
	Path string `toml:"path"`
}

func FromToml(path string) (*Config, error) {
	conf := &Config{}
	_, err := toml.DecodeFile(path, conf)
	return conf, err
}
