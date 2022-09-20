package model

import (
	"LdapAdmin/common/model"
	"LdapAdmin/db"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strings"
)

type Api struct {
	ID          int    `gorm:"type:int;primaryKey;autoIncrement;not null;comment:the id of api" json:"id"`
	Active      int    `gorm:"type:int;not null;comment:the active status of api: 1 active 2 archived" json:"active"`
	Status      int    `gorm:"type:int;not null;comment:the status of api: 1 enable 2 disable" json:"status"`
	Method      string `gorm:"type:varchar(20);not null;comment:the type of api" json:"method"`
	Name        string `gorm:"type:varchar(128);not null;comment:the name of api" json:"name"`
	Path        string `gorm:"type:varchar(100);not null;comment:the path of api" json:"path"`
	Description string `gorm:"type:varchar(510);not null;comment:the description of api_group" json:"description"`
	ApiGroupID  int    `gorm:"type:int;not null;comment:the id of group which the api belong" json:"api_group_id"`
	model.StringModel
}

type ApiGroup struct {
	ID          int    `gorm:"type:int;primaryKey;autoIncrement;not null;comment:the id of api_group" json:"id"`
	Active      int    `gorm:"type:int;not null;comment:the active status of api_group: 1 active 2 archived" json:"active"`
	Status      int    `gorm:"type:int;not null;comment:the status of api: 1 enable 2 disable" json:"status"`
	Name        string `gorm:"type:varchar(20);not null;comment:the name of api_group" json:"name"`
	Prefix      string `gorm:"type:varchar(50);unique;not null;comment:the prefix of api" json:"prefix"`
	Description string `gorm:"type:varchar(510);;comment:the description of api_group" json:"description"`
	ApiList     []Api  `gorm:"foreignKey:api_group_id;associate_foreignKey:id" json:"apiList,omitempty"`
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
	Status      int    `json:"status" binding:"required,oneof=1 2"`          //The status of api: 1 enable 2 disable
	Method      string `json:"method" binding:"required,oneof=GET POST PUT"` //The type of api: GET POST PUT
	Name        string `json:"name" binding:"required"`                      //The name of api
	Path        string `json:"path" binding:"required"`                      //The path of api
	Description string `json:"description"`                                  //The description of api
	ApiGroupID  int    `json:"api_group_id" binding:"required"`              //The api group id that is api belong to
}

type DeleteApiReq struct {
	IDS []int `json:"ids" binding:"required"` //The array of id which belong the api, that you want to delete
}

type GetApiListReq struct {
	Active      int    `form:"active" binding:"required,oneof=1 2"` //Search apis by the active status of api: 1 active 2 archived
	Status      int    `form:"status" binding:"oneof=0 1 2"`        //Search apis by the status of api: 1 enable 2 disable
	Method      string `form:"method"`                              //The type of api: GET POST PUT
	Name        string `form:"name"`                                //Search apis by name
	Path        string `form:"path"`                                //Search apis by the path of api
	Description string `form:"description"`                         //Search apis by description
	ApiGroupID  int    `form:"apiGroupId"`                          //Search apis by the group id of api
	model.PaginationOption
}

type ModifyApiReq struct {
	ID          int    `json:"id" binding:"required"`             //The id that you want to modify
	Type        int    `json:"type" binding:"required,oneof=1 2"` //The type of modify: 1 normal, 2 unarchived
	Status      int    `json:"status"`                            //The status of api: 1 enable 2 disable
	Method      string `json:"method"`                            //The type of api
	Name        string `json:"name"`                              //The name of api to modifying
	Path        string `json:"path"`                              //The path of api to modifying
	Description string `json:"description"`                       //The description of api to modifying
	ApiGroupID  int    `json:"api_group_id"`                      //The group id of api to modifying
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
	Name        string      `json:"name" binding:"required"`   //The name of group to api
	Prefix      string      `json:"prefix" binding:"required"` //The prefix of path which belongs to group
	Description string      `json:"description"`               //The description of group
	ApiList     []AddApiReq `json:"api_list"`                  //The apis of group
}

