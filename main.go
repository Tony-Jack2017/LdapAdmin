package main

import (
	"LdapAdmin/config"
)

func main() {
	config.InitConfig()
	//ldap.InitLdap()
	//db.InitDatabase()

	//generate rsa keys
	//if err := util.GenerateRsaKey(1024, config.Conf.System.RsaKeyFolder); err != nil {
	//	fmt.Println(err.Error())
	//}
}
