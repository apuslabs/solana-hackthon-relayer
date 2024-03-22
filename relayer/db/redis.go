package db

import (
	"apus-relayer/relayer/config"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func Init() {
	host := fmt.Sprintf("%s:%d", config.GetStr("redis.host"), config.GetInt("redis.port"))
	pwd := config.GetStr("redis.password")
	db := config.GetInt("redis.db")
	rdb = redis.NewClient(&redis.Options{
		Addr:     host,
		Password: pwd,
		DB:       db,
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		fmt.Print("连接Redis失败：", err)
		panic(err.Error())
	}
}
