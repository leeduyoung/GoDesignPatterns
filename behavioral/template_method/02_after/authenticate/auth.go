package authenticate

import "fmt"

type AuthInterface interface {
	authenticate(id, password string) string
}

type Auth struct {
	AuthInterface
}

func NewAbstractAuth(auth AuthInterface) *Auth {
	return &Auth{
		AuthInterface: auth,
	}
}

// Auth 퍼블릭 메서드로 공개된다
func (a *Auth) Auth(id, password string) (string, error) {
	userName := a.authenticate(id, password)

	fmt.Println("step2. 인증 실패시 예외처리")

	fmt.Println("step3. 인증 성공시 인증정보 제공")

	return userName, nil
}
