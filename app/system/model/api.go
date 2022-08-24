package model

import (
	"LdapAdmin/common/model"
	"LdapAdmin/db"
)

type Api struct {
	ID          int    `gorm:"type:int;primaryKey;autoIncrement;not null;comment:the id of api" json:"id"`
	ApiGroupID  int    `gorm:"type:int;not null;comment:the id of group which the api belong" json:"api_group_id"`
	Name        string `gorm:"type:varchar(128);not null;comment:the name of api" json:"name"`
	Path        string `gorm:"type:varchar(100);not null;comment:the path of api" json:"path"`
	Description string `gorm:"type:varchar(510);not null;comment:the description of api_group" json:"description"`
	model.StringModel
}

type ApiGroup struct {
	ID          int    `gorm:"type:int;primaryKey;autoIncrement;not null;comment:the id of api_group" json:"id"`
	Name        string `gorm:"type:varchar(20);not null;comment:the name of api_group" json:"name"`
	Description string `gorm:"type:varchar(510);not null;comment:the description of api_group" json:"description"`
	ApiList     []Api  `gorm:"foreignKey:api_group_id;associate_foreignKey:id" json:"apiList"`
	model.StringModel
}

var localApi Api

func (a *Api) TableName() string {
	return "ldap_admin_apis"
}

type AddApiReq struct {
	Name    string `json:"name"`    //The name of api
	Path    string `json:"path"`    //The path of api
	Creator string `json:"creator"` //Who is created the api
}

type DeleteApiReq struct {
	IDS []int `json:"ids"` //The array of id which belong the api, that you want to delete
}

type GetApiListReq struct {
}

type ModifyApiReq struct {
}

func AddApi(api Api) (int, error) {
	if err := db.DB.
		Create(&api).
		Error; err != nil {
		return 0, err
	}
	return api.ID, nil
}

func DeleteApi(ids []int) error {
	if err := db.DB.
		Delete(&Api{}, "id IN (?)", ids).
		Error; err != nil {
		return err
	}
	return nil
}

func GetApiList() {

}

func ModifyApi() error {
	return nil
}
