package model

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

// Item is a structure fot the table items
type Item struct {
	ID         uuid.UUID      `json:"id" gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	AddisberID string         `json:"addisber_id" validate:"required"`
	Title      string         `json:"title" validate:"required"`
	Price      string         `json:"price" validate:"required"`
	URL        string         `json:"url" validate:"required"`
	PictureURL string         `json:"picture_url" validate:"required"`
	CreatedAt  time.Time      `json:"created_at,omitempty"`
	UpdatedAt  time.Time      `json:"updated_at,omitempty"`
	DeletedAt  gorm.DeletedAt `json:"-" gorm:"index"`
}
