package apiserver

type APIServer struct {
	config *Config
}

// New ...
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
	}
}

// Bootstrap
func (s *APIServer) Bootstrap() error {
	return nil
}
