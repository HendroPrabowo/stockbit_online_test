package config

type AppConfig struct {
	ServiceName string `mapstructure:"service_name"`
	Server      Server `mapstructure:"server"`
}

type Server struct {
	Port string
}
