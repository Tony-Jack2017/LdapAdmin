package model

type LoginReq struct {
	Account  string `json:"account" binding:"required"`  //the account for login
	Password string `json:"password" binding:"required"` //the password for login
}
