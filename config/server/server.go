package server

// ConfigServer add here and yml file if there is more configuration
type ConfigServer struct {
	Debug string `yaml:"debug"`
	Grpc  Server `yaml:"grpc"`
}

type Server struct {
	Host    string `yaml:"host"`
	Port    int    `yaml:"port"`
	TLS     bool   `yaml:"tls"`
	Timeout int    `yaml:"timeout"`
}
