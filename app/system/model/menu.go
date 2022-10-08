package model

import (
	"LdapAdmin/common/model"
	"LdapAdmin/db"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strings"
	"time"
)

type Menu struct {
	ID          int    `gorm:"type:int;primaryKey;autoIncrement;not null;comment:the id of menu" json:"id"`
	Active      int    `gorm:"type:int;not null;comment:the active status of menu : 1 active 2 archived" json:"active"`
	Status      int    `gorm:"type:int;not null;comment:the use status of menu: 1 enable 2 disable" json:"status"`
	ParentID    int    `gorm:"type:int;comment:the parent id of menu" json:"parent_id"`
	Name        string `gorm:"type:varchar(255);not null;comment:the name of menu" json:"name"`
	Path        string `gorm:"type:varchar(126);unique;not null;comment:the path of route" json:"path"`
	Description string `gorm:"type:varchar(510);comment:the description for menu" json:"description"`
	Children    []Menu `gorm:"foreignKey:ParentID" json:"children,omitempty"`
	model.StringModel
}

var localMenu Menu

func (m *Menu) TableName() string {
	return "ldap_admin_menus"
}

type AddMenuReq struct {
	Status      int    `json:"status" binding:"required,oneof=1 2"` //Menu's used status
	Name        string `json:"name" binding:"required"`             //Menu's name
	Path        string `json:"path" binding:"required"`             //Menu's path
	Description string `json:"description"`                         //The description of menu
	ParentID    int    `json:"parent_id"`                           //The id of menu's parent
}

type DeleteMenuReq struct {
	IDS  []int `json:"ids" binding:"gt=0"`       //The array of id that is you want to delete
	Type int   `json:"type" binding:"oneof=1 2"` //The way of deleting the menu, 1 archived 2 forever
}

type GetMenuListReq struct {
	Active      int    `form:"active" binding:"required,oneof=1 2"` //Search menu's by active: 1 active, 2 archived
	Type        int    `form:"type" binding:"required,oneof=1 2"`   //Search type: 1 normal, 2 cascade
	Status      int    `form:"status" binding:"oneof=0 1 2"`        //The status of the menu: 1 enable 2 disable
	Name        string `form:"name"`                                //Search menus by name
	Path        string `form:"path"`                                //Search menus by path
	Description string `form:"description"`                         //Search menus by description
	ParentID    int    `form:"parent_id"`                           //Search menus by parent id
	model.PaginationOption
}

type ModifyMenuReq struct {
	ID          int    `json:"id" binding:"required"`              //The id for modify
	Type        int    `json:"type" binding:"required,oneof= 1 2"` //The type for modify: 1 normal, 2 unarchived
	OldPath     string `json:"old_path" binding:"required"`        //The old path
	Path        string `json:"path"`                               //The new path
	Status      int    `json:"status"`                             //The new status: 1 enable 2 disable
	Name        string `json:"name"`                               //The new name
	Description string `json:"description"`                        //The new description
}

func AddMenu(menu Menu) (int, error) {
	if err := db.DB.Table(localMenu.TableName()).
		Create(&menu).
		Error; err != nil {
		db.DB.Rollback()
		return 0, err
	}
	return menu.ID, nil
}

func DeleteMenu(req *DeleteMenuReq) error {
	switch req.Type {
	case 1:
		if err := recursionDelete(req.IDS, 1).Error; err != nil {
			db.DB.Rollback()
			return err
		}
	case 2:
		if err := recursionDelete(req.IDS, 2).Error; err != nil {
			db.DB.Rollback()
			return err
		}
	default:
		return errors.New("the request type of delete menu is only supported 1 or 2")
	}
	return nil
}

func GetMenuList(req *GetMenuListReq) ([]Menu, int64, error) {
	var menus []Menu
	var conn *gorm.DB
	switch req.Active {
	case 1:
		conn = db.DB.Model(&Menu{}).Order("id")
	case 2:
		conn = db.DB.Model(&Menu{}).Order("id").Unscoped()
	default:
		return nil, 0, errors.New("the request active type of get menu list is only supported 1 or 2")
	}
	switch req.Type {
	case 1:
	case 2:
		conn = conn.Preload("Children", recursionPreload, "active", 1)
	default:
		return nil, 0, errors.New("the request type of get menu list is only supported 1 or 2")
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
	if req.Type == 2 {
		conn = conn.Where("parent_id IS NULL")
	}
	var count int64
	if err := conn.Count(&count).Offset((req.Page - 1) * req.Size).Limit(req.Size).Find(&menus).Error; err != nil {
		return nil, 0, err
	}
	return menus, count, nil
}

func GetMenuByIDAndPath(id int, path string) (*Menu, error) {
	var menu Menu
	conn := db.DB.Table(localMenu.TableName())
	if id != 0 {
		conn.Where("id = ?", id)
	}
	if path != "" {
		conn.Where("path = ?", path)
	}
	if err := conn.First(&menu).
		Error; err != nil {
		return nil, err
	}
	return &menu, nil
}

func ModifyMenu(id int, menu Menu, req *ModifyMenuReq) error {
	conn := db.DB.Table(localMenu.TableName()).Where("id = ?", id)
	switch req.Type {
	case 1:
		conn = conn.Updates(&menu)
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

func ModifyMenuPathBatch(oldPath string, path string) error {
	if err := db.DB.Table(localMenu.TableName()).
		Where("path LIKE ?", fmt.Sprintf("%s%%", oldPath)).
		Update("path", gorm.Expr(fmt.Sprintf("overlay(path placing '%s' from 1 for %d)", path, len(oldPath)))).
		Error; err != nil {
		return err
	}
	return nil
}

/* $ Tools */

func recursionPreload(db *gorm.DB) *gorm.DB {
	return db.Preload("Children", recursionPreload)
}

func recursionDelete(ids []int, deleteType int) *gorm.DB {
	if deleteType == 1 {
		return db.DB.Exec("with recursive menus as ("+
			"select * from ldap_admin_menus where id IN (?)"+
			" union all "+
			"select m1.* from ldap_admin_menus m1, menus m2 where m2.id = m1.parent_id"+
			") update ldap_admin_menus set deleted_at = ?, active = ?", ids, time.Now(), 2)
	}
	if deleteType == 2 {
		return db.DB.Exec("with recursive menus as ("+
			"select * from ldap_admin_menus where id IN (?)"+
			" union all "+
			"select m1.* from ldap_admin_menus m1, menus m2 where m2.id = m1.parent_id"+
			") delete from ldap_admin_menus", ids)
	}
	return db.DB
}
