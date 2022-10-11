package main

import (
	base "LdapAdmin/app/system/model"
	"LdapAdmin/config"
	"LdapAdmin/db"
	"LdapAdmin/router"
)

func generateTable() {
	db.DB.AutoMigrate(
		&base.ApiGroup{},
		&base.Api{},
		//&base.Menu{},
		//&model.Token{},
	)
}

func main() {
	config.InitConfig()
	//ldap.InitLdap()
	db.InitDatabase()

	// $ generate table
	generateTable()

	router.InitRouter()
	// $ generate rsa keys
	//if err := util.GenerateRsaKey(1024, config.Conf.System.RsaKeyFolder); err != nil {
	//	fmt.Println(err.Error())
	//}
}
