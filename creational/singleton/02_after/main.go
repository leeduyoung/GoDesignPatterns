package main

import (
	"fmt"
	"github.com/leeduyoung/GoDesignPatterns/creational/singleton/02_after/singleton"
)

func main() {
	instance := singleton.Instance()
	instance.Data = "Change Singleton Data"
	anotherInstance := singleton.Instance()

	fmt.Println(instance.Data)
	fmt.Println(anotherInstance.Data)
}
