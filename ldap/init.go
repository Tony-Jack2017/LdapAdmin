package ldap

import (
	"crypto/tls"
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"go-ldap/common/util"
	"go-ldap/config"
)

var LDAP *ldap.Conn

func InitLdap() {
	conn, err := ldap.DialTLS("tls", "192.168.2.10:389",&tls.Config{
		InsecureSkipVerify: true,
	})
	if err != nil {
		util.PrintlnDangerous("connect the ldap server failed !!!")
		util.PrintlnDangerous("error: ", err.Error())
	}

	// verify admin bindDN and Password to build ldap connect
	err = conn.Bind(config.Conf.Ldap.AdminDN, config.Conf.Ldap.AdminPW)

	if err != nil {
		util.PrintlnDangerous("verify administrator account failed !!!")
		util.PrintlnDangerous("error: ", err.Error())
	}

	LDAP = conn

	showDsn := fmt.Sprintf("%s:******@tcp(%s)",
		config.Conf.Ldap.AdminDN,
		config.Conf.Ldap.Host + config.Conf.Ldap.Port,
	)

	util.PrintlnSuccess("connect ldap server success >>>>")
	util.PrintlnSuccess(showDsn)
}