package model

type User struct {
	ID       int    `json:"id"`
	Account  string `json:"account"`
	Password string `json:"password"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Mobile   string `json:"mobile"`
}

func TableName() string {
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
