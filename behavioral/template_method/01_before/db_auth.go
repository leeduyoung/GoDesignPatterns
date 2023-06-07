package main

import "fmt"

type DBAuth struct {
}

func (dba *DBAuth) Authenticate(id, password string) (string, error) {
	// step1. 사용자 정보로 인증 확인
	userName := dbLogin(id, password)

	dba.handleFailedAuthentication()
	dba.handleSuccessAuthentication()

	return userName, nil
}

func (dba *DBAuth) handleSuccessAuthentication() {
	fmt.Println("step3. 인증 성공시 인증정보 제공")
}

func (dba *DBAuth) handleFailedAuthentication() {
	fmt.Println("step2. 인증 실패시 예외처리")
}

func dbLogin(id, password string) string {
	// DO SOMETHING...
	fmt.Println("step1. DB 사용자 정보로 인증 확인")
	return id
}
