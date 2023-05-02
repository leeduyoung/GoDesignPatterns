package main

import "fmt"

type LDAPAuth struct {
}

func (ldapa *LDAPAuth) Authenticate(id, password string) string {
	// step1. 사용자 정보로 인증 확인
	return ldapLogin(id, password)
}

func ldapLogin(id, password string) string {
	// DO SOMETHING...
	fmt.Println("step1. LDAP 사용자 정보로 인증 확인")
	return id
}
