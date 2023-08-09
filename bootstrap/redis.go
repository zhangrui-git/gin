package bootstrap

import (
	"by/user/container"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func InitRedis(e *gin.Engine, app *container.Application) {
	redisConf := app.Config("redis")
	defaultConf := redisConf.Sub("redis.connections.default")
	opt := redis.Options{
		Addr:     defaultConf.GetString("host"),
		Password: defaultConf.GetString("password"),
		DB:       0,
		PoolSize: 10,
	}
	rdb := redis.NewClient(&opt)
	app.Store("redis", rdb)
}
