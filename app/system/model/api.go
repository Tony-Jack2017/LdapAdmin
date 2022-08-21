package model

import "LdapAdmin/common/model"

type Api struct {
	ID   int    `gorm:"type:int;primaryKey;autoIncrement;not null;comment:the id of api" json:"id"`
	Name string `gorm:"type:varchar(128);not null;comment:the name of api" json:"name"`
	Path string `gorm:"type:varchar(100);not null;comment:the path of api" json:"path"`
	model.StringModel
}

var localApi Api

func (a *Api) TableName() string {
	return "ldap_admin_apis"
}

func AddApi() {

}

func DeleteApi() {

}

func GetApiList() {

}

func ModifyApi() {

}
