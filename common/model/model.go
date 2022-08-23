package model

import (
	"LdapAdmin/common/util"
)

type StringModel struct {
	CreatedAt *util.LocalTime `gorm:"type:timestamp;comment:the time of created data" json:"created_at"`
	UpdatedAt *util.LocalTime `gorm:"type:timestamp;comment:the time of updated data" json:"updated_at"`
	DeletedAt *util.LocalTime `gorm:"type:timestamp;comment:the time of deleted data" json:"deleted_at"`
}
