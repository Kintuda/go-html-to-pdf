package config

type AppConfig struct {
	Env         string `env:"ENV"`
	HttpPort    string `env:"HTTP_PORT"`
	PostgresDns string `env:"POSTGRES_DNS"`
}
