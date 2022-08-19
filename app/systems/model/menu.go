package model

type Menu struct {
	ID    int    `gorm:"type:int;primaryKey;autoIncrement;" json:"id"`
	Name  string `gorm:"" json:"name"`
	Route string `gorm:"" json:"route"`
}

var localMenu Menu

func (m *Menu) TableName() string {
	return "ldap_admin_menus"
}
