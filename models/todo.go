package models

import (
	"gorm.io/gorm"
	"time"
)

type Todo struct {
	ID              uint   `gorm:"primaryKey" json:"id"`
	ActivityGroupID uint   `json:"activity_group_id"`
	Title           string `gorm:"not_null" json:"title"`
	IsActive        bool   `gorm:"not_null;default:true" json:"is_active"`
	Priority        string `gorm:"not_null;default:very-high"json:"priority"`

	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}
