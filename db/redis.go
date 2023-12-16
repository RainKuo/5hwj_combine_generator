package db

import (
	"fmt"
	"github.com/go-redis/redis"
)

type RedisDriver struct {
	cli *redis.Client
}

func (rd *RedisDriver) ConnectRedis() {
	rd.cli = redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379", Password: "", DB: 13})
	_, err := rd.cli.Ping().Result()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (rd *RedisDriver) Set(key string, data interface{}) {
	rd.cli.Set(key, data, 0)
}
