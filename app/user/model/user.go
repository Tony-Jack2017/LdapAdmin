package model

import "LdapAdmin/common/model"

type User struct {
	ID           int    `gorm:"type:int;comment:user's id" json:"id"`
	Active       int    `json:"active"`
	Account      string `gorm:"" json:"account"`
	Password     string `gorm:"" json:"password"`
	FullName     string `gorm:"" json:"full_name"`
	Surname      string `gorm:"" json:"surname"`
	GivenName    string `gorm:"" json:"given_name"`
	DisplayName  string `gorm:"" json:"display_name"`
	Avatar       string `gorm:"" json:"avatar"`
	Gender       int    `gorm:"" json:"gender"`
	Email        string `gorm:"" json:"email"`
	Mobile       string `gorm:"" json:"mobile"`
	Address      string `gorm:"" json:"address"`
	Birthday     string `gorm:"" json:"birthday"`
	Introduction string `gorm:"" json:"introduction"`
	Status       int    `json:"status"`
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
