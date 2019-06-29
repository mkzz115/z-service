package confutil

import "github.com/BurntSushi/toml"

type Config struct {
	Server  ServerConfig      `toml:"server"`
	Client  ClientConfig      `toml:"thrift_client"`
	Account ServerConfig      `toml:"account"`
	Forum   ServerConfig      `toml:"forum"`
	Oneconf ServerConfig      `toml:"oneconf"`
	Order   ServerConfig      `toml:"order"`
	Depends DependServeConfig `toml:"depends_serve"`
	Logger  LogConfig         `toml:"logger"`
}

type ServerConfig struct {
	Addr string `toml:"addr"`
}

type LogConfig struct {
	Path string `toml:"path"`
}

type ClientConfig struct {
	Addr    string `toml:"addr"`
	Timeout int32  `toml:"timeout"`
}

type DependServeConfig struct {
	MysqlConfig MysqlConfig `toml:"mysql_config"`
}

type MysqlConfig struct {
	Dsn string `toml:"dsn"`
}

func FromToml(path string) (*Config, error) {
	conf := &Config{}
	_, err := toml.DecodeFile(path, conf)
	return conf, err
}
