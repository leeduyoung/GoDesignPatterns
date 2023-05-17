package main

import (
	"fmt"
)

func main() {
	instance := Instance()
	instance.Data = "Change Singleton Data"
	anotherInstance := Instance()

	fmt.Println(instance.Data)
	fmt.Println(anotherInstance.Data)
}

type NotSingleton struct {
	Data string
}

var instance *NotSingleton

func Instance() *NotSingleton {
	return &NotSingleton{
		Data: "Hello world!",
	}
}
