package configs

type GraderConfig struct {
	Addr string
}

func NewGraderConfig() *GraderConfig {
	return &GraderConfig{
		Addr: "localhost:8081",
	}
}
