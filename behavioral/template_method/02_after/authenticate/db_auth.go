package authenticate

import "fmt"

type DBAuth struct {
}

// authenticate 패키지 내에서만 호출되도록 제한
func (dba *DBAuth) authenticate(id, password string) string {
	// step1. 사용자 정보로 인증 확인
	return dbLogin(id, password)
}

func dbLogin(id, password string) string {
	// DO SOMETHING...
	fmt.Println("step1. DB 사용자 정보로 인증 확인")
	return id
}
