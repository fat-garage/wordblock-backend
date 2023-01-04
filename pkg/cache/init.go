package cache

import (
	"context"
	"fmt"
	"github.com/fat-garage/wordblock-backend/pkg/conf"
	"log"

	"github.com/go-redis/redis/v8"
)

var Redis *redis.Client

func Init() {
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Cfg.Redis.Host, conf.Cfg.Redis.Port),
		Password: conf.Cfg.Redis.Password,
		DB:       conf.Cfg.Redis.Db,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal(err)
	}
	Redis = client
}
