package config

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

const (
	yamlPath = "configs"
	yamlFile = "config"
	envFile  = ".env"
)

type Config struct {
	Postgres PostgresConfig
	// Redis         RedisConfig
	// Cache         CacheConfig
	HTTPServer    HTTPServerConfig
	HandlerConfig HandlerConfig
}

type PostgresConfig struct {
	Username string
	Password string
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	DBName   string `mapstructure:"db_name"`
	SSLMode  string `mapstructure:"ssl_mode"`
}

// type RedisConfig struct {
// 	Password string
// 	Host     string `mapstructure:"host"`
// 	Port     string `mapstructure:"port"`
// 	DBName   int    `mapstructure:"db_name"`
// }

// type CacheConfig struct {
// 	Expiration time.Duration `mapstructure:"expiration"`
// }

type HTTPServerConfig struct {
	Port           string        `mapstructure:"port"`
	ReadTimeout    time.Duration `mapstructure:"read_timeout"`
	WriteTimeout   time.Duration `mapstructure:"write_timeout"`
	MaxHeaderBytes int           `mapstructure:"max_header_bytes"`
}

type HandlerConfig struct {
	Books BooksHandlerConfig
}

type BooksHandlerConfig struct {
	Limit  int `mapstructure:"limit"`
	Offset int `mapstructure:"offset"`
}

func InitConfig() (*Config, error) {
	if err := read(); err != nil {
		return nil, err
	}

	var cfg Config
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func unmarshal(cfg *Config) error {
	if err := viper.UnmarshalKey("postgres", &cfg.Postgres); err != nil {
		return fmt.Errorf("unmarshal postgres config: %v", err)
	}

	// if err := viper.UnmarshalKey("redis", &cfg.Redis); err != nil {
	// 	return fmt.Errorf("unmarshal redis config: %v", err)
	// }

	// if err := viper.UnmarshalKey("cache", &cfg.Cache); err != nil {
	// 	return fmt.Errorf("unmarshal cache config: %v", err)
	// }

	if err := viper.UnmarshalKey("http_server", &cfg.HTTPServer); err != nil {
		return fmt.Errorf("unmarshal http server config: %v", err)
	}

	if err := viper.UnmarshalKey("books_handler", &cfg.HandlerConfig.Books); err != nil {
		return fmt.Errorf("unmarshal books_handler config: %v", err)
	}

	cfg.Postgres.Username = os.Getenv("POSTGRES_USER")
	cfg.Postgres.Password = os.Getenv("POSTGRES_PASSWORD")

	return nil
}

func read() error {
	if err := godotenv.Load(envFile); err != nil {
		return fmt.Errorf("read env file: %s: %s", envFile, err)
	}

	viper.AddConfigPath(yamlPath)
	viper.SetConfigName(yamlFile)
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("read yaml file: %v: %v", envFile, err)
	}

	return nil
}
