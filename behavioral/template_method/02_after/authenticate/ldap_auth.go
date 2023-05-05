package authenticate

import "fmt"

type LDAPAuth struct {
}

// authenticate 패키지 내에서만 호출되도록 제한
func (ldapa *LDAPAuth) authenticate(id, password string) string {
	// step1. 사용자 정보로 인증 확인
	return ldapLogin(id, password)
}

func ldapLogin(id, password string) string {
	// DO SOMETHING...
	fmt.Println("step1. LDAP 사용자 정보로 인증 확인")
	return id
}
