package util

import (
	"os"
	"testing"

	"github.com/caarlos0/env"
	"github.com/stretchr/testify/assert"
)

func TestMissingRequiredVars(t *testing.T) {
	cfg := Config{}

	os.Unsetenv("STAGE")

	err := env.Parse(&cfg)

	assert.Contains(t, err.Error(), "Required environment variable STAGE is not set")
}

func TestConfigStagePropertySet(t *testing.T) {
	os.Setenv("STAGE", "local")
	os.Setenv("SENDGRID_API_KEY", "key")
	os.Setenv("MAILGUN_API_KEY", "key")
	pwd, _ := os.Getwd()
	os.Chdir("../")
	defer func() { os.Chdir(pwd) }()

	cfg := LoadConfig()

	assert.Equal(t, "local", cfg.Stage)
	assert.Equal(t, 20000, cfg.HystrixTimeout)
}

func TestConfigEnvironmentOverridesFileProps(t *testing.T) {
	os.Setenv("STAGE", "local")
	os.Setenv("SENDGRID_API_KEY", "key")
	os.Setenv("MAILGUN_API_KEY", "key")
	os.Setenv("METRICS_COMMAND_NAME", "XXX")
	pwd, _ := os.Getwd()
	os.Chdir("../")
	defer func() { os.Chdir(pwd) }()

	cfg := LoadConfig()

	assert.Equal(t, "local", cfg.Stage)
	assert.Equal(t, "XXX", cfg.MetricsCommandName)
}
