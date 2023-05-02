package main

import "fmt"

func main() {
	// DB AUTHENTICATE
	userName, _ := new(DBAuth).Authenticate("kaye", "qwer1234")
	fmt.Println(userName)

	// LDAP AUTHENTICATE
	userName, _ = new(LDAPAuth).Authenticate("dylee", "qwer1234")
	fmt.Println(userName)
}
