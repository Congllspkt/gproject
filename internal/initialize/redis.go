package initialize

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gproject/internal/initialize/global"
)

var ctx = context.Background()

func InitRedis() {
	r := global.Config.Redis

	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", r.Host, r.Port),
		Password: r.Password,
		DB:       r.Database,
		PoolSize: 10,
	})

	_, err := rdb.Ping(ctx).Result()

	if err != nil {
		global.Logger.Error("Redis Init Error")
		panic("Init Redis fail")
	}
	global.Logger.Info("Redis Init Success")
	global.Rdb = rdb
}
