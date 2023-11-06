package config

type Config struct {
	Database DatabaseConfig `yaml:"database"`
	System   System         `yaml:"system"`
}
