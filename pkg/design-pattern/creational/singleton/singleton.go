package singleton

import "sync"

type Singleton interface {
	doSomething()
}

type singleton struct{}

func (s *singleton) doSomething() {}

var (
	instance *singleton
	once     sync.Once
)

func GetInstance() Singleton {
	once.Do(func() {
		instance = &singleton{}
	})

	return instance
}
