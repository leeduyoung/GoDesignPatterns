package main

import "fmt"

type DBAuth struct {
}

func (dba *DBAuth) Authenticate(id, password string) string {
	// step1. 사용자 정보로 인증 확인
	return dbLogin(id, password)
}

func dbLogin(id, password string) string {
	// DO SOMETHING...
	fmt.Println("step1. DB 사용자 정보로 인증 확인")
	return id
}
