package main

import "fmt"

type AuthInterface interface {
	Authenticate(id, password string) string
}

type Auth struct {
	AuthInterface
}

func NewAbstractAuth(auth AuthInterface) *Auth {
	return &Auth{
		AuthInterface: auth,
	}
}

func (a *Auth) Auth(id, password string) (string, error) {
	userName := a.Authenticate(id, password)

	fmt.Println("step2. 인증 실패시 예외처리")

	fmt.Println("step3. 인증 성공시 인증정보 제공")

	return userName, nil
}
