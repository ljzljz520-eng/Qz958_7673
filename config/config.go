package config

import "os"

type Config struct {
	Address     string
	Database    string
	Environment string
}

func Load() Config {
	c := Config{Address: ":8080", Database: "bakery.db", Environment: "development"}
	if v := os.Getenv("BAKERY_ADDR"); v != "" {
		c.Address = v
	}
	if v := os.Getenv("BAKERY_DB"); v != "" {
		c.Database = v
	}
	if v := os.Getenv("BAKERY_ENV"); v != "" {
		c.Environment = v
	}
	return c
}
func (c Config) Production() bool { return c.Environment == "production" }
func (c Config) Validate() error {
	if c.Address == "" || c.Database == "" {
		return os.ErrInvalid
	}
	return nil
}
