package model

import (
	"gorm.io/gorm"
	"time"
)

type StringModel struct {
	CreatedAt time.Time      `gorm:"type:timestamp;autoCreateTime;comment:the time of created data" json:"created_at"`
	UpdatedAt time.Time      `gorm:"type:timestamp;autoUpdateTime;comment:the time of updated data" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"type:timestamp;comment:the time of deleted data" json:"deleted_at"`
}
