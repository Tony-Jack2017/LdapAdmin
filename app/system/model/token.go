package model

import (
	"LdapAdmin/common/model"
	"LdapAdmin/db"
)

type Token struct {
	ID          int    `json:"id"`
	Account     string `json:"account"`
	IP          string `json:"ip"`
	TokenString string `json:"token_string"`
	LoginTime   string `json:"login_time"`
	Active      int    `json:"active"`
	model.StringModel
}

var localToken Token

func (token *Token) TableName() string {
	return "ldap_admin_tokens"
}

type AddTokenReq struct {
	Account       string `json:"account"`
	IP            string `json:"ip"`
	TokenString   string `json:"token_string"`
	LastLoginTime string `json:"last_login_time"`
}

type DeleteTokenReq struct {
}

type GetTokensReq struct {
}

func AddToken(token Token) error {
	if err := db.DB.Table(localToken.TableName()).Create(&token).Error; err != nil {
		return err
	}
	return nil
}

func ModifyToken() {

}
