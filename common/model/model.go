package model

import "time"

type StringModel struct {
	CreatedAt time.Time `gorm:"" json:"created_at"`
	UpdatedAt time.Time `gorm:"" json:"updated_at"`
	DeletedAt time.Time `gorm:"" json:"deleted_at"`
}

type IntModel struct {
	CreatedAt int `json:"created_at"`
	UpdatedAt int `json:"updated_at"`
	DeletedAt int `json:"deleted_at"`
}
