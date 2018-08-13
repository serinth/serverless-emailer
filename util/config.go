package util

import (
	"github.com/BurntSushi/toml"
	"github.com/caarlos0/env"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

type Config struct {
	Stage                        string `env:"STAGE,required"`
	HystrixTimeout               int    `env:"HYSTRIX_TIMEOUT"`
	HystrixMaxConcurrentRequests int    `env:"HYSTRIX_MAX_CONCURRENT_REQUESTS"`
	HystrixErrorPercentThreshold int    `env:"HYSTRIX_ERROR_THRESHOLD"`
	MetricsCommandName           string `env:"METRICS_COMMAND_NAME"`
	SendGridAPIKey               string `env:"SENDGRID_API_KEY,required"`
	SendGridURL                  string `env:"SENDGRID_URL"`
	MailGunAPIKey                string `env:"MAILGUN_API_KEY,required"`
	MailGunURL                   string `env:"MAILGUN_URL"`
}

func LoadConfig() *Config {

	initialEnv := os.Getenv("STAGE")
	if len(strings.TrimSpace(initialEnv)) == 0 {
		log.Fatal("'STAGE' variable not set, exiting.")
	}
	cfg := Config{}

	if _, err := toml.DecodeFile("configs/"+initialEnv+".toml", &cfg); err != nil {
		log.Fatalf("Could not load %s config with error: %s", err.Error())
	}

	err := env.Parse(&cfg)

	if err != nil {
		log.Fatalf("Failed to load env variables. %+v\n", err)
	}

	return &cfg
}
