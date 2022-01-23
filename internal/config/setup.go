package config

func Load() Config {
	return Config{
		Port: "8080",
		Env:  "dev",
	}
}
