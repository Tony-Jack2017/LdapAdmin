package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
)

var Conf Config

type Config struct {
	System   *SystemConfig   `json:"system"`
	Redis    *RedisConfig    `json:"redis"`
	Database *DatabaseConfig `json:"database"`
	Ldap     *LdapConfig     `json:"ldap"`
}

type SystemConfig struct {
	Mode         string `yaml:"mode"`
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	TokenExpired int    `yaml:"tokenExpired"`
	TokenSecret  string `yaml:"tokenSecret"`
	RsaKeyFolder string `yaml:"rsaKeyFolder"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Password string `yaml:"password"`
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
	Port    int    `yaml:"port"`
	AdminDN string `yaml:"adminDN"`
	AdminPW string `yaml:"adminPW"`
	BaseDN  string `yaml:"baseDN"`
}

func InitConfig() {

	workDir, err := os.Getwd()
	if err != nil {
		fmt.Println("read the directory failed!!! ")
		fmt.Println("error: ", err.Error())
		return
	}

	viper.SetConfigName("config_work")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(workDir + "/conf")

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("read config failed !!!")
		fmt.Println("error: ", err.Error())
		return
	}

	if err := viper.Unmarshal(&Conf); err != nil {
		fmt.Println("unmarshal the config to the variable 'Conf' is failed !!!")
		fmt.Println("error: ", err.Error())
		return
	}

	// when viper change, the conf also change
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.Unmarshal(&Conf); err != nil {
			fmt.Println("unmarshal the config to the variable 'Conf' is failed !!!")
			fmt.Println("error: ", err.Error())
			return
		}
	})
}
