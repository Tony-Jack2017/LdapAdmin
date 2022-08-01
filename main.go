package main

import (
	"go-ldap/config"
	"go-ldap/ldap"
)

func main() {
	config.InitConfig()
	ldap.InitLdap()
}
