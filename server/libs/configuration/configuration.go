package configuration

import (
	"log"
)

type AppConfig struct {
	DatabaseURL         string `envconfig:"DATABASE_URL"`
	LogDestination      string `envconfig:"LOG_DEST" default:"stdout"`
	LogFormat           string `envconfig:"LOG_FORMAT" default:"text"`
	LogLevel            string `envconfig:"LOG_LEVEL" default:"info"`
	HTTPPort            string `envconfig:"HTTPPort" default:"15123"`
	GithubPersonalToken string `envconfig:"GITHUB_PERSONAL_TOKEN"`
}

var conf *AppConfig

func init() {
	var c AppConfig
	err := Process("", &c)
	if err != nil {
		log.Fatalln("Failed getting environment variables.", err.Error())
	}
	conf = &c
}

func Load() *AppConfig {
	return conf
}
