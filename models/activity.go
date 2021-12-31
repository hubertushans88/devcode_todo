package models

import (
	"gorm.io/gorm"
	"time"
)

type Activity struct{
	ID 			uint			`gorm:"primaryKey" json:"id"`
	Email 		*string 		`json:"email"`
	Title		string			`json:"title"`

	CreatedAt 	time.Time      `json:"created_at"`
	UpdatedAt 	time.Time      `json:"updated_at"`
	DeletedAt 	gorm.DeletedAt `json:"deleted_at"`
}
