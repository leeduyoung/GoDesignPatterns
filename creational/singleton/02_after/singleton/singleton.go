package singleton

import "sync"

type Singleton struct {
	Data string
}

var instance *Singleton
var once sync.Once

func Instance() *Singleton {
	once.Do(func() {
		instance = &Singleton{
			Data: "Hello world!",
		}
	})
	return instance
}
