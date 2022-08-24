package model

import (
	"LdapAdmin/common/model"
	"LdapAdmin/db"
)

type Menu struct {
	ID          int    `gorm:"type:int;primaryKey;autoIncrement;not null;comment:menu's id" json:"id"`
	Active      int    `gorm:"type:int;not null;comment:menu's status 1 active 2 inactive" json:"active"`
	Name        string `gorm:"type:varchar(255);not null;comment:menu's name" json:"name"`
	Path        string `gorm:"type:varchar(126);not null;comment:the path of route" json:"path"`
	Description string `gorm:"type:varchar(510);comment:the description for menu" json:"description"`
	model.StringModel
}

var localMenu Menu

func (m *Menu) TableName() string {
	return "ldap_admin_menus"
}

type AddMenuReq struct {
	Name        string `json:"name" binding:"required"` //menu's name
	Path        string `json:"path" binding:"required"` //menu's path
	Description string `json:"description"`             //the description of menu
}

type DeleteMenuReq struct {
	IDS []int `json:"ids"` //the array of id that is you want to delete
}

type GetMenuListReq struct {
	Name        string `form:"name"`        //search menus by name
	Path        string `form:"path"`        //search menus by path
	Description string `form:"description"` //search menus by description
	model.PaginationOption
}

type ModifyMenuReq struct {
	ID          int    `json:"id" binding:"required"` //the id for modify
	Name        string `json:"name"`                  //the new name
	Path        string `json:"path"`                  //the new path
	Description string `json:"description"`           //the new description
}

func AddMenu(menu Menu) (int, error) {
	if err := db.DB.Table(localMenu.TableName()).
		Create(&menu).
		Error; err != nil {
		return 0, err
	}
	return menu.ID, nil
}

func DeleteMenu(ids []int) error {
	if err := db.DB.Table(localMenu.TableName()).
		Where("id IN (?)", ids).
		Delete(&Menu{}).
		Error; err != nil {
		return err
	}
	return nil
}

func GetMenuList() {

}

func ModifyMenu(id int, menu Menu) error {
	if err := db.DB.Table(localMenu.TableName()).
		Where("id = ?", id).
		Updates(&menu).
		Error; err != nil {
		return err
	}
	return nil
}
