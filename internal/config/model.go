package config

import (
	"github.com/alexedwards/scs/v2"
)

type Config struct {
	Port    string
	Env     string
	Session *scs.SessionManager
}
