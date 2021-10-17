package config

type AppConfig struct {
	ServiceName string   `mapstructure:"service_name"`
	Server      Server   `mapstructure:"server"`
	Omdb        Omdb     `mapstructure:"omdb"`
	Postgres    Postgres `mapstructure:"postgres"`
}

type Server struct {
	Port string
}

type Omdb struct {
	Url    string
	ApiKey string
}

type Postgres struct {
	Addr     string
	User     string
	Password string
	Database string
}
