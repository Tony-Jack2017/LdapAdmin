package model

import (
	"LdapAdmin/common/model"
	"LdapAdmin/db"
)

type Api struct {
	ID          int    `gorm:"type:int;primaryKey;autoIncrement;not null;comment:the id of api" json:"id"`
	Active      int    `gorm:"type:int;not null;comment:the active status of api" json:"active"`
	Name        string `gorm:"type:varchar(128);not null;comment:the name of api" json:"name"`
	Path        string `gorm:"type:varchar(100);not null;comment:the path of api" json:"path"`
	Description string `gorm:"type:varchar(510);not null;comment:the description of api_group" json:"description"`
	ApiGroupID  int    `gorm:"type:int;not null;comment:the id of group which the api belong" json:"api_group_id"`
	model.StringModel
}

type ApiGroup struct {
	ID          int    `gorm:"type:int;primaryKey;autoIncrement;not null;comment:the id of api_group" json:"id"`
	Active      int    `gorm:"type:int;not null;comment:the active status of api_group" json:"active"`
	Name        string `gorm:"type:varchar(20);not null;comment:the name of api_group" json:"name"`
	Path        string `gorm:"type:varchar(50);not null;comment:the path of api_group" json:"path"`
	Description string `gorm:"type:varchar(510);not null;comment:the description of api_group" json:"description"`
	ApiList     []Api  `gorm:"foreignKey:api_group_id;associate_foreignKey:id" json:"apiList"`
	model.StringModel
}

var localApi Api

func (a *Api) TableName() string {
	return "ldap_admin_apis"
}

type AddApiReq struct {
	Name        string `json:"name"`        //The name of api
	Path        string `json:"path"`        //The path of api
	Description string `json:"description"` //The description of api
}

type DeleteApiReq struct {
	IDS []int `json:"ids"` //The array of id which belong the api, that you want to delete
}

type GetApiListReq struct {
	Active      int    `form:"active"`      //Search apis by the active status of api
	Name        string `form:"name"`        //Search apis by name
	Path        string `form:"path"`        //Search apis by the path of api
	Description string `form:"description"` //Search apis by description
	ApiGroupID  int    `form:"apiGroupId"`  //Search apis by the group id of api
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

func GetApiList(req *GetApiListReq) ([]Api, int64, error) {
	return nil, 0, nil
}

func ModifyApi() error {
	return nil
}
