package dto

type Config struct {
	Webserver Webserver `toml:"webserver"`
	Database  Database  `toml:"database"`
	Redis     Redis     `toml:"redis"`
}

type Webserver struct {
	Port string `toml:"port"`
}

type Database struct {
	Host     string `toml:"host"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	Port     string `toml:"port"`
	Schema   string `toml:"schema"`
}

type Redis struct {
	Host     string `toml:"host"`
	Port     string `toml:"Port"`
	Password string `toml:"Password"`
}
