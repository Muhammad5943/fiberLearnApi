package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Note struct {
	*gorm.Model // Adds some metadata to the table

	ID       uuid.UUID `gorm:"type:uuid"` // Explicitly specify
	Title    string    `json:"title"`
	SubTitle string    `json:"sub_title"`
	Text     string    `json:"text"`
}
