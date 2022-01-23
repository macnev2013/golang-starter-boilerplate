package config

import (
	"github.com/alexedwards/scs/v2"
)

type Config struct {
	Port string
	Env  string

	Session *scs.SessionManager
	DB      DBConfig
}

type (
	// DBConfig defines database configurations
	DBConfig struct {
		Host     string
		Port     string
		Password string
		Name     string
	}
)
