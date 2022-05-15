package model

import (
	"time"
)

type Common struct {
	ID        uint64    `json:"id" gorm:"primarykey"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
