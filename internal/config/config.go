package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

// Config 应用配置
type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Swagger  SwaggerConfig  `yaml:"swagger"`
	Database DatabaseConfig `yaml:"database"`
	Redis    RedisConfig    `yaml:"redis"`
}

// ServerConfig 服务器配置
type ServerConfig struct {
	Port        string `yaml:"port"`
	BasePath    string `yaml:"base_path"`
	Environment string `yaml:"environment"`
}

// SwaggerConfig Swagger 配置
type SwaggerConfig struct {
	Enabled     bool     `yaml:"enabled"`
	Title       string   `yaml:"title"`
	Version     string   `yaml:"version"`
	Description string   `yaml:"description"`
	Host        string   `yaml:"host"`
	Schemes     []string `yaml:"schemes"`
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

// RedisConfig Redis 配置
type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

// Load 从文件加载配置
func Load(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	// 设置默认值
	if cfg.Server.Port == "" {
		cfg.Server.Port = "8002"
	}
	if cfg.Server.BasePath == "" {
		cfg.Server.BasePath = "/api/v3"
	}
	if cfg.Server.Environment == "" {
		cfg.Server.Environment = "local"
	}

	return &cfg, nil
}

