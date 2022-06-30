package pkg

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Redis    RedisConfig
	Postgres PostgresConfig
	Server   ServerConfig
	LogLevel string
}

type ServerConfig struct {
	Port      int
	PublicUri string
}

type RedisConfig struct {
	Host     string
	Port     int
	Password string
}

type PostgresConfig struct {
	Host     string
	Port     int
	Password string
}

func GetConfig() (Config, error) {
	var cfg Config

	err := viper.Unmarshal(&cfg)

	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func GetPostgresConnectionString() string {
	var cfg Config
	_ = viper.Unmarshal(&cfg)

	return fmt.Sprintf(`postgres://postgres:%s@%s:%d?sslmode=disable`, cfg.Postgres.Password, cfg.Postgres.Host, cfg.Postgres.Port)
}

func bindEnv(key string, defaultValue interface{}, envName string) {
	viper.SetDefault(key, defaultValue)
	_ = viper.BindEnv(key, envName)
}

func init() {

	bindEnv("server.port", 5003, "SERVER_PORT")
	bindEnv("server.debugPort", 5001, "DEBUG_PORT")

	bindEnv("loglevel", "debug", "LOG_LEVEL")

	bindEnv("redis.host", "localhost", "REDIS_HOST")
	bindEnv("redis.port", 6379, "REDIS_PORT")
	bindEnv("redis.password", "", "REDIS_PASSWORD")

	bindEnv("postgres.host", "localhost", "DB_HOST")
	bindEnv("postgres.port", 5432, "DB_PORT")
	bindEnv("postgres.password", "postgres", "DB_PASSWORD")
}
