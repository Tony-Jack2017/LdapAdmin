package model

type LoginReq struct {
	IP       string `json:"ip" binding:"required"`
	Account  string `json:"account" binding:"required"`  //the account for login
	Password string `json:"password" binding:"required"` //the password for login
}
