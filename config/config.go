package config

import (
	"LdapAdmin/common/util"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

var Conf Config

type Config struct {
	System   *SystemConfig
	Database *DatabaseConfig
	Ldap     *LdapConfig
}

type SystemConfig struct {
	Mode string `yaml:"mode"`
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Dbname   string `yaml:"dbname"`
}

type LdapConfig struct {
	Host    string `yaml:"host"`
	Port    string `yaml:"port"`
	AdminDN string `yaml:"adminDN"`
	AdminPW string `yaml:"adminPW"`
	BaseDN  string `yaml:"baseDN"`
}

func InitConfig() {

	workDir, err := os.Getwd()
	if err != nil {
		util.PrintlnDangerous("read the directory failed!!! ")
		util.PrintlnDangerous("error: ", err.Error())
		return
	}

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/conf")

	if err := viper.ReadInConfig(); err != nil {
		util.PrintlnDangerous("read config failed !!!")
		util.PrintlnDangerous("error: ", err.Error())
		return
	}

	if err := viper.Unmarshal(&Conf); err != nil {
		util.PrintlnDangerous("unmarshal the config to the variable 'Conf' is failed !!!")
		util.PrintlnDangerous("error: ", err.Error())
		return
	}

	// when viper change, the conf also change
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.Unmarshal(&Conf); err != nil {
			util.PrintlnDangerous("unmarshal the config to the variable 'Conf' is failed !!!")
			util.PrintlnDangerous("error: ", err.Error())
			return
		}
	})
}
