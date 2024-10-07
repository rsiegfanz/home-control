package configs

type FetcherConfig struct {
	Url   string `env:"URL"`
	Modus string `env:"MODUS"`
}
