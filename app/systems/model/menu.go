package model

type Menu struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Route string `json:"route"`
}

var localMenu Menu

func (m *Menu) TableName() string {
	return "ldap_admin_menus"
}
