package service

type Config struct {
	IP   string
	Port string
}

func InitConfig() *Config {
	// Default configuration
	return &Config{
		IP:   "0.0.0.0",
		Port: "3788",
	}
}
