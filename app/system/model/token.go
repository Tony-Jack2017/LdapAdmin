package model

import (
	"LdapAdmin/common/model"
	"LdapAdmin/db"
	"strings"
)

type Token struct {
	ID          int    `gorm:"type:int;primaryKey;autoIncrement;comment:token's id" json:"id"`
	Account     string `gorm:"type:varchar(20);not null;comment:the user's account" json:"account"`
	IP          string `gorm:"type:varchar(20);not null;comment:the ip of where is created token" json:"ip"`
	TokenString string `gorm:"type:varchar(255);not null;comment:token string" json:"token_string"`
	LoginTime   string `gorm:"type:varchar(50);not null;comment:the time of user last login" json:"login_time"`
	Active      int    `gorm:"type:int;not null;comment:token's status 1 active 2 inactive'" json:"active"`
	model.StringModel
}

var localToken Token

func (token *Token) TableName() string {
	return "ldap_admin_tokens"
}

type AddTokenReq struct {
	Account       string `json:"account"`         //required
	IP            string `json:"ip"`              //required
	TokenString   string `json:"token_string"`    //required
	LastLoginTime string `json:"last_login_time"` //required
}

type GetTokenReq struct {
	Account string `json:"account"`
	IP      string `json:"ip"`
}

type ModifyTokenReq struct {
	ID            int    `json:"id"` //required
	IP            string `json:"ip"`
	TokenString   string `json:"token_string"`
	LastLoginTime string `json:"last_login_time"`
}

func AddToken(token Token) (int, error) {
	if err := db.DB.Table(localToken.TableName()).
		Create(&token).
		Error; err != nil {
		return 0, err
	}
	return token.ID, nil
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
		Updates(token).
		Error; err != nil {
		return err
	}
	return nil
}
