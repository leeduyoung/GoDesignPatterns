package main

import "fmt"

func main() {
	var (
		dbAuth   AuthInterface = new(DBAuth)
		ldapAuth AuthInterface = new(LDAPAuth)
	)

	// DB AUTHENTICATE
	userName, _ := NewAbstractAuth(dbAuth).Auth("kaye", "qwer1234")
	fmt.Println(userName)

	// LDAP AUTHENTICATE
	userName, _ = NewAbstractAuth(ldapAuth).Auth("dylee", "qwer1234")
	fmt.Println(userName)
}
