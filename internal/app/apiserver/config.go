package apiserver

// Config ...
type Config struct {
	BindAddress string `toml:"bind_address"`
}

func NewConfig() *Config {
	return &Config{
		BindAddress: ":8080",
	}
}
