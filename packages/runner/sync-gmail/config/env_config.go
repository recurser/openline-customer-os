package config

import (
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/logger"
	"github.com/openline-ai/openline-customer-os/packages/server/customer-os-common-module/tracing"
)

type Config struct {
	Neo4jDb struct {
		Target                string `env:"NEO4J_TARGET,required"`
		User                  string `env:"NEO4J_AUTH_USER,required,unset"`
		Pwd                   string `env:"NEO4J_AUTH_PWD,required,unset"`
		Realm                 string `env:"NEO4J_AUTH_REALM"`
		MaxConnectionPoolSize int    `env:"NEO4J_MAX_CONN_POOL_SIZE" envDefault:"100"`
		LogLevel              string `env:"NEO4J_LOG_LEVEL" envDefault:"WARNING"`
	}
	PostgresDb struct {
		Host            string `env:"POSTGRES_HOST,required"`
		Port            string `env:"POSTGRES_PORT,required"`
		User            string `env:"POSTGRES_USER,required,unset"`
		Db              string `env:"POSTGRES_DB,required"`
		Password        string `env:"POSTGRES_PASSWORD,required,unset"`
		MaxConn         int    `env:"POSTGRES_DB_MAX_CONN"`
		MaxIdleConn     int    `env:"POSTGRES_DB_MAX_IDLE_CONN"`
		ConnMaxLifetime int    `env:"POSTGRES_DB_CONN_MAX_LIFETIME"`
		LogLevel        string `env:"POSTGRES_LOG_LEVEL" envDefault:"WARN"`
	}

	Anthropic struct {
		ApiPath          string `env:"ANTHROPIC_API_PATH,required" envDefault:"WARN"`
		ApiKey           string `env:"ANTHROPIC_API_KEY,required" envDefault:"WARN"`
		SummaryPrompt    string `env:"ANTHROPIC_SUMMARY_PROMPT,required" envDefault:"WARN"`
		ActionItemsPromp string `env:"ANTHROPIC_ACTION_ITEMS_PROMPT,required" envDefault:"WARN"`
	}

	OpenAi struct {
		ApiPath string `env:"OPENAI_API_PATH,required" envDefault:"WARN"`
		ApiKey  string `env:"OPENAI_API_KEY,required" envDefault:"WARN"`
	}

	ValidationApi struct {
		Path string `env:"VALIDATION_API_PATH" validate:"required"`
		Key  string `env:"VALIDATION_API_KEY" validate:"required"`
	}

	SyncData struct {
		CronSync string `env:"CRON_SYNC" envDefault:"0 */1 * * * *"`
	}

	Jaeger tracing.JaegerConfig
	Logger logger.Config
}
