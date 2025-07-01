package swissknife

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

// SaveStructToJSONFile save structure to JSON file
func SaveStructToJSONFile(i interface{}, filepath string) error {
	jsonr, err := json.Marshal(i)
	if err != nil {
		return err
	}

	return SaveStringToFile(filepath, string(jsonr))
}

// SaveStructToJSONFileIndent save structure to JSON file
func SaveStructToJSONFileIndent(i interface{}, filepath string) error {
	jsonr, err := json.MarshalIndent(i, "", "	")
	if err != nil {
		return err
	}

	return SaveStringToFile(filepath, string(jsonr))
}

// SaveStructToYamlFile save structure to YAML file (in plain text)
func SaveStructToYamlFile(i interface{}, filepath string) error {
	dataEncoded, err := yaml.Marshal(i)
	if err != nil {
		return err
	}

	return SaveStringToFile(filepath, string(dataEncoded))
}

type DBConfig struct {
	Host          string `envconfig:"DB_HOST" default:"localhost"`
	Port          int    `envconfig:"DB_PORT" default:"3306"`
	Name          string `envconfig:"DB_NAME"`
	User          string `envconfig:"DB_USER"`
	Password      string `envconfig:"DB_PASSWORD"`
	ConnTimeoutMS int    `envconfig:"DB_CONN_TIMEOUT" default:"5000"`

	MaxOpenConns        int `envconfig:"DB_MAX_OPEN_CONNS" default:"10"`
	MaxIdleConns        int `envconfig:"DB_MAX_IDLE_CONNS" default:"5"`
	ConnMaxLifetimeMins int `envconfig:"DB_CONN_MAX_LIFETIME_MINS" default:"5"`

	GormDebugMode bool `envconfig:"DB_GORM_DEBUG_MODE" default:"false"`
}

type RMQConfig struct {
	UseTLS   bool   `envconfig:"AMQP_USE_TLS" default:"false"`
	Host     string `envconfig:"AMQP_HOST" default:"localhost"`
	Port     int    `envconfig:"AMQP_PORT" default:"5672"`
	User     string `envconfig:"AMQP_USER"`
	Password string `envconfig:"AMQP_PASSWORD"`
}

type SentryConfig struct {
	DSN           string  `envconfig:"SENTRY_DSN"`
	EnableTracing bool    `envconfig:"SENTRY_ENABLE_TRACING" default:"true"`
	SampleRate    float64 `envconfig:"SENTRY_SAMPLE_RATE" default:"0.2"`
}

type RedisConfig struct {
	Host     string `envconfig:"REDIS_HOST" default:"0.0.0.0"`
	Port     string `envconfig:"REDIS_PORT" default:"6379"`
	Password string `envconfig:"REDIS_PASSWORD" default:""`
}
