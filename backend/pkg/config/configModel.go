package config

type Config struct {
	HouseServer struct {
		Url string `yaml:"url"`
	} `yaml:"houseServer"`
	DataPaths struct {
		LatestMeasurements string `yaml:"latestMeasurements"`
	} `yaml:"dataPaths"`
}
