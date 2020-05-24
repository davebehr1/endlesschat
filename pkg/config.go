package pkg

import (
	"github.com/spf13/viper"
)

type Config struct {
	Redis  RedisConfig
	Server ServerConfig
}

type ServerConfig struct {
	Port      int
	PublicUri string
}

type RedisConfig struct {
	Host string
	Port int
}

func GetConfig() (Config, error) {
	var cfg Config

	err := viper.Unmarshal(&cfg)

	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func bindEnv(key string, defaultValue interface{}, envName string) {
	viper.SetDefault(key, defaultValue)
	_ = viper.BindEnv(key, envName)
}

func init() {

	bindEnv("server.port", 5000, "SERVER_PORT")
	bindEnv("server.debugPort", 5001, "DEBUG_PORT")

	bindEnv("redis.host", "localhost", "REDIS_HOST")
	bindEnv("redis.port", 6379, "REDIS_PORT")
}
