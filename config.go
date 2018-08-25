package langmap

import (
	"encoding/json"
	"os"
)

type Config struct {
	Address  string
	Database DatabaseConfig
}

type DatabaseConfig struct {
	Driver string
	Source string
}

func (c *Config) FromFile(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}

	err = json.NewDecoder(f).Decode(c)
	if err != nil {
		return err
	}

	return nil
}
