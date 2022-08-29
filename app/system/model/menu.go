package model

import (
	"LdapAdmin/common/model"
	"LdapAdmin/db"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"strings"
)

type Menu struct {
	ID          int    `gorm:"type:int;primaryKey;autoIncrement;not null;comment:menu's id" json:"id"`
	Active      int    `gorm:"type:int;not null;comment:menu's active's status 1 active 2 archived" json:"active"`
	Status      int    `gorm:"type:int;not null;comment:menu's use's status 1 enable 2 disable" json:"status"`
	ParentID    int    `gorm:"type:int;not null;default:0;comment:menu's parent id" json:"parent_id"`
	Name        string `gorm:"type:varchar(255);not null;comment:menu's name" json:"name"`
	Path        string `gorm:"type:varchar(126);not null;comment:the path of route" json:"path"`
	Description string `gorm:"type:varchar(510);comment:the description for menu" json:"description"`
	Children    []Menu `gorm:"foreignKey:ParentID;associate_foreignKey:ID" json:"children"`
	model.StringModel
}

var localMenu Menu

func (m *Menu) TableName() string {
	return "ldap_admin_menus"
}

type AddMenuReq struct {
	Status      int    `json:"status" binding:"required"` //Menu's used status
	Name        string `json:"name" binding:"required"`   //Menu's name
	Path        string `json:"path" binding:"required"`   //Menu's path
	Description string `json:"description"`               //The description of menu
	ParentID    int    `json:"parent_id"`                 //The id of menu's parent
}

type DeleteMenuReq struct {
	IDS  []int `json:"ids"`  //The array of id that is you want to delete
	Type int   `json:"type"` //The way of deleting the menu, 1 archived 2 forever
}

type GetMenuListReq struct {
	Active      int    `json:"active"`      //Search menu's by active
	Name        string `form:"name"`        //Search menus by name
	Path        string `form:"path"`        //Search menus by path
	Description string `form:"description"` //Search menus by description
	ParentID    int    `json:"parent_id"`   //Search menus by parent id
	model.PaginationOption
}

type ModifyMenuReq struct {
	ID          int    `json:"id" binding:"required"` //The id for modify
	Status      int    `json:"status"`                //The new status
	Name        string `json:"name"`                  //The new name
	Path        string `json:"path"`                  //The new path
	Description string `json:"description"`           //The new description
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
		if err := db.DB.Table(localMenu.TableName()).
			Where("id IN (?)", req.IDS).
			Update("active", 2).
			Delete(&Menu{}).
			Error; err != nil {
			db.DB.Rollback()
			return err
		}
	case 2:
		if err := db.DB.Table(localMenu.TableName()).Unscoped().
			Where("id IN (?)", req.IDS).
			Delete(&Menu{}).
			Error; err != nil {
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
		conn = db.DB.Table(localMenu.TableName()).Order("id")
	case 2:
		conn = db.DB.Table(localMenu.TableName()).Order("id").Unscoped()
	default:
		return nil, 0, errors.New("the request type of delete menu is only supported 1 or 2")
	}

	name := strings.TrimSpace(req.Name)
	if name != "" {
		conn = conn.Where("name LIKE ?", fmt.Sprintf("%%%s%%", name))
	}
	path := strings.TrimSpace(req.Path)
	if name != "" {
		conn = conn.Where("path LIKE ?", fmt.Sprintf("%%%s%%", path))
	}
	description := strings.TrimSpace(req.Description)
	if name != "" {
		conn = conn.Where("description LIKE ?", fmt.Sprintf("%%%s%%", description))
	}

	var count int64
	if err := conn.Count(&count).Offset((req.Page - 1) * req.Size).Limit(req.Size).Find(&menus).Error; err != nil {
		return nil, 0, err
	}
	return menus, count, nil

}

func ModifyMenu(id int, menu Menu) error {
	if err := db.DB.Table(localMenu.TableName()).
		Where("id = ?", id).
		Updates(&menu).
		Error; err != nil {
		db.DB.Rollback()
		return err
	}
	return nil
}
