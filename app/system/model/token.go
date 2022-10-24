package model

import (
	"LdapAdmin/common/model"
	"LdapAdmin/database/db"
	"strings"
)

type Token struct {
	ID          int    `gorm:"type:int;primaryKey;autoIncrement;comment:the id of token" json:"id"`
	Account     string `gorm:"type:varchar(20);unique;not null;comment:the account of token" json:"account"`
	IP          string `gorm:"type:varchar(15);not null;comment:the ip where will use token" json:"ip"`
	TokenString string `gorm:"type:varchar(255);not null;comment:the token strings" json:"token_string"`
	model.StringModel
}

var localToken Token

func (token *Token) TableName() string {
	return "ldap_admin_tokens"
}

type AddTokenReq struct {
	Account     string `json:"account"`
	IP          string `json:"ip"`
	TokenString string `json:"token_string"`
}

type GetTokenReq struct {
	Account string `json:"account"` //Search token by account
	IP      string `json:"IP"`      //Search token by ip
}

func AddToken(token Token) error {
	if err := db.DB.Table(localToken.TableName()).
		Create(&token).
		Error; err != nil {
		return err
	}
	return nil
}

func GetToken(req *GetTokenReq) (*Token, error) {
	var token Token
	conn := db.DB.Table(localToken.TableName()).Order("id")
	account := strings.TrimSpace(req.Account)
	if account != "" {
		conn = conn.Where("account = ?", account)
	}
	ip := strings.TrimSpace(req.IP)
	if ip != "" {
		conn = conn.Where("ip = ?", ip)
	}
	if err := conn.Find(&token).Error; err != nil {
		return nil, err
	}
	return &token, nil
}

func ModifyToken(id int, token Token) error {
	if err := db.DB.Table(localToken.TableName()).
		Where("id = ?", id).
		Updates(&token).
		Error; err != nil {
		return err
	}
	return nil
}
