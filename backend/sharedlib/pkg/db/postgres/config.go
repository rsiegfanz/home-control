package postgres

type Config struct {
	Host     string `env:"HOST"`
	User     string `env:"USER"`
	Password string `env:"PASSWORD"`
	DbName   string `env:"DBNAME"`
	Port     int16  `env:"PORT"`
}
