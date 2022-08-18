package model

import (
	"LdapAdmin/common/model"
	"time"
)

type Token struct {
	ID            int       `json:"id"`
	Account       string    `json:"account"`
	IP            string    `json:"ip"`
	TokenString   string    `json:"token_string"`
	LastLoginTime time.Time `json:"last_login_time"`
	model.StringModel
}

var localToken Token

func (token *Token) TableName() string {
	return "ldap_admin_tokens"
}

func AddToken() {

}

func DeleteToken() {

}

func GetTokens() {

}

func ModifyToken() {

}
