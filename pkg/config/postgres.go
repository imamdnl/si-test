package config

import (
	"errors"
	"os"
)

type PostgresConfig struct {
	Host string
	Port string

	Username string
	Password string

	Database string
	Schema   string
}

func (cfg PostgresConfig) GetHost() string {
	return cfg.Host
}

func (cfg PostgresConfig) GetPort() string {
	return cfg.Port
}

func (cfg PostgresConfig) GetUsername() string {
	return cfg.Username
}

func (cfg PostgresConfig) GetPassword() string {
	return cfg.Password
}

func (cfg PostgresConfig) GetDatabase() string {
	return cfg.Database
}

func (cfg PostgresConfig) GetSchema() string {
	return cfg.Schema
}

func NewPostgresConfig() (PostgresConfig, error) {
	host := os.Getenv("DB_POSTGRES_HOST")
	if host == "" {
		return PostgresConfig{}, errors.New("postgres host is required")
	}

	port := os.Getenv("DB_POSTGRES_PORT")
	if port == "" {
		return PostgresConfig{}, errors.New("postgres port is required")
	}

	username := os.Getenv("DB_POSTGRES_USERNAME")
	if username == "" {
		return PostgresConfig{}, errors.New("postgres username is required")
	}

	password := os.Getenv("DB_POSTGRES_PASSWORD")
	if password == "" {
		Logger().Warn("warning, postgres password is empty")
	}

	database := os.Getenv("DB_POSTGRES_NAME")
	if database == "" {
		return PostgresConfig{}, errors.New("postgres name is required")
	}

	schema := os.Getenv("DB_POSTGRES_SCHEMA")
	if schema == "" {
		return PostgresConfig{}, errors.New("postgres schema is required")
	}

	return PostgresConfig{
		Host:     host,
		Port:     port,
		Username: username,
		Password: password,
		Database: database,
		Schema:   schema,
	}, nil
}
