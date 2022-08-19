package model

import (
	"LdapAdmin/common/model"
	"LdapAdmin/db"
)

type Api struct {
	ID       int    `gorm:"type:int;primaryKey;autoIncrement;comment:接口id" json:"id"`
	Name     string `gorm:"" json:"name"`
	Path     string `gorm:"" json:"path"`
	Creator  string `json:"creator"`
	Operator string `json:"operator"`
	model.IntModel
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
