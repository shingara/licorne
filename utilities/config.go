package utilities

import (
	"os"
)

type Config struct {
	Database string
	Host string
	Port string
}

var (
	CurrentConfig *Config
)

func GetConfig() *Config {
	if CurrentConfig == nil {
		c := &Config{ os.Getenv("MONGO_DATABASE"),
		os.Getenv("DB_1_PORT_27017_TCP_ADDR"),
		os.Getenv("DB_1_PORT_27017_TCP_PORT") }

		if c.Database == "" {
			c.Database = "licorne-dev"
		}
		if c.Host == "" {
			c.Host = "127.0.0.1"
		}

		if c.Port == ""{
			c.Port = "27017"
		}
		CurrentConfig = c
	}
	return CurrentConfig
}
