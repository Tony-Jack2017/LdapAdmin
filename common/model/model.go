package model

import "time"

type StringModel struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}

type IntModel struct {
	CreatedAt int8 `json:"created_at"`
	UpdatedAt int8 `json:"updated_at"`
	DeletedAt int8 `json:"deleted_at"`
}
