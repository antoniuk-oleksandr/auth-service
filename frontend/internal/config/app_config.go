package config

type BackendConfig struct {
	Addr string `env:"BACKEND_ADDR" validate:"required"`
}

type ServerConfig struct {
	Port string `env:"SERVER_PORT" envDefault:"3000" validate:"required"`
}

type AppConfig struct {
	BackendConfig BackendConfig
	ServerConfig  ServerConfig
}
