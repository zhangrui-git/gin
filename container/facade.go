package container

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func (c *Application) Config(name string) *viper.Viper {
	v, ok := c.Load(name)
	if ok {
		return v.(*viper.Viper)
	}
	return nil
}

func Config(name string) *viper.Viper {
	return app.Config(name)
}

func (c *Application) Database() *gorm.DB {
	db, ok := c.Load("db")
	if ok {
		return db.(*gorm.DB)
	}
	return nil
}

func Database() *gorm.DB {
	return app.Database()
}

func (c *Application) Redis(name string) *redis.Client {
	rdb, ok := c.Load("redis")
	if ok {
		return rdb.(*redis.Client)
	}
	return nil
}

func Redis(name string) *redis.Client {
	return app.Redis(name)
}
