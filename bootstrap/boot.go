package bootstrap

import (
	"by/user/container"
	"github.com/gin-gonic/gin"
	"log"
)

func Boot(e *gin.Engine) {
	log.Default().Println("初始化容器")
	app := container.GetApp()
	app.Store("engine", e)

	log.Default().Println("初始化配置")
	InitConfig(e, app)

	log.Default().Println("初始化数据库")
	InitDatabase(e, app)

	log.Default().Println("初始化Redis")
	InitRedis(e, app)

	log.Default().Println("初始化路由")
	InitRouter(e, app)

	log.Default().Printf("服务启动 %s", app.StartTime.Format("2006-01-02 15:04:05"))
}
