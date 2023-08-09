package container

import "sync"

type Service map[string]*sync.Pool

var svc Service

func init() {
	svc = make(Service)
}

func (s Service) Register(name string, f func() any) {
	s[name] = &sync.Pool{New: f}
}

func (s Service) Get(name string) any {
	return s[name].Get()
}

func (s Service) Put(name string, o any) {
	s[name].Put(o)
}

func GetPool() Service {
	return svc
}
