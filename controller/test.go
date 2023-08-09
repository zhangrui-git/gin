package controller

import (
	"by/user/container"
	"by/user/redisKey"
	"by/user/response"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"log"
	"net/http"
)

func Config(ctx *gin.Context) {
	mysqlConf := container.Config("mysql")
	log.Println(mysqlConf.Get("mysql.connections"))

	esConf := container.Config("es")
	log.Println(esConf.Get("elasticsearch.connections"))
}

func Redis(ctx *gin.Context) {
	var err error
	rdbCtx := context.Background()
	rdb := container.Redis("redis")
	var result, redisErr = rdb.Get(rdbCtx, redisKey.UserModel(31257)).Result()
	log.Println(result)

	switch {
	case redisErr == redis.Nil || result == "":
		ctx.JSON(http.StatusOK, response.Success(struct{}{}))
		break
	case redisErr != nil:
		log.Println(err.Error())
		ctx.JSON(http.StatusOK, response.Failed(400000))
		break
	default:
		//var mu bean.UserBean
		//err = json.Unmarshal([]byte(result), &mu)
		//if err != nil {
		//	log.Println(err.Error())
		//	ctx.JSON(http.StatusOK, response.Failed(400000))
		//} else {
		//	ctx.JSON(http.StatusOK, response.Success(mu))
		//}
		break
	}

	// 返回json串
	//ctx.Header("Content-Type", "application/json; charset=utf-8")
	//ctx.String(http.StatusOK, "%s", result)

	// 返回struct
	//var u bean.UserBean
	//err = json.Unmarshal([]byte(result), &u)
	//if err != nil {
	//	log.Println(err.Error())
	//}
	//ctx.JSON(http.StatusOK, u)
}
