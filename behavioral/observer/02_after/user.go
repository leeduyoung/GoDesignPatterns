package main

import "fmt"

type User struct {
	name string
}

func NewUser(name string) *User {
	return &User{
		name: name,
	}
}

func (u *User) Name() string {
	return u.name
}

func (u *User) HandleMessage(message string) {
	prefix := "[" + u.name + " 화면]"
	fmt.Println(prefix + message)
}
