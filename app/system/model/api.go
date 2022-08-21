package model

type Api struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
}

var localApi Api

func (a *Api) TableName() string {
	return "ldap_admin_apis"
}

func AddApi() {
	
}

func DeleteApi() {
	
}

func GetApiList() {
	
}

func ModifyApi()  {
	
}