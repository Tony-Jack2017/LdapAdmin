package main

import (
	"LdapAdmin/config"
	"LdapAdmin/db"
	"LdapAdmin/ldap"
)

func main() {
	config.InitConfig()
	ldap.InitLdap()
	db.InitDatabase()
}
