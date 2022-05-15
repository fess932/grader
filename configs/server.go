package configs

type ServerConfig struct {
	NatsHost string
	NatsPort int
	Host     string
}

func NewServerConfig() ServerConfig {
	return ServerConfig{
		NatsHost: "localhost",
		NatsPort: 4488,
		Host:     ":8080",
	}
}
