package configs

type ServerConfig struct {
	NatsHost string
	NatsPort int
	Host     string

	GraderAddr string

	QueueWorkers int
	FilesDir     string
}

func NewServerConfig() ServerConfig {
	return ServerConfig{
		NatsHost:     "localhost",
		NatsPort:     4488,
		Host:         ":8080",
		FilesDir:     "./files",
		QueueWorkers: 10,
		GraderAddr:   "localhost:8081",
	}
}
