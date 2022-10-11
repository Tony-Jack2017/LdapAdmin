package model

import (
	"database/sql/driver"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type LocalTime time.Time // local time

type StringModel struct {
	CreatedAt LocalTime      `gorm:"type:timestamp;autoCreateTime;comment:the time of created data" json:"created_at"`
	UpdatedAt LocalTime      `gorm:"type:timestamp;autoUpdateTime;comment:the time of updated data" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"type:timestamp;comment:the time of deleted data" json:"deleted_at"`
}

// MarshalJSON
// @override
func (t *LocalTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format("2006-01-02 15:04:05"))), nil
}

func (t LocalTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	tlt := time.Time(t)
	if tlt.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return tlt, nil
}

func (t *LocalTime) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*t = LocalTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
