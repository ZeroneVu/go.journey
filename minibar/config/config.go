package config

import (
	"log"

	"github.com/BurntSushi/toml"
)

// Config ...
type Config struct {
	Server Server `toml:"server"`
}

// Server ...
type Server struct {
	Port  int  `toml:"port"`
	Oauth bool `toml:"oauth"`
}

// NewConfig ...
// Parses the config file using a 3rd party toml parser
func NewConfig(dir string) Config {
	var conf Config

	// Parse the config toml file
	if _, err := toml.DecodeFile(dir, &conf); err != nil {
		log.Fatal(err)
	}

	// If the server config is nil then populate the port to the default
	if conf.Server.Port == 0 {
		conf.Server.Port = 8080
	}

	return conf
}
