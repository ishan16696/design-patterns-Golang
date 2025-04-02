package singleton

import "sync"

type Singleton struct {
	Data string
	Info int
}

var (
	instance *Singleton
	once     sync.Once
)

func GetInstance(info int) *Singleton {
	once.Do(func() {
		instance = &Singleton{
			Data: "I am thread-safe!",
			Info: info,
		}
	})
	return instance
}
