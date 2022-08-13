package db

import (
	"LdapAdmin/common/util"
	"LdapAdmin/config"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s  dbname=%s password=%s",
		config.Conf.Database.Host,
		config.Conf.Database.Port,
		config.Conf.Database.Username,
		config.Conf.Database.Dbname,
		config.Conf.Database.Password,
	)
	dialector := postgres.Open(dsn)
	db, err := gorm.Open(dialector, &gorm.Config{})

	if err != nil {
		util.PrintlnDangerous("connect the database failed !!!")
		util.PrintlnDangerous(" error: ", err.Error())
	}
	DB = db
	showDsn := fmt.Sprintf("host=%s port=%s user=%s @dbname=%s",
		config.Conf.Database.Host,
		config.Conf.Database.Port,
		config.Conf.Database.Username,
		config.Conf.Database.Dbname,
	)
	util.PrintlnSuccess("init the database success >>>>>")
	util.PrintlnSuccess("connection: ", showDsn)
}
