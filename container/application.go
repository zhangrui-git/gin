package container

import (
	"sync"
	"time"
)

type Application struct {
	m         sync.Map
	Server    string
	Version   string
	StartTime time.Time
}

var app *Application

func init() {
	app = &Application{
		m:         sync.Map{},
		Server:    "by-api",
		Version:   "0.0.1",
		StartTime: time.Now(),
	}
}

func (c *Application) Store(key, value any) {
	c.m.Store(key, value)
}

func (c *Application) Load(key any) (any, bool) {
	return c.m.Load(key)
}

func (c *Application) Delete(key any) {
	c.m.Delete(key)
}

func GetApp() *Application {
	return app
}
