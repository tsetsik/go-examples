package config

import "github.com/go-playground/validator"

type Config struct {
	Host       string
	Port       int
	JobWorkers int
}

func LoadConfig(host string, port, jobWorkers int) (*Config, error) {
	return &Config{
		Host:       host,
		Port:       port,
		JobWorkers: jobWorkers,
	}, nil
}

func (c *Config) Validate() error {
	validator := validator.New()
	return validator.Struct(c)
}
