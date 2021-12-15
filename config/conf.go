package config

import "fmt"

// DefaultConfigFile contains the config path to be used if no file is supplied.
const DefaultConfigFile = "config.yml"

type Conf struct {
	Database DatabaseConf `mapstructure:"database"`
}

type DatabaseConf struct {
	User     string `mapstructure:"user"`
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Name     string `mapstructure:"name"`
	Port     string `mapstructure:"port"`
}

func (d *DatabaseConf) Uri() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?retryWrites=true&tls=false", d.User, d.Password, d.Host, d.Port, d.Name)
}
