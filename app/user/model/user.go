package model

import (
	"LdapAdmin/common/model"
	"LdapAdmin/common/util"
)

type User struct {
	ID           int             `gorm:"type:int;primaryKey;autoIncrement;not null;comment:user's id" json:"id"`
	Active       int             `gorm:"type:int;not null;comment:user's exist status: 1 exist 2 not_exist" json:"active"`
	Status       int             `gorm:"type:int;not nul;comment:user's status: 1 active 2 inactive" json:"status"`
	Account      string          `gorm:"type:varchar(20);not null;unique;comment:user's account" json:"account"`
	Password     string          `gorm:"type:varchar(126);not null;comment:user's password" json:"password"`
	FullName     string          `gorm:"type:varchar(36);not null;comment:user's true name" json:"full_name"`
	Surname      string          `gorm:"type:varchar(18);not null;comment:user's surname" json:"surname"`
	GivenName    string          `gorm:"type:varchar(18);not nul;comment:user's given_name" json:"given_name"`
	DisplayName  string          `gorm:"type:varchar(36);not nul;comment:user's display_name" json:"display_name"`
	Avatar       string          `gorm:"type:varchar(255);comment:user's avatar" json:"avatar"`
	Gender       int             `gorm:"type:int;not null;comment:user's gender, 1 male 2 female 3 other" json:"gender"`
	Email        string          `gorm:"type:varchar(60);unique;comment:user's email" json:"email"`
	Mobile       string          `gorm:"type:varchar(20);unique;comment:user's mobile" json:"mobile"`
	Address      string          `gorm:"type:varchar(255);comment:user's address" json:"address"`
	Birthday     *util.LocalTime `gorm:"type:timestamp;not null;comment:user's birth date" json:"birthday"`
	Introduction string          `gorm:"type:varchar(900);comment:user's introduction" json:"introduction"`
	model.StringModel
}

func (u *User) TableName() string {
	return "ldap_admin_users"
}

type AddUserReq struct {
}

type DeleteUserReq struct {
}

type GetUsersReq struct {
}

type ModifyUserReq struct {
}

func AddUser() {
}

func DeleteUser() {
}

func GetUsers() {
}

func ModifyUser() {
}
