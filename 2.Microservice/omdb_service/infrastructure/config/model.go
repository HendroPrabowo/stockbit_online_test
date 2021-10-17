package config

type AppConfig struct {
	ServiceName string `mapstructure:"service_name"`
	Server      Server `mapstructure:"server"`
	Omdb        Omdb   `mapstructure:"omdb"`
}

type Server struct {
	Port string
}

type Omdb struct {
	Url    string
	ApiKey string
}
