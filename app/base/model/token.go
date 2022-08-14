package model

import "time"

type Token struct {
	Account       string    `json:"account"`
	IP            string    `json:"ip"`
	TokenString   string    `json:"token_string"`
	LastLoginTime time.Time `json:"last_login_time"`
}
