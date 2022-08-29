package model

import (
	"LdapAdmin/common/model"
	"LdapAdmin/db"
)

type Api struct {
	ID          int    `gorm:"type:int;primaryKey;autoIncrement;not null;comment:the id of api" json:"id"`
	Active      int    `gorm:"type:int;not null;comment:the active status of api: 1 active 2 archived" json:"active"`
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
	Prefix      string `gorm:"type:varchar(50);not null;comment:the prefix of api" json:"prefix"`
	Description string `gorm:"type:varchar(510);not null;comment:the description of api_group" json:"description"`
	ApiList     []Api  `gorm:"foreignKey:api_group_id;associate_foreignKey:id" json:"apiList"`
	model.StringModel
}

var localApi Api
var localApiGroup ApiGroup

func (a *Api) TableName() string {
	return "ldap_admin_apis"
}
func (a *ApiGroup) TableName() string {
	return "ldap_admin_api_groups"
}

/* $ Api */

type AddApiReq struct {
	Name        string `json:"name"`         //The name of api
	Path        string `json:"path"`         //The path of api
	Description string `json:"description"`  //The description of api
	ApiGroupID  int    `json:"api_group_id"` //The api group id that is api belong to
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
	ID          int    `json:"id" binding:"required"` //The id that you want to modify
	Name        string `json:"name"`                  //The name of api to modifying
	Path        string `json:"path"`                  //The path of api to modifying
	Description string `json:"description"`           //The description of api to modifying
	ApiGroupID  int    `json:"api_group_id"`          //The group id of api to modifying
}

type GetApiListResp struct {
	ID           int    `json:"id"`
	Active       int    `json:"active"`
	Name         string `json:"name"`
	Path         string `json:"path"`
	Description  string `json:"description"`
	ApiGroupID   int    `json:"api_group_id"`
	ApiGroupName string `json:"api_group_name"`
	model.StringModel
}

/* $ ApiGroup */

type AddApiGroupReq struct {
	Name        string      `json:"name"`        //The name of group to api
	Prefix      string      `json:"prefix"`      //The prefix of path which belongs to group
	Description string      `json:"description"` //The description of group
	ApiList     []AddApiReq `json:"api_list"`    //The apis of group
}

type DeleteApiGroupReq struct {
	IDS []int `json:"ids" binding:"required"` //The id of group to delete
}

type GetApiGroupListReq struct {
}

type ModifyApiGroupReq struct {
}

/* $ Api Sql Operations */

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
		db.DB.Rollback()
		return err
	}
	return nil
}

func DeleteApiByGroupID(groupId int) error {
	if err := db.DB.Table(localApi.TableName()).
		Where("api_group_id = ?", groupId).
		Error; err != nil {
		db.DB.Rollback()
		return err
	}
	return nil
}

func GetApiList(req *GetApiListReq) ([]Api, int64, error) {
	return nil, 0, nil
}

func ModifyApi(id int, api Api) error {
	if err := db.DB.
		Where("id = ?", id).
		Updates(&api).
		Error; err != nil {
	}
	return nil
}

/* $ ApiGroup Sql Operations */

func AddApiGroup(group ApiGroup) (int, error) {
	if err := db.DB.Table(localApiGroup.TableName()).
		Create(&group).
		Error; err != nil {
		return 0, err
	}
	return group.ID, nil
}

func DeleteApiGroup(ids []int) error {
	if err := db.DB.Table(localApiGroup.TableName()).
		Where("id IN (?)", ids).
		Delete(&ApiGroup{}).
		Error; err != nil {
		return err
	}
	return nil
}

func GetApiGroupList(req *GetApiGroupListReq) {}

func ModifyApiGroup(id int, group ApiGroup) error {
	if err := db.DB.Table(localApiGroup.TableName()).
		Where("id = ?", id).
		Updates(&group).
		Error; err != nil {
		return err
	}
	return nil
}

/* $ Tools */

func convertApiToResp() {
}
