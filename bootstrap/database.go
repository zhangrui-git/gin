package bootstrap

import (
	"by/user/container"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func InitDatabase(e *gin.Engine, app *container.Application) {
	mysqlConf := app.Config("mysql")
	defaultConf := mysqlConf.Sub("mysql.connections.default")
	str := "%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local"
	dsn := fmt.Sprintf(
		str,
		defaultConf.GetString("username"),
		defaultConf.GetString("password"),
		defaultConf.GetString("host"),
		defaultConf.GetInt("port"),
		defaultConf.GetString("database"),
		defaultConf.GetString("charset"),
	)
	db, err1 := gorm.Open(mysql.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	if err1 != nil {
		log.Panicf("Mysql连接失败:%s", err1.Error())
	}
	sqlDB, err2 := db.DB()
	if err2 != nil {
		log.Panicf("Mysql配置错误:%s", err2.Error())
	}
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(20)
	app.Store("db", db)
}
