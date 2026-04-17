package environment

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/creasty/defaults"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	APP_NAME      string `mapstructure:"APP_NAME" default:"go-portfolio"`
	APP_ENV       string `mapstruture:"APP_ENV" default:"local"`
	APP_SECURE    bool   `mapstructure:"APP_SECURE" default:"false"`
	APP_DEBUG     bool   `mapstructure:"APP_DEBUG" default:"true"`
	APP_HOST      string `mapstructure:"APP_HOST" default:"localhost"`
	APP_PROTOCOL  string `mapstructure:"APP_PROTOCOL" default:"tcp"`
	APP_HTTP_PORT int    `mapstructure:"APP_HTTP_PORT" default:"3000"`

	DB_HOST     string `mapstructure:"DB_HOST" default:"localhost"`
	DB_PORT     int    `mapstructure:"DB_PORT" default:"5432"`
	DB_NAME     string `mapstructure:"DB_NAME" default:""`
	DB_USER     string `mapstructure:"DB_USER" default:""`
	DB_PASSWORD string `mapstructure:"DB_PASSWORD" default:""`
	DB_MAX_IDLE int    `mapstructure:"DB_MAX_IDLE" default:"10"`
	DB_MAX_OPEN int    `mapstructure:"DB_MAX_OPEN" default:"100"`

	RATE_LIMITER_RPS   float64 `mapstructure:"RATE_LIMITER_RPS" default:"20"`
	RATE_LIMITER_BURST int     `mapstructure:"RATE_LIMITER_BURST" default:"5"`
	ALLOWED_ORIGINS    string  `mapstructure:"ALLOWED_ORIGINS"`

	// Add more configuration fields as needed
}

type Environment string

const (
	LOCAL Environment = "local"
	DEV   Environment = "development"
	PROD  Environment = "production"
)

func ProvideConfig() (*Config, error) {
	cfg := Config{}
	if err := defaults.Set(&cfg); err != nil {
		log.Fatalf("Error setting default values: %v", err)
	}

	// Detect environment
	envStr := strings.ToLower(os.Getenv("APP_ENV"))
	env := Environment(envStr)
	if env == "" {
		env = LOCAL
	}

	// Local → load .env file first
	if env == LOCAL {
		_ = godotenv.Load(".env")
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Auto-bind each struct field by key
	t := reflect.TypeOf(cfg)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		key := strings.ToUpper(field.Name)
		if err := viper.BindEnv(key); err != nil {
			return nil, fmt.Errorf("Failed to bind env %s : %w", key, err)
		}
	}

	// Fill struct
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}
