package bootstrap

import (
	"by/user/container"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
)

func InitConfig(e *gin.Engine, app *container.Application) {
	var configPath = "./config"
	var configType = "yaml"
	var err error

	// app
	appViper := viper.New()
	appViper.AddConfigPath(configPath)
	appViper.SetConfigType(configType)
	appViper.SetConfigName("app")
	err = appViper.ReadInConfig()
	if err != nil {
		log.Printf("配置读取错误：%s", err.Error())
	} else {
		appViper.ConfigFileUsed()
		appViper.WatchConfig()
		app.Store("app", appViper)
	}

	// mysql
	mysqlViper := viper.New()
	mysqlViper.AddConfigPath(configPath)
	mysqlViper.SetConfigType(configType)
	mysqlViper.SetConfigName("mysql")
	err = mysqlViper.ReadInConfig()
	if err != nil {
		log.Printf("配置读取错误：%s", err.Error())
	} else {
		mysqlViper.ConfigFileUsed()
		mysqlViper.WatchConfig()
		app.Store("mysql", mysqlViper)
	}

	// redis
	redisViper := viper.New()
	redisViper.AddConfigPath(configPath)
	redisViper.SetConfigType(configType)
	redisViper.SetConfigName("redis")
	err = redisViper.ReadInConfig()
	if err != nil {
		log.Printf("配置读取错误：%s", err.Error())
	} else {
		redisViper.ConfigFileUsed()
		redisViper.WatchConfig()
		app.Store("redis", redisViper)
	}

	// elasticsearch
	esViper := viper.New()
	esViper.AddConfigPath(configPath)
	esViper.SetConfigType(configType)
	esViper.SetConfigName("elasticsearch")
	err = esViper.ReadInConfig()
	if err != nil {
		log.Printf("配置读取错误：%s", err.Error())
	} else {
		esViper.ConfigFileUsed()
		esViper.WatchConfig()
		app.Store("es", esViper)
	}
}