type DeleteApiGroupReq struct {
	IDS []int `json:"ids" binding:"required"` //The id of group to delete
}

type GetApiGroupListReq struct {
	Active      int    `form:"active" binding:"required,oneof=1 2"` //Search group by the active status: 1 active, 2 archived
	Type        int    `form:"type" binding:"required,oneof=1 2"`   //Search type: 1 normal, 2 cascade
	Name        string `form:"name"`                                //Search group by the name
	Prefix      string `form:"prefix"`                              //Search group by the prefix
	Description string `form:"description"`                         //Search group by the description
	model.PaginationOption
}

type ModifyApiGroupReq struct {
	ID   int `json:"id" binding:"required, oneof=1 2"`  //The id of the group will be modified
	Type int `json:"type" binding:"required,oneof=1 2"` //The type of modify: 1 normal, 2 unarchived

}

/* $ Api Sql Operations */

func AddApi(api Api) (int, error) {
	if err := db.DB.
		Create(&api).
		Error; err != nil {
		db.DB.Rollback()
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
	var apis []Api
	var conn *gorm.DB
	switch req.Active {
	case 1:
		conn = db.DB.Model(&Api{}).Order("id")
	case 2:
		conn = db.DB.Model(&Api{}).Order("id").Unscoped()
	default:
		return nil, 0, errors.New("the request active type of get menu list is only supported 1 or 2")
	}
	name := strings.TrimSpace(req.Name)
	if name != "" {
		conn = conn.Where("name LIKE ?", fmt.Sprintf("%%%s%%", name))
	}
	path := strings.TrimSpace(req.Path)
	if path != "" {
		conn = conn.Where("path LIKE ?", fmt.Sprintf("%%%s%%", path))
	}
	description := strings.TrimSpace(req.Description)
	if description != "" {
		conn = conn.Where("description LIKE ?", fmt.Sprintf("%%%s%%", description))
	}
	if req.Status != 0 {
		conn = conn.Where("status = ?", req.Status)
	}
	var count int64
	if err := conn.Count(&count).Offset((req.Page - 1) * req.Size).Limit(req.Size).Find(&apis).Error; err != nil {
		return nil, 0, err
	}
	return apis, count, nil
}

func ModifyApi(id int, api Api, req *ModifyApiReq) error {
	conn := db.DB.Table(localApi.TableName()).Where("id = ?", id)
	switch req.Type {
	case 1:
		conn = conn.Updates(&api)
	case 2:
		conn = conn.Update("deleted_at", nil)
	default:
		return errors.New("the request type of get menu list is only supported 1 or 2")
	}
	if err := conn.Error; err != nil {
		db.DB.Rollback()
		return err
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

func GetApiGroupList(req *GetApiGroupListReq) ([]ApiGroup, int64, error) {
	var groups []ApiGroup
	var conn *gorm.DB
	switch req.Active {
	case 1:
		conn = db.DB.Model(&ApiGroup{}).Order("id")
	case 2:
		conn = db.DB.Model(&ApiGroup{}).Order("id").Unscoped()
	default:
		return nil, 0, errors.New("the request active type of get menu list is only supported 1 or 2")
	}
	switch req.Type {
	case 1:
	case 2:
		conn = conn.Preload("Children", "active", 1)
	default:
		return nil, 0, errors.New("the request type of get menu list is only supported 1 or 2")
	}
	name := req.Name
	if name != "" {
		conn = conn.Where("name LIKE ?", fmt.Sprintf("%%%s%%", name))
	}
	prefix := strings.TrimSpace(req.Prefix)
	if prefix != "" {
		conn = conn.Where("description LIKE ?", fmt.Sprintf("%%%s%%", prefix))
	}
	description := strings.TrimSpace(req.Description)
	if description != "" {
		conn = conn.Where("description LIKE ?", fmt.Sprintf("%%%s%%", description))
	}
	var count int64
	if err := conn.Count(&count).Offset((req.Page - 1) * req.Size).Limit(req.Size).Find(&groups).Error; err != nil {
		return nil, 0, err
	}
	return groups, count, nil

}

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
