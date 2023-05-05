package main

import (
	"fmt"
	"github.com/leeduyoung/GoDesignPatterns/behavioral/template_method/02_after/authenticate"
)

func main() {
	var (
		dbAuth   authenticate.AuthInterface = new(authenticate.DBAuth)
		ldapAuth authenticate.AuthInterface = new(authenticate.LDAPAuth)
	)

	// DB AUTHENTICATE
	userName, _ := authenticate.NewAbstractAuth(dbAuth).Auth("kaye", "qwer1234")
	fmt.Println(userName)

	// LDAP AUTHENTICATE
	userName, _ = authenticate.NewAbstractAuth(ldapAuth).Auth("dylee", "qwer1234")
	fmt.Println(userName)
}
