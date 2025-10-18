package config

import "github.com/go-playground/validator"

type Config struct {
	Host       string `validate:"required,hostname|ip"`
	Port       int    `validate:"required,min=1,max=65535"`
	InfraPort  int    `validate:"required,min=1,max=65535"`
	JobWorkers int    `validate:"required,min=1,max=100"`
}

func LoadConfig(host string, port, infraPort, jobWorkers int) (*Config, error) {
	return &Config{
		Host:       host,
		Port:       port,
		InfraPort:  infraPort,
		JobWorkers: jobWorkers,
	}, nil
}

func (c *Config) Validate() error {
	validator := validator.New()
	return validator.Struct(c)
}
