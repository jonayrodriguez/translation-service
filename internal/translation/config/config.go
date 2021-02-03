package config

import (
	"github.com/jonayrodriguez/translation-service/internal/log"
)

// Config represents application configuration
type Config struct {
	Service ServiceConfig
	Server  ServerConfig
	Logging log.Logging
	DB      DBConfig
}

// ServiceConfig represents the service configuration
type ServiceConfig struct {
	Name string
}

// ServerConfig represents the server configuration
type ServerConfig struct {
	Port int
}

// DBConfig represents the DB configuration
type DBConfig struct {
	Host     string
	Port     int
	Schema   string
	Username string
	Password string
}
