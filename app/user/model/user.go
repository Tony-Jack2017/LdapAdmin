package model

import "LdapAdmin/common/model"

type User struct {
	ID           int    `json:"id"`
	Account      string `json:"account"`
	Password     string `json:"password"`
	FullName     string `json:"full_name"`
	Surname      string `json:"surname"`
	GivenName    string `json:"given_name"`
	DisplayName  string `json:"display_name"`
	Avatar       string `json:"avatar"`
	Gender       int    `json:"gender"`
	Email        string `json:"email"`
	Mobile       string `json:"mobile"`
	Address      string `json:"address"`
	Birthday     string `json:"birthday"`
	Introduction string `json:"introduction"`
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

func AddUser()  {

}

func DeleteUser() {

}

func GetUsers() {

}

func ModifyUser() {

}