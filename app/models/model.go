package models

import (
	"time"
)

type BaseModel struct {
	ID uint64 `gorm:"primaryKey" json:"id"`
}

type TimestampsField struct {
	CreatedAt time.Time `json:"created_at" gorm:"type:DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3)"`
	UpdatedAt time.Time `json:"updated_at" gorm:"type:DATETIME(3) NOT NULL DEFAULT CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3)"`
}
