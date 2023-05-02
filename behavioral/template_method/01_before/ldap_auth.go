package main

import "fmt"

type LDAPAuth struct {
}

func (ldapa *LDAPAuth) Authenticate(id, password string) (string, error) {
	// step1. 사용자 정보로 인증 확인
	userName := ldapLogin(id, password)

	fmt.Println("step2. 인증 실패시 예외처리")

	fmt.Println("step3. 인증 성공시 인증정보 제공")

	return userName, nil
}

func ldapLogin(id, password string) string {
	// DO SOMETHING...
	fmt.Println("step1. LDAP 사용자 정보로 인증 확인")
	return id
}
