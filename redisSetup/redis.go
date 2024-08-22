package repository

import (
	"fmt"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var (
	defaultRedis *redis.Client
)

type Redis struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

func InitRedis() {
	redisCfg := Redis{
		Host: viper.GetString("redis.host"),
		Port: viper.GetInt("redis.port"),
		Password: viper.GetString("redis.password"),
		DB:       viper.GetInt("redis.db"),
	}
	addr := fmt.Sprintf("%s:%d", redisCfg.Host, redisCfg.Port)
	defaultRedis = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: redisCfg.Password, // no password set
		DB:       redisCfg.DB,       // use default DB
	})
}

func GetRedisClient() *redis.Client {
	return defaultRedis
}
