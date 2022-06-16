package models

import (
	"time"
)

type Todo struct {
	Id        int       `gorm:"primaryKey" json:"id"`
	Completed bool      `json:"completed" `
	Text      string    `json:"text"`
	CreatedAt time.Time `gorm:"autoCreateTime:true" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoCreateTime:true" json:"updated_at"`
}
