package model

import (
	"time"

	"gorm.io/gorm"
)

type TimeAt struct {
	CreatedAt *time.Time      `json:"created_at,omitempty"`
	UpdatedAt *time.Time      `json:"updated_at,omitempty"`
	DeletedAt *gorm.DeletedAt `json:"deleted_at,omitempty"`
}
