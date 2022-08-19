package model

import (
	"LdapAdmin/common/model"
	"time"
)

type Token struct {
	ID            int       `gorm:"type:int;primaryKey;autoIncrement;comment:token's id" json:"id"`
	Account       string    `json:"account"`
	IP            string    `gorm:"type:varchar(128);comment:the source ip of token" json:"ip"`
	TokenString   string    `gorm:"type:varchar()" json:"token_string"`
	LastLoginTime time.Time `json:"last_login_time"`
	model.StringModel
}

var localToken Token

func (token *Token) TableName() string {
	return "ldap_admin_tokens"
}

func AddToken() {

}

func GetTokens() {

}

func ModifyToken() {

}
